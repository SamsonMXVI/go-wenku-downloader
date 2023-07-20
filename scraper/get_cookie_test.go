package scraper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCookie(t *testing.T) {
	err := GetCookie()
	require.NoError(t, err)
}
