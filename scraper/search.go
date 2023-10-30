package scraper

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/samsonmxvi/go-wenku-downloader/scraper/enums"
	"github.com/samsonmxvi/go-wenku-downloader/util"
)

func Search(str string, searchType enums.SearchType, page string) (*PageResult, error) {
	var searchResult *PageResult = &PageResult{}
	searchKeyByte, err := util.Utf8ToGbk(str)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("searchtype", enums.SearchTypeTextReq[searchType])
	params.Add("searchkey", string(searchKeyByte))
	params.Add("page", page)

	doc, err := Get(fmt.Sprintf("%s?%s", SEARCH_URL, params.Encode()))
	if err != nil {
		return nil, err
	}

	if docText := doc.Text(); !strings.Contains(docText, "搜索结果") {
		novelId := 0
		re := regexp.MustCompile(`/modules/article/uservote\.php\?id=(\d+)`)
		docHtml, _ := doc.Html()
		match := re.FindStringSubmatch(docHtml)
		if len(match) > 1 {
			novelId, err = strconv.Atoi(match[1])
			if err != nil {
				return nil, err
			}
		}

		novel := &Novel{
			NovelId: novelId,
		}
		getNovelDetailsDoc(doc, novel)

		searchResult.NovelArray = append(searchResult.NovelArray, novel)

		return searchResult, err
	}

	novelArray, err := getTableGridNovel(doc)
	if err != nil {
		return nil, err
	}

	searchResult.NovelArray = append(searchResult.NovelArray, novelArray...)
	searchResult.TotalPage = getTotalPage(doc)

	return searchResult, nil
}
