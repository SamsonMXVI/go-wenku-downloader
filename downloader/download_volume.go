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

func DownloadVolume(volume *scraper.Volume, dirPath string) error {
	var imageArray []string
	imageDirPath := path.Join(dirPath, ImageFolderName)

	if err := util.CheckDir(dirPath); err != nil {
		return err
	}

	chaterArray, err := getChapterArray(volume)
	if err != nil {
		return err
	}

	progressBar := pb.StartNew(len(chaterArray))
	for i, chapter := range chaterArray {
		progressBar.SetTemplateString(getChapterTemplateString(volume.Name, i))
		// check file exist
		chapterFile := path.Join(dirPath, fmt.Sprintf("%d.json", chapter.Index))
		if util.CheckFileExist(chapterFile) {
			getChapterContentFromFile(chapterFile, chapter)
		} else {
			err := getChaterContent(chapter)
			if err != nil {
				return err
			}
			// save chapter to file
			err = DownloadChapter(chapter, dirPath)
			if err != nil {
				log.Printf("download chapter error %v \n", err)
				return err
			}
		}

		if chapter.Content.Images != nil && len(chapter.Content.Images) != 0 {
			imageArray = append(imageArray, chapter.Content.Images...)
		}

		progressBar.Increment()
	}
	progressBar.Finish()

	for _, imageURL := range imageArray {
		success := false
		for i := 0; i < 3; i++ {
			err := DownloadImage(imageURL, imageDirPath)
			if err == nil {
				success = true
				break
			} else {
				time.Sleep(3 * time.Second) // temp fix rate limit
				continue
			}
		}
		if !success {
			return fmt.Errorf("图片下载错误")
		}
	}

	return nil
}

func getChapterArray(volume *scraper.Volume) ([]*scraper.Chapter, error) {
	for i := 0; i < 3; i++ {
		chaterArray, err := scraper.GetChapterArray(volume)
		if err == nil {
			time.Sleep(1 * time.Second)
			return chaterArray, nil
		} else {
			log.Printf("获取章节列表失败 %v, 重试第%v次 \n", err, i)
			time.Sleep(6 * time.Second) // temp fix rate limit
			continue
		}
	}
	return nil, fmt.Errorf("获取章节列表失败")
}

func getChaterContent(chapter *scraper.Chapter) error {
	var err error
	for i := 0; i < 3; i++ {
		err = scraper.GetChapterContent(chapter)
		if err == nil {
			time.Sleep(1 * time.Second)
			return nil
		} else {
			time.Sleep(6 * time.Second) // temp fix rate limit
			log.Printf("获取章节内容失败 %v, 重试第%v次 \n", err, i)
			continue
		}
	}
	return err
}

func getChapterContentFromFile(path string, chapter *scraper.Chapter) {
	file, _ := os.ReadFile(path)
	_ = json.Unmarshal([]byte(file), chapter)
}
