package prompt

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/samsonmxvi/go-wenku-downloader/scraper"
)

func promptVolumeSelect(catalogueUrl string) ([]*scraper.Volume, error) {
	var volumeOptions []string
	var selectedVolume []*scraper.Volume

	volumeArray, err := scraper.GetNovelVolumeArray(catalogueUrl)
	if err != nil {
		return nil, err
	}
	for _, volume := range volumeArray {
		volumeOptions = append(volumeOptions, volume.Name)
	}
	volumeSelected := []string{}
	prompt := &survey.MultiSelect{
		Message: "请选择下载第几卷:",
		Options: volumeOptions,
	}
	survey.AskOne(prompt, &volumeSelected)
	for _, v := range volumeSelected {
		for _, volume := range volumeArray {
			if v == volume.Name {
				selectedVolume = append(selectedVolume, volume)
			}
		}
	}
	return selectedVolume, nil
}
