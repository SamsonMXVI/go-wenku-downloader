package scraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetChapterArray(volume *Volume) ([]*Chapter, error) {
	doc, err := Get(volume.CatalogueUrl)
	if err != nil {
		return nil, err
	}
	chapterArray := make([]*Chapter, 0)
	rows := doc.Find("tbody").Children()
	insertMap(rows, &chapterArray, volume.RowNumber, volume.EndRow, volume.Name, volume.CatalogueUrl)
	return chapterArray, nil
}

func insertMap(rows *goquery.Selection, chapterArray *[]*Chapter, start int, end int, volumeName string, catalogueUrl string) {
	rows.Slice(start, end).Find("a").Each(func(i int, s *goquery.Selection) {
		chapterIndex := i + 1
		chapterTitle := s.Text()
		chapterUrl, _ := s.Attr("href")
		chapterUrl = strings.ReplaceAll(catalogueUrl, "index.htm", chapterUrl)
		chapter := &Chapter{
			Index: chapterIndex,
			Title: chapterTitle,
			Url:   chapterUrl,
		}
		*chapterArray = append(*chapterArray, chapter)
	})
}
