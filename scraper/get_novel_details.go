package scraper

import (
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
				novel.Length = getTextAfterColon(text)
			}
		}
	})

	// Get cover
	novel.Cover = doc.Find("#content img").AttrOr("src", "")

	// Get tag
	novel.Tag = getTextAfterColon(doc.Find("#content span b").Eq(2).Text())

	// Get recentChapter
	novel.RecentChapter = doc.Find("#content span").Eq(5).Text()

	// Get desc
	novel.Desc = doc.Find("#content span").Eq(7).Text()

	// Get catalogueUrl
	novel.CatalogueUrl = doc.Find("#content").Children().First().Children().Eq(5).Children().First().Children().First().Children().First().Children().Eq(1).Children().First().AttrOr("href", "")

	return novel, nil
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
