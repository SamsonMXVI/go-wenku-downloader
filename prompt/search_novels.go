package prompt

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"

	"github.com/manifoldco/promptui"
	"github.com/samsonmxvi/go-wenku-downloader/scraper"
	"github.com/samsonmxvi/go-wenku-downloader/scraper/enums"
)

func searchNovels(searchText string, searchType enums.SearchType) error {
	var pageIndex int = 1
	var items []string
	var c *color.Color
	var novelId int = 0

	prompt := promptui.Select{
		Size: 9,
	}

	for {
		items = []string{}
		var totalPage int

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

		prompt.Label = fmt.Sprintf("请选择需要下载的小说-当前页面(%d/%d)", pageIndex, totalPage)

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
		prompt.Items = items
		selectedIndex, result, err := prompt.Run()
		if err != nil {
			return err
		}

		if result == "返回" {
			return nil
		}
		if result == "上一页" {
			pageIndex -= 1
			continue
		}
		if result == "下一页" {
			pageIndex += 1
			continue
		}

		if result != "" {
			novel := searchResult.NovelArray[selectedIndex]
			novelId, err = getNovelIdFromUrl(novel.CatalogueUrl)
			if err != nil {
				return err
			}
			break
		}
	}
	if novelId != 0 {
		download(novelId)
	}
	return nil
}
