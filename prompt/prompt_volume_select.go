package prompt

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/samsonmxvi/go-wenku-downloader/scraper"
)

func promptVolumeSelect(catalogueUrl string) ([]*scraper.Catalogue, error) {
	var volumeOptions []string
	var selectedCatalogue []*scraper.Catalogue

	catalogueArray, err := scraper.GetCatalogueArray(catalogueUrl)
	if err != nil {
		return nil, err
	}
	for _, catalogue := range catalogueArray {
		volumeOptions = append(volumeOptions, catalogue.Volume.Name)
	}
	volumeSelected := []string{}
	prompt := &survey.MultiSelect{
		Message: "请选择下载第几卷:",
		Options: volumeOptions,
	}

	err = survey.AskOne(prompt, &volumeSelected)
	if err != nil {
		return nil, err
	}

	for _, v := range volumeSelected {
		for _, catalogue := range catalogueArray {
			if v == catalogue.Volume.Name {
				selectedCatalogue = append(selectedCatalogue, catalogue)
			}
		}
	}
	return selectedCatalogue, nil
}
