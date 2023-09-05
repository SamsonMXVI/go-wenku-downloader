package scraper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCFCookie(t *testing.T) {
	err := GetCFCookie()
	require.NoError(t, err)
}
