package scraper

import (
	"fmt"
	"net/url"

	"github.com/samsonmxvi/go-wenku-downloader/scraper/enums"
)

func GetTop(topSoftType enums.TopSortType, page string) (*PageResult, error) {
	var topResult *PageResult = &PageResult{}

	params := url.Values{}
	params.Add("sort", enums.TopSoftTextReq[topSoftType])
	params.Add("page", page)

	doc, err := Get(fmt.Sprintf("%s?%s", TOP_URL, params.Encode()))
	if err != nil {
		return nil, err
	}

	novelArray, err := getTableGridNovel(doc)
	if err != nil {
		return nil, err
	}

	topResult.NovelArray = append(topResult.NovelArray, novelArray...)
	topResult.TotalPage = getTotalPage(doc)

	return topResult, nil
}
