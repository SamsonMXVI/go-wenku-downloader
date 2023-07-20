package scraper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	doc, err := Get("https://www.wenku8.net/novel/1/1973/index.htm")
	require.NoError(t, err)
	require.NotEmpty(t, doc)
	fmt.Printf("doc.Text() %v ", doc.Text())
}
