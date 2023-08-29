package scraper

import (
	"github.com/PuerkitoBio/goquery"
)

func GetNovelVolumeArray(catalogueUrl string) ([]*Volume, error) {
	// get document
	doc, err := Get(catalogueUrl)
	if err != nil {
		return nil, err
	}
	return getVolumeArrayFromDoc(doc, catalogueUrl)
}

func getVolumeArrayFromDoc(doc *goquery.Document, catalogueUrl string) ([]*Volume, error) {
	// get Index, Name, RowNumber, CatalogueUrl, create volume
	volumeArray := make([]*Volume, 0)
	doc.Find("table td[colspan]").Each(func(i int, s *goquery.Selection) {
		index := i
		name := s.Text()
		rowNumber := s.Parent().Index()
		volume := &Volume{
			Index:        index,
			Name:         name,
			RowNumber:    rowNumber,
			CatalogueUrl: catalogueUrl,
		}
		volumeArray = append(volumeArray, volume)
	})

	// calculate EndRow from document
	rows := doc.Find("tbody").Children()
	volumesLen := len(volumeArray)
	tempEndRow := rows.Length()
	for i := volumesLen - 1; i >= 0; i-- {
		volumeArray[i].EndRow = tempEndRow
		tempEndRow = volumeArray[i].RowNumber
		volumeArray[i].ChapterCount = getChapterCount(rows, volumeArray[i].RowNumber, volumeArray[i].EndRow)
	}
	return volumeArray, nil
}

func getChapterCount(rows *goquery.Selection, start int, end int) int {
	return rows.Slice(start, end).Find("a").Length()
}
