package scraper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetNovelDetails retrieves the novel details based on novelId
func GetNovelDetails(novelId int) (*Novel, error) {

	novel := &Novel{
		NovelId: novelId,
	}

	doc, err := Get(BASE_URL + strconv.Itoa(novelId) + ".htm")
	if err != nil {
		return nil, err
	}

	if doc.Find(".blocktitle").First().Text() == "出现错误！" {
		return nil, fmt.Errorf("没有这本小说")
	}

	getNovelDetailsDoc(doc, novel)

	return novel, nil
}

func getNovelDetailsDoc(doc *goquery.Document, novel *Novel) {
	mEqCopyright := 0
	mEqAnimate := 0

	containsCopyright := strings.Contains(doc.Text(), "因版权问题，文库不再提供该小说的在线阅读与下载服务！")
	containsAnimate := strings.Contains(doc.Text(), "本作已动画化")

	if containsCopyright {
		mEqCopyright = 2
	}
	if !containsAnimate {
		mEqAnimate = 1
	}

	// Get novelName
	novel.NovelName = doc.Find("#content span b").First().Text()

	// Get library, author, status, lastUpdateTime, length
	info := doc.Find("#content").Children().First().Children().First().Children().First().Children()
	info.Children().Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if text != "" {
			text = strings.TrimSpace(text)
			switch i {
			case 1:
				novel.Library = getTextAfterColon(text)
			case 2:
				novel.Author = getTextAfterColon(text)
			case 3:
				novel.Status = getTextAfterColon(text)
			case 4:
				novel.LastUpdateTime = getTextAfterColon(text)
			case 5:
				if !containsCopyright {
					novel.Length = getTextAfterColon(text)
				}
			}
		}
	})

	// Get cover
	novel.Cover = doc.Find("#content img").AttrOr("src", "")

	// Get tag
	novel.Tag = getTextAfterColon(doc.Find("#content span b").Eq(2 - mEqAnimate).Text())

	// Get recentChapter
	if !containsCopyright {
		novel.RecentChapter = doc.Find("#content span").Eq(5 - mEqCopyright - mEqAnimate).Text()
	}

	// Get desc
	novel.Desc = doc.Find("#content span").Eq(7 - mEqCopyright - mEqAnimate).Text()

	// Get catalogueUrl
	html, _ := doc.Find("#content").Children().First().Html()
	re := regexp.MustCompile(`<a href="(https://www\.wenku8\.net/novel/[^"]+)">小说目录</a>`)
	match := re.FindStringSubmatch(html)
	if match != nil {
		novel.CatalogueUrl = match[1]
	}
}

// Utility function to extract text after colon in a string
func getTextAfterColon(text string) string {
	re := regexp.MustCompile(`：(.+)$`)
	matches := re.FindStringSubmatch(text)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
