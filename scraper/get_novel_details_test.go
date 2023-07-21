package scraper

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNovelDetail(t *testing.T) {
	novel, err := GetNovelDetails(1973)
	require.NoError(t, err)
	require.NotEmpty(t, novel)
	s, _ := json.MarshalIndent(novel, "", "\t")
	fmt.Println(string(s))

	// copyright animate
	// novel, err := GetNovelDetails(1587)
	// require.NoError(t, err)
	// require.NotEmpty(t, novel)
	// fmt.Printf("%v", novel.LastUpdateTime)

	// // animate
	// novel, err := GetNovelDetails(2975)
	// require.NoError(t, err)
	// require.NotEmpty(t, novel)
	// fmt.Printf("%v", novel.Desc)
}
