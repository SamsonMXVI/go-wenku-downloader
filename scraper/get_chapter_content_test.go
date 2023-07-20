package scraper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetChapterContent(t *testing.T) {
	// 插图 https://www.wenku8.net/novel/1/1973/75978.htm
	// 文字 https://www.wenku8.net/novel/1/1973/111020.htm
	imageChapter := &Chapter{
		Index: 0,
		Title: "",
		Url:   "https://www.wenku8.net/novel/1/1973/75978.htm",
	}
	err := GetChapterContent(imageChapter)
	require.NoError(t, err)
	require.NotEmpty(t, imageChapter.Content.Images)
	articleChapter := &Chapter{
		Index: 0,
		Title: "",
		Url:   "https://www.wenku8.net/novel/1/1973/111020.htm",
	}
	err = GetChapterContent(articleChapter)
	require.NoError(t, err)
	require.NotEmpty(t, articleChapter.Content.Article)
}
