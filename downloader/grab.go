package downloader

import (
	"fmt"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"gopkg.in/cheggaaa/pb.v2"
)

func Grab(path string, url string) error {
	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(path, url)

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

	// create and start new bar
	bar := pb.StartNew(int(resp.Size()))

Loop:
	for {
		select {
		case <-t.C:
			bar.SetCurrent(resp.BytesComplete())

		case <-resp.Done:
			bar.SetCurrent(resp.Size())
			bar.Finish()
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		return err
	}

	if resp.HTTPResponse.StatusCode != 200 {
		return fmt.Errorf("http request error")
	}

	return nil
}
