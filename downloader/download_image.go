package downloader

import (
	"fmt"
	"path"

	"github.com/samsonmxvi/go-wenku-downloader/util"
)

func DownloadImage(url string, dirPath string) error {
	// image name
	imgName := util.GetUrlLastString(url)
	// image file path
	filePath := path.Join(dirPath, imgName)

	// check file already download
	if util.CheckFileExist(filePath) {
		return nil
	}

	// check dir if not exit create
	if err := util.CheckDir(dirPath); err != nil {
		return fmt.Errorf("创建路径失败：%v", err)
	}

	return Grab(filePath, url)
}
