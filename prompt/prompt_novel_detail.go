package prompt

import (
	"encoding/json"
	"os"
	"path"
	"strconv"

	"github.com/fatih/color"
	"github.com/samsonmxvi/go-wenku-downloader/scraper"
	"github.com/samsonmxvi/go-wenku-downloader/util"
)

func getNovelDetails(novelId int) (novel *scraper.Novel, err error) {
	novel = &scraper.Novel{}
	filePath := path.Join(strconv.Itoa(novelId), "metadata.json")
	if util.CheckFileExist(filePath) {
		getChapterContentFromFile(filePath, novel)
	} else {
		novel, err = scraper.GetNovelDetails(novelId)
		if err != nil {
			return nil, err
		}
	}
	return novel, nil
}

func promptNovelDetails(novel *scraper.Novel) {
	c := color.New(color.FgCyan)
	c.Printf("名称：%v\n", novel.NovelName)
	c.Printf("简介：%v\n", novel.Desc)
	c.Printf("作者: %v\n", novel.Author)
	c.Printf("标签: %v\n", novel.Tag)
	c.Printf("完结状态：%v\n", novel.Status)
	c.Printf("最新章节：%v\n", novel.RecentChapter)
	c.Printf("全文长度: %v\n", novel.Length)
	c.Printf("上次更新时间：%v\n", novel.LastUpdateTime)
}

func getChapterContentFromFile(path string, novel *scraper.Novel) {
	file, _ := os.ReadFile(path)
	_ = json.Unmarshal([]byte(file), novel)
}
