package scraper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCookie(t *testing.T) {
	err := GetCFCookie()
	require.NoError(t, err)
	err = GetCookie()
	require.NoError(t, err)
	fmt.Printf("format string %v \n", Cookie)
}
