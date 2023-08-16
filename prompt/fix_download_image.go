package prompt

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/samsonmxvi/go-wenku-downloader/downloader"
	"github.com/samsonmxvi/go-wenku-downloader/scraper"
)

func FixDownloadImage(novelId int) error {
	downloadPath := strconv.Itoa(novelId)

	_, err := os.Stat(downloadPath)
	if os.IsNotExist(err) {
		return nil
	}

	// 读取目录中的文件和文件夹
	files, err := os.ReadDir(downloadPath)
	if err != nil {
		return err
	}

	// 遍历并打印出所有文件夹
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		// fmt.Println(file.Name())
		volumePath := path.Join(downloadPath, file.Name())

		// 读取目录中的文件和文件夹
		volumePathFiles, err := os.ReadDir(volumePath)
		if err != nil {
			return err
		}

		// 遍历并打印出所有 .json 文件
		for _, jsonFile := range volumePathFiles {
			if !jsonFile.IsDir() && strings.HasSuffix(jsonFile.Name(), ".json") {
				chapter := &scraper.Chapter{}

				// fmt.Println(file.Name())
				jsonPath := path.Join(volumePath, jsonFile.Name())
				jsonByte, err := os.ReadFile(jsonPath)
				if err != nil {
					continue
				}

				json.Unmarshal(jsonByte, chapter)

				if len(chapter.Content.Images) > 0 {
					for _, imgUrl := range chapter.Content.Images {
						fmt.Printf("正在修复图片 %v %v \n", jsonPath, imgUrl)
						for i := 0; i < 3; i++ {
							err := downloader.DownloadImage(imgUrl, path.Join(volumePath, downloader.ImageFolderName))
							if err == nil {
								break
							} else {
								time.Sleep(6 * time.Second) // temp fix rate limit
								continue
							}
						}
					}
				}
			}
		}
	}

	return nil
}
