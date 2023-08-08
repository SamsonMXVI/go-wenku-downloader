package prompt

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
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
	novel, err := scraper.GetNovelDetails(novelId)
	if err != nil {
		return fmt.Errorf("获取小说信息失败: %e", err)
	} else if novel == nil {
		return nil
	}

	// download novel metadata and coverImg
	downloader.DownloadNovelMetadata(novel, downloadPath)
	time.Sleep(1 * time.Second)

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
	time.Sleep(1 * time.Second)

	// get selected volume list
	volumeArray, err := scraper.GetNovelVolumeArray(novel.CatalogueUrl)
	if err != nil {
		return fmt.Errorf("下载小说卷信息失败: %e", err)
	}

	time.Sleep(1 * time.Second)

	// get coverIndex from input
	converIndex := 1

	// download volume
	for _, volume := range volumeArray {
		time.Sleep(3 * time.Second)
		volumeName := strings.ReplaceAll(volume.Name, "/", "")
		volumePath := path.Join(downloadPath, volumeName)
		if _, err := os.Stat(volumePath); !os.IsNotExist(err) {
			// return nil
			continue
		}
		err = downloader.DownloadVolumeAll(volume, volumePath)
		if err != nil {
			log.Printf("download volume error %v", err)
			return fmt.Errorf("下载小说卷失败: %e", err)
		}
		err = createEpub(novel, volume.Name, volume.ChapterCount, converIndex, downloadPath)
		if err != nil {
			log.Printf("create epub failed: %v", err)
			return fmt.Errorf("下载小说卷失败: %e", err)
		}
	}

	return nil
}
