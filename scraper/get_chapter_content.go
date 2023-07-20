package scraper

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetChapterContent(chapter *Chapter) error {
	doc, err := Get(chapter.Url)
	if err != nil {
		return err
	}
	if strings.TrimSpace(doc.Find("#contentmain span").First().Text()) == "null" {
		content := ""
		doc, err := AndroidGet(chapter.Url)
		if err != nil {
			return err
		}
		content = doc.Find("body").Text()

		content = strings.Replace(content, "&nbsp;", "", -1)
		content = strings.Replace(content, "更多精彩热门日本轻小说、动漫小说，轻小说文库(http://www.wenku8.com) 为你一网打尽！", "", -1)

		picReg := regexp.MustCompile(`http:\/\/pic\.wenku8\.com\/pictures\/[\/0-9]+.jpg`)
		picRegL := regexp.MustCompile(`http:\/\/pic\.wenku8\.com\/pictures\/[\/0-9]+.jpg\([0-9]+K\)`)
		images := picReg.FindAllString(content, -1)
		content = picRegL.ReplaceAllString(content, "")
		content = picReg.ReplaceAllString(content, "")

		chapterContent := &ChapterContent{
			Images:  images,
			Article: content,
		}
		chapter.Content = chapterContent
		return nil
	}

	content := doc.Find("#content").Text()
	content = strings.Replace(content, "本文来自 轻小说文库(http://www.wenku8.com)", "", -1)
	content = strings.Replace(content, "台版 转自 轻之国度", "", -1)
	content = strings.Replace(content, "最新最全的日本动漫轻小说 轻小说文库(http://www.wenku8.com) 为你一网打尽！", "", -1)

	images := []string{}
	doc.Find("img").Each(func(i int, imgEle *goquery.Selection) {
		src, _ := imgEle.Attr("src")
		images = append(images, src)
	})

	chapterContent := &ChapterContent{
		Images:  images,
		Article: content,
	}
	chapter.Content = chapterContent
	return nil
}
