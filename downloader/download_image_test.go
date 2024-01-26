package downloader

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadImage(t *testing.T) {
	/*
		http://pic.wenku8.com/pictures/1/1973/75978/90772.jpg http://pic.wenku8.com/pictures/1/1973/75978/90773.jpg http://pic.wenku8.com/pictures/1/1973/75978/90774.jpg http://pic.wenku8.com/pictures/1/1973/75978/90775.jpg http://pic.wenku8.com/pictures/1/1973/75978/90776.jpg http://pic.wenku8.com/pictures/1/1973/75978/90777.jpg http://pic.wenku8.com/pictures/1/1973/75978/90778.jpg http://pic.wenku8.com/pictures/1/1973/75978/90779.jpg http://pic.wenku8.com/pictures/1/1973/75978/90780.jpg http://pic.wenku8.com/pictures/1/1973/75978/90781.jpg http://pic.wenku8.com/pictures/1/1973/75978/90782.jpg http://pic.wenku8.com/pictures/1/1973/75978/90783.jpg http://pic.wenku8.com/pictures/1/1973/75978/90784.jpg http://pic.wenku8.com/pictures/1/1973/75978/90785.jpg http://pic.wenku8.com/pictures/1/1973/75978/90786.jpg http://pic.wenku8.com/pictures/1/1973/75978/90787.jpg http://pic.wenku8.com/pictures/1/1973/75978/90788.jpg http://pic.wenku8.com/pictures/1/1973/75978/90789.jpg http://pic.wenku8.com/pictures/1/1973/75978/90790.jpg http://pic.wenku8.com/pictures/1/1973/75978/90791.jpg http://pic.wenku8.com/pictures/1/1973/75978/90792.jpg
	*/
	err := DownloadImage("http://pic.wenku8.com/pictures/1/1973/75978/90774.jpg", "./test/插图/")
	require.NoError(t, err)
}
