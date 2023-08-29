package scraper

func GetCatalogueArray(catalogueUrl string) ([]*Catalogue, error) {
	catalogueArray := make([]*Catalogue, 0)

	doc, err := Get(catalogueUrl)
	if err != nil {
		return nil, err
	}

	volumeArray, err := getVolumeArrayFromDoc(doc, catalogueUrl)

	if err != nil {
		return nil, err
	}

	for i, volume := range volumeArray {
		catalogueArray = append(catalogueArray, &Catalogue{})
		chapterArray, err := getChapterArrayFromDoc(doc, volume)
		if err != nil {
			return nil, err
		}
		catalogueArray[i].Volume = *volume
		catalogueArray[i].ChapterArray = chapterArray
	}

	return catalogueArray, nil
}
