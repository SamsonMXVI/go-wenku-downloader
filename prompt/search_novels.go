package prompt

import (
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"

	"github.com/samsonmxvi/go-wenku-downloader/scraper"
	"github.com/samsonmxvi/go-wenku-downloader/scraper/enums"
)

func searchNovels(searchText string, searchType enums.SearchType) error {
	var pageIndex int = 1
	var items []string
	var c *color.Color
	var novelId int = 0

	prompt := survey.Select{
		PageSize: 9,
	}

	for {
		items = []string{}
		var totalPage int
		var selectedIndex int

		// get search result
		searchResult, err := scraper.Search(searchText, searchType, strconv.Itoa(pageIndex))
		if err != nil {
			return err
		}

		if len(searchResult.NovelArray) == 1 && searchResult.TotalPage == "" {
			novelId = searchResult.NovelArray[0].NovelId
			break
		}

		// totalpage convert string to int
		totalPage, err = strconv.Atoi(searchResult.TotalPage)
		if err != nil {
			return err
		}

		prompt.Message = fmt.Sprintf("请选择需要下载的小说-当前页面(%d/%d)", pageIndex, totalPage)

		// generate prompt item
		for _, sR := range searchResult.NovelArray {
			str := ""
			c = color.New(color.FgGreen)
			str += c.Sprintf(sR.NovelName)
			c = color.New(color.FgRed)
			str += c.Sprintf(" 作者: %s", sR.Author)
			c = color.New(color.FgBlue)
			str += c.Sprintf(" 标签: %s", sR.Tag)
			items = append(items, str)
		}
		if pageIndex != 1 {
			items = append(items, "上一页")
		}
		if pageIndex != totalPage {
			items = append(items, "下一页")
		}
		items = append(items, "返回")

		// start prompt
		prompt.Options = items
		err = survey.AskOne(&prompt, &selectedIndex)

		if err != nil {
			return err
		}

		if prompt.Options[selectedIndex] == "返回" {
			return nil
		}
		if prompt.Options[selectedIndex] == "上一页" {
			pageIndex -= 1
			continue
		}
		if prompt.Options[selectedIndex] == "下一页" {
			pageIndex += 1
			continue
		}

		if prompt.Options[selectedIndex] != "" {
			novel := searchResult.NovelArray[selectedIndex]
			novelId, err = getNovelIdFromUrl(novel.CatalogueUrl)
			if err != nil {
				return err
			}
			break
		}
	}
	if novelId != 0 {
		err := download(novelId)
		if err != nil {
			return err
		}
	}
	return nil
}
