package downloader

import (
	"testing"

	"github.com/samsonmxvi/go-wenku-downloader/scraper"
	"github.com/stretchr/testify/require"
)

func TestDownloadVolume(t *testing.T) {
	volumeArray, err := scraper.GetCatalogueArray("https://www.wenku8.net/novel/1/1973/index.htm")
	require.NoError(t, err)
	require.NotEmpty(t, volumeArray)
	_, err = DownloadVolume(volumeArray[0], "./", true)
	require.NoError(t, err)
}
