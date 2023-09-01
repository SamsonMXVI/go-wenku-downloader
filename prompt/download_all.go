package prompt

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/samsonmxvi/go-wenku-downloader/downloader"
	"github.com/samsonmxvi/go-wenku-downloader/scraper"
	"github.com/samsonmxvi/go-wenku-downloader/util"
)

func downloadAll(novelId int) error {
	downloadPath := strconv.Itoa(novelId)
	if err := util.CheckDir(downloadPath); err != nil {
		return fmt.Errorf("创建目录失败: %e", err)
	}

	// get novel metadata
	novel, err := getNovelDetails(int(novelId))
	if err != nil {
		if err == scraper.ErrorNoNovel {
			os.Remove(downloadPath)
		}
		fmt.Printf("获取小说信息失败: %e \n", err)
		return nil
	}
	promptNovelDetails(novel)

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
	// catalogueArray, err := scraper.GetCatalogueArray(novel.CatalogueUrl)
	// if err != nil {
	// 	return fmt.Errorf("下载小说卷信息失败: %e", err)
	// }
	catalogueArray, err := getCatalogueArray(novel.CatalogueUrl)
	if err != nil {
		return err
	}

	// get coverIndex from input
	converIndex := 1

	// download volume
	for _, catalogue := range catalogueArray {
		volumePath := path.Join(downloadPath, formatFilename(catalogue.Volume.Name))
		// if _, err := os.Stat(volumePath); !os.IsNotExist(err) {
		// 	// return nil
		// 	continue
		// }
		updated, err := downloader.DownloadVolume(catalogue, volumePath, true)
		if err != nil {
			log.Printf("download volume error %v", err)
			return fmt.Errorf("下载小说卷失败: %e", err)
		}
		if updated {
			err = createEpub(novel, catalogue.Volume.Name, catalogue.Volume.ChapterCount, converIndex, downloadPath)
			if err != nil {
				log.Printf("create epub failed: %v", err)
				return fmt.Errorf("下载小说卷失败: %e", err)
			}
		}
	}

	return nil
}

func getCatalogueArray(catalogueUrl string) (catalogueArray []*scraper.Catalogue, err error) {
	for i := 0; i < 3; i++ {
		catalogueArray, err = scraper.GetCatalogueArray(catalogueUrl)
		if err == nil {
			time.Sleep(time.Second)
			return catalogueArray, nil
		} else {
			time.Sleep(6 * time.Second) // temp fix rate limit
			continue
		}
	}
	return nil, fmt.Errorf("下载小说卷信息失败: %e", err)
}
