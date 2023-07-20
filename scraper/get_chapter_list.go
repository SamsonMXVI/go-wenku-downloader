package scraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetChapterList(catalogueUrl string, volume *Volume) ([]*Chapter, error) {
	doc, err := Get(catalogueUrl)
	if err != nil {
		return nil, err
	}
	chapterList := make([]*Chapter, 0)
	rows := doc.Find("tbody").Children()
	insertMap(volume.RowNumber, volume.EndRow, volume.Name, rows, &chapterList, volume.CatalogueUrl)
	return chapterList, nil
}

func insertMap(start int, end int, volumeName string, rows *goquery.Selection, chapterList *[]*Chapter, catalogueUrl string) {
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
		*chapterList = append(*chapterList, chapter)
	})
}
