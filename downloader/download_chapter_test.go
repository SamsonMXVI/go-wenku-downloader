package downloader

import (
	"testing"

	"github.com/samsonmxvi/go-wenku-downloader/scraper"
	"github.com/stretchr/testify/require"
)

func TestDownloadChapter(t *testing.T) {
	//版权  https://www.wenku8.net/novel/0/471/17514.htm
	//版权 插图 https://www.wenku8.net/novel/0/471/17513.htm
	//正常 https://www.wenku8.net/novel/1/1973/69567.htm
	//正常 插图 https://www.wenku8.net/novel/1/1973/69759.htm
	chapter := &scraper.Chapter{
		Index: 1,
		Title: "KEYWORDS",
		Url:   "https://www.wenku8.net/novel/1/1973/69567.htm",
	}
	err := scraper.GetChapterContent(chapter)
	require.NoError(t, err)
	err = DownloadChapter(chapter, "./test/第一卷")
	require.NoError(t, err)
	chapterImage := &scraper.Chapter{
		Index: 1,
		Title: "插图",
		Url:   "https://www.wenku8.net/novel/0/471/17513.htm",
	}
	err = scraper.GetChapterContent(chapterImage)
	require.NoError(t, err)
	err = DownloadChapter(chapterImage, "./test/第二卷")
	require.NoError(t, err)

}
