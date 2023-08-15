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

	chaterArray, err := scraper.GetChapterArray(volume)
	time.Sleep(1 * time.Second) // temp fix rate limit

	if err != nil {
		return fmt.Errorf("获取章节列表失败")
	}

	progressBar := pb.StartNew(len(chaterArray))
	for i, chapter := range chaterArray {
		progressBar.SetTemplateString(getChapterTemplateString(volume.Name, i))
		// check file exist
		chapterFile := path.Join(dirPath, fmt.Sprintf("%d.json", chapter.Index))
		if util.CheckFileExist(chapterFile) {
			getChapterContentFromFile(chapterFile, chapter)
		} else {
			getChaterContent(chapter)
			// save chapter to file
			err = DownloadChapter(chapter, dirPath)
			if err != nil {
				log.Printf("download chapter error %v \n", err)
			}
		}

		if chapter.Content.Images != nil && len(chapter.Content.Images) != 0 {
			imageArray = append(imageArray, chapter.Content.Images...)
		}

		progressBar.Increment()
	}
	progressBar.Finish()

	for _, imageURL := range imageArray {
		for i := 0; i < 3; i++ {
			err := DownloadImage(imageURL, imageDirPath)
			if err == nil {
				break
			} else {
				time.Sleep(1 * time.Second) // temp fix rate limit
			}
		}
	}

	return nil
}

func getChaterContent(chapter *scraper.Chapter) {
	for i := 0; i < 3; i++ {
		err := scraper.GetChapterContent(chapter)
		if err == nil {
			time.Sleep(1 * time.Second)
			break
		} else {
			time.Sleep(3 * time.Second) // temp fix rate limit
			log.Printf("get chapter content error %v, retry %v \n", err, i)
			continue
		}
	}
}

func getChapterContentFromFile(path string, chapter *scraper.Chapter) {
	file, _ := os.ReadFile(path)
	_ = json.Unmarshal([]byte(file), chapter)
}
