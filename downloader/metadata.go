package downloader

import (
	"encoding/json"
	"os"
	"path"

	"github.com/samsonmxvi/go-wenku-downloader/scraper"
	"github.com/samsonmxvi/go-wenku-downloader/util"
)

func DownloadNovelMetadata(novel *scraper.Novel, dirPath string) error {
	filePath := path.Join(dirPath, "metadata.json")

	if err := util.CheckDir(dirPath); err != nil {
		return err
	}

	novelJson, err := json.MarshalIndent(novel, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, novelJson, 0644)
	if err != nil {
		return err
	}

	return nil
}

func DownloadVolumeMetadata(volume *scraper.Volume, dirPath string) error {
	filePath := path.Join(dirPath, "metadata.json")

	if err := util.CheckDir(dirPath); err != nil {
		return err
	}

	novelJson, err := json.MarshalIndent(volume, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, novelJson, 0644)
	if err != nil {
		return err
	}

	return nil
}
