package scraper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetChapterMap(t *testing.T) {
	Novel, err := GetNovelDetails(1973)
	require.NoError(t, err)
	require.NotEmpty(t, Novel)
	catalogueUrl := Novel.CatalogueUrl

	volumes, err := GetNovelVolumeArray(catalogueUrl)
	require.NoError(t, err)
	fmt.Println("Volume Map:")
	require.NoError(t, err)
	chaterMap, err := GetChapterArray(volumes[0])
	require.NoError(t, err)
	require.NotEmpty(t, chaterMap)
	for _, chapter := range chaterMap {
		fmt.Printf("Chapter %d: %s (%s)\n", chapter.Index, chapter.Title, chapter.Url)
	}
	fmt.Println("Total Chapters:", len(chaterMap))
}
