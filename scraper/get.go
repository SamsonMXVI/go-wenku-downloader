package scraper

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/samsonmxvi/go-wenku-downloader/util"
)

func Get(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// set cookie header
	req.Header.Set("Cookie", Cookie)
	req.Header.Add("User-Agent", UserAgent)

	// send request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http request error")
	}

	// get response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// transcoding body from gbk to utf8
	decodedBody, err := util.GbkToUtf8(body)
	if err != nil {
		return nil, err
	}

	// create goquery.Document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(decodedBody)))
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func AndroidGet(url string) (*goquery.Document, error) {
	v := strings.Split(url[:strings.LastIndex(url, ".")], "/")
	bookId := v[len(v)-2]
	chapterId := v[len(v)-1]
	payload := strings.NewReader(encode(bookId, chapterId))
	req, err := http.NewRequest("POST", "http://app.wenku8.com/android.php", payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Dalvik/2.1.0 (Linux; U; Android 7.1.2; unknown Build/NZH54D)")

	// set cookie header
	req.Header.Set("Cookie", Cookie)

	// send request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http request error")
	}

	// get response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// create goquery.Document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func encode(bookId string, chapterId string, lang_optional ...int) string {
	lang := 0
	if len(lang_optional) > 0 {
		lang = lang_optional[0]
	}
	str := fmt.Sprintf("action=book&do=text&aid=%s&cid=%s&t=%d", bookId, chapterId, lang)
	encodedStr := base64.StdEncoding.EncodeToString([]byte(str))
	timetoken := time.Now().Unix()
	return fmt.Sprintf("&appver=1.13&request=%s&timetoken=%d", encodedStr, timetoken)
}
