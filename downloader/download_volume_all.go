package downloader

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/samsonmxvi/go-wenku-downloader/scraper"
	"github.com/samsonmxvi/go-wenku-downloader/util"
	"gopkg.in/cheggaaa/pb.v2"
)

func DownloadVolumeAll(volume *scraper.Volume, dirPath string) error {
	var imageArray []string
	var downloadError *DownloadVolumnError
	imageDirPath := path.Join(dirPath, ImageFolderName)
	errorPath := path.Join(dirPath, ErrorJsonName)
	if checkError(dirPath) {
		if err := reDownloadError(dirPath); err != nil {
			return fmt.Errorf("重新下载缺失文件失败")
		}
		return nil
	}

	if err := util.CheckDir(dirPath); err != nil {
		return err
	}

	chaterArray, err := scraper.GetChapterArray(volume)

	time.Sleep(1 * time.Second)

	if err != nil {
		return fmt.Errorf("获取章节列表失败")
	}

	progressBar := pb.StartNew(len(chaterArray))

	for i, chapter := range chaterArray {
		progressBar.SetTemplateString(getChapterTemplateString(volume.Name, i))
		// check file exist
		if util.CheckFileExist(path.Join(dirPath, fmt.Sprintf("%d.json", chapter.Index))) {
			progressBar.Increment()
			continue
		}

		err := scraper.GetChapterContent(chapter)
		time.Sleep(3 * time.Second)

		if err != nil {
			log.Printf("get chapter content error %v", err)
			downloadError.Chapter = append(downloadError.Chapter, chapter)
		}

		if chapter.Content.Images != nil && len(chapter.Content.Images) != 0 {
			imageArray = append(imageArray, chapter.Content.Images...)
		}

		err = DownloadChapter(chapter, dirPath)
		if err != nil {
			log.Printf("download chapter error %v", err)
			downloadError.Chapter = append(downloadError.Chapter, chapter)
		}
		progressBar.Increment()
	}
	progressBar.Finish()

	for _, imageURL := range imageArray {
		success := false
		time.Sleep(1 * time.Second)
		for i := 0; i < 3; i++ {
			err := DownloadImage(imageURL, imageDirPath)
			if err == nil {
				success = true
				break
			}
		}
		if !success {
			// download error
			downloadError.Image = append(downloadError.Image, imageURL)
		}
	}

	if downloadError != nil {
		json, err := json.MarshalIndent(downloadError, "", " ")
		if err != nil {
			return fmt.Errorf("marshal downloadError failed: %v", err)
		}
		if err = os.WriteFile(errorPath, json, 0644); err != nil {
			return fmt.Errorf("save downloadError failed: %v", err)
		}
	}

	return nil
}
