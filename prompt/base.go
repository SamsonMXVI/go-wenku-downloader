package prompt

import (
	"fmt"
	"log"
	"os"

	"github.com/samsonmxvi/go-wenku-downloader/scraper/enums"
)

type Questions int

const (
	ViewPopularNovels Questions = iota
	SearchNovels
	DownloadNovel
	DoNothing
)

var QuestionsText = []string{
	"查看(今日更新/热门轻小说/总推荐榜/...)",
	"搜索小说",
	"下载小说",
	"什么也不做",
}

func InitPrompt() {
	for {
		selectedIndex, _ := getSelectedIndex("你打算做什么", QuestionsText)
		questionTwo(Questions(selectedIndex))
	}
}
func questionTwo(question Questions) {
	switch question {
	case ViewPopularNovels:
		selectedIndex, err := getSelectedIndex("请选择分类", enums.TopSoftText)
		if err != nil {
			log.Printf("Search failed %v\n", err)
			return
		}
		promptTopList(enums.TopSortType(selectedIndex))

	case SearchNovels:
		selectedIndex, err := getSelectedIndex("请选择搜索类型", enums.SearchTypeText)
		if err != nil {
			log.Printf("Search failed %v\n", err)
			return
		}
		str, err := getInputString(fmt.Sprintf("请输入要搜索的%s", enums.SearchTypeText[selectedIndex]))
		if err != nil {
			log.Printf("Search failed %v\n", err)
			return
		}
		err = searchNovels(str, enums.SearchType(selectedIndex))
		if err != nil {
			log.Printf("Search failed %v\n", err)
			return
		}

	case DownloadNovel:
		novelId, err := inputNovelId()
		if err != nil {
			log.Printf("Prompt failed %v\n", err)
			return
		}
		err = download(novelId)
		if err != nil {
			log.Printf("download failed %v\n", err)
			return
		}

	case DoNothing:
		os.Exit(1)
	default:
		fmt.Println()
	}
}
