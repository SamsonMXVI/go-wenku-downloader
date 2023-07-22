package scraper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTop(t *testing.T) {
	err := GetCookie()
	require.NoError(t, err)

	result, err := GetTop(TopSoftDailyLeaderboard, "1")
	require.NoError(t, err)
	require.NotEmpty(t, result.NovelArray)
	require.True(t, len(result.NovelArray) > 1)

	for _, v := range result.NovelArray {
		fmt.Printf("NovelName: %s\n", v.NovelName)
		fmt.Printf("Author: %s\n", v.Author)
		fmt.Printf("Tag: %s\n", v.Tag)
		fmt.Printf("CatalogueUrl: %s\n", v.CatalogueUrl)
	}

	fmt.Printf("total page %s\n", result.TotalPage)
}
