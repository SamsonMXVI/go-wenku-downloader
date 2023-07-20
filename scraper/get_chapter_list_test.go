package scraper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetChapterList(t *testing.T) {
	Novel, err := GetNovelDetails(1973)
	require.NoError(t, err)
	require.NotEmpty(t, Novel)
	catalogueUrl := Novel.CatalogueUrl

	volumes, err := GetNovelVolumns(catalogueUrl)
	require.NoError(t, err)
	fmt.Println("Volume Map:")
	require.NoError(t, err)
	chaterList, err := GetChapterList(catalogueUrl, volumes[0])
	require.NoError(t, err)
	require.NotEmpty(t, chaterList)
	for _, chapter := range chaterList {
		fmt.Printf("Chapter %d: %s (%s)\n", chapter.Index, chapter.Title, chapter.Url)
	}
	fmt.Println("Total Chapters:", len(chaterList))
}
