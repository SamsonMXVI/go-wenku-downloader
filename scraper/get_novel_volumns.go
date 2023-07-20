package scraper

import "github.com/PuerkitoBio/goquery"

func GetNovelVolumns(catalogueUrl string) ([]*Volume, error) {
	// get document
	doc, err := Get(catalogueUrl)
	if err != nil {
		return nil, err
	}

	// get Index, Name, RowNumber, CatalogueUrl, create volume
	volumes := make([]*Volume, 0)
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
		volumes = append(volumes, volume)
	})

	// calculate EndRow from document
	rows := doc.Find("tbody").Children()
	volumesLen := len(volumes)
	tempEndRow := rows.Length()
	for i := volumesLen - 1; i >= 0; i-- {
		volumes[i].EndRow = tempEndRow
		tempEndRow = volumes[i].RowNumber
	}
	return volumes, nil
}
