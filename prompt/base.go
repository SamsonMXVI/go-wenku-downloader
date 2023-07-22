package prompt

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
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
	prompt := promptui.Select{
		Label: "你打算做什么",
		Items: QuestionsText,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "> {{ . | cyan }}",
			Inactive: "  {{ . | white }}",
			Selected: "{{ . | green }}",
		},
	}

	for {
		_, question, err := prompt.Run()
		if err != nil {
			log.Fatal(err)
		}
		questionTwo(question)
	}
}
func questionTwo(question string) {
	switch question {
	case QuestionsText[ViewPopularNovels]:
		selectedIndex, err := getSelectedIndex("请选择分类", enums.TopSoftText)
		if err != nil {
			return
		}
		promptTopList(enums.TopSortType(selectedIndex))
	case QuestionsText[SearchNovels]:
		selectedIndex, err := getSelectedIndex("请选择搜索类型", enums.SearchTypeText)
		if err != nil {
			return
		}
		str, err := getInputString(fmt.Sprintf("请输入要搜索的%s", enums.SearchTypeText[selectedIndex]))
		if err != nil {
			return
		}
		searchNovels(str, enums.SearchType(selectedIndex))
	case QuestionsText[DownloadNovel]:
		novelId, err := inputNovelId()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		download(novelId)
	case QuestionsText[DoNothing]:
		os.Exit(1)
	default:
		fmt.Println()
	}
}
