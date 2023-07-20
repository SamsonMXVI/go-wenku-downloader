package scraper

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNovelDetail(t *testing.T) {
	Novel, err := GetNovelDetails(1973)
	require.NoError(t, err)
	require.NotEmpty(t, Novel)
	s, _ := json.MarshalIndent(Novel, "", "\t")
	fmt.Println(string(s))
}
