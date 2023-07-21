package downloader

import "fmt"

func getImageTemplateString(imageName string) string {
	return fmt.Sprintf(`正在下载图片%s: {{counters . }} {{bar . "[" "=" ">" " " "]"}} {{percent . }} {{speed . }}`, imageName)
}

func getChapterTemplateString(volumeName string, i int) string {
	return fmt.Sprintf(`正在下载《%s》章节第%d章节: {{counters . }} {{bar . "[" "=" ">" " " "]"}} {{percent . }} {{speed . }}`, volumeName, i)
}
