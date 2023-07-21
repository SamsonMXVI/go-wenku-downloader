package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/samsonmxvi/go-wenku-downloader/util"
	"gopkg.in/cheggaaa/pb.v2"
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

	progressBar := pb.New(0)
	defer progressBar.Finish()
	progressBar.SetTemplateString(getImageTemplateString(imgName))

	// get image
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("获取图片失败：%v", err)
	}
	defer response.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("创建文件失败 %v", err)
	}
	defer file.Close()

	progressBar.SetTotal(int64(response.ContentLength))
	progressBar.Start()
	progressWriter := &util.ProgressWriter{
		Writer:       file,
		ProgressBar:  progressBar,
		CurrentBytes: 0,
	}

	_, err = io.Copy(progressWriter, response.Body)
	if err != nil {
		return err
	}

	return nil
}
