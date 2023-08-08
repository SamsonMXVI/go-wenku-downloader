package prompt

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bmaupin/go-epub"
	"github.com/samsonmxvi/go-wenku-downloader/downloader"
	"github.com/samsonmxvi/go-wenku-downloader/scraper"
	"github.com/samsonmxvi/go-wenku-downloader/util"
)

func download(novelId int) error {
	downloadPath := strconv.Itoa(novelId)
	if err := util.CheckDir(downloadPath); err != nil {
		return fmt.Errorf("创建目录失败: %e", err)
	}

	// get novel metadata
	novel, err := promptNovelDetails(int(novelId))
	if err != nil {
		return fmt.Errorf("获取小说信息失败: %e", err)
	}

	// download novel metadata and coverImg
	downloader.DownloadNovelMetadata(novel, downloadPath)

	// download cover image
	success := false
	for i := 0; i < 3; i++ {
		err = downloader.DownloadImage(novel.Cover, downloadPath)
		if err == nil {
			success = true
			break
		}
	}
	if !success {
		return fmt.Errorf("download cover image failed %v", err)
	}

	// get selected volume list
	volumeArray, err := promptVolumeSelect(novel.CatalogueUrl)
	if err != nil {
		return fmt.Errorf("下载小说卷信息失败: %e", err)
	}

	// get coverIndex from input
	converIndex, err := inputCoverIndex()
	if err != nil {
		return fmt.Errorf("prompt failed %v", err)
	}

	// download volume
	for _, volume := range volumeArray {
		time.Sleep(3 * time.Second) // temp fix rate limit
		volumePath := path.Join(downloadPath, formatFilename(volume.Name))
		err = downloader.DownloadVolume(volume, volumePath)
		if err != nil {
			log.Printf("download volume error %v", err)
			continue
		}
		err = createEpub(novel, volume.Name, volume.ChapterCount, converIndex, downloadPath)
		if err != nil {
			log.Printf("create epub failed: %v", err)
			continue
		}
	}

	return nil
}

func createEpub(novel *scraper.Novel, volumeName string, chapterCount int, coverIndex int, downloadPath string) error {
	formatedVolumeName := formatFilename(volumeName)
	var internalImagePath []string
	// output epub path
	var epubFilePath string = path.Join(downloadPath, fmt.Sprintf("%s %s.epub", novel.NovelName, formatedVolumeName))
	// volume path
	var volumePath string = path.Join(downloadPath, formatedVolumeName)
	var imagePath string = path.Join(volumePath, downloader.ImageFolderName)
	var coverPath string = path.Join(downloadPath, util.GetUrlLastString(novel.Cover))

	// create epub
	epub := epub.NewEpub(novel.NovelName + " " + volumeName)
	epub.SetAuthor(novel.Author)

	// add coverImage to epub
	internalConverPath, err := util.AddImage(epub, coverPath)
	if err != nil {
		return fmt.Errorf("add image to epub failed")
	}
	internalImagePath = append(internalImagePath, internalConverPath)

	for i := 1; i <= chapterCount; i++ {
		file, err := os.ReadFile(path.Join(volumePath, fmt.Sprintf("%d.json", i)))
		if err != nil {
			return err
		}
		chapter := &scraper.Chapter{}
		err = json.Unmarshal([]byte(file), &chapter)
		if err != nil {
			return err
		}
		jsonByte, err := json.MarshalIndent(chapter.Content.Article, "", " ")
		jsonStr := strings.Replace(string(jsonByte), "\"", "", -1)
		if err != nil {
			return err
		}
		xhtml := util.CreateSectionXhtml(chapter.Title, jsonStr)
		if len(chapter.Content.Images) != 0 {
			for _, img := range chapter.Content.Images {
				imgFile := path.Join(imagePath, util.GetUrlLastString(img))
				internalPath, _ := util.AddImage(epub, imgFile)
				xhtml = util.AddImageToXhtml(internalPath, xhtml)
				internalImagePath = append(internalImagePath, internalPath)
			}
		}
		err = util.AddSectionXhtml(epub, chapter.Title, xhtml)
		if err != nil {
			return err
		}
	}

	tempConverPath := internalImagePath[0]
	if coverIndex < len(internalImagePath) {
		tempConverPath = internalImagePath[coverIndex]
	}
	epub.SetCover(tempConverPath, "")

	err = epub.Write(epubFilePath)
	if err != nil {
		return err
	}
	return nil
}

func formatFilename(str string) string {
	newFilename := strings.ReplaceAll(str, "/", "-")
	re := regexp.MustCompile(`[<>:"/\\|?*\t]`)
	newFilename = re.ReplaceAllString(newFilename, "")
	return newFilename
}
