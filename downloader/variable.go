package downloader

import "time"

var (
	ErrorJsonName   = "error.json"
	ImageFolderName = "images"
	DownloadTimer   = time.Second
	RetryTimes      = 6
	RetryTimer      = 6 * time.Second
)
