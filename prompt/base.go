package prompt

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
)

type Questions int

const (
	ViewPopularNovels Questions = iota
	SearchNovels
	DownloadNovel
	DoNothing
)

var QuestionsText = []string{
	"查看热门小说 --待实现",
	"搜索小说 --待实现",
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
		popularNovel()
	case QuestionsText[SearchNovels]:
		searchNovels()
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
