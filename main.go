package main

import (
	"fmt"
	"log"
	"os"

	"github.com/samsonmxvi/go-wenku-downloader/prompt"
	"github.com/samsonmxvi/go-wenku-downloader/scraper"
	"github.com/urfave/cli"
)

func main() {
	err := scraper.GetCookie()
	if err != nil {
		log.Fatalf("登陆失败 %v \n", err)
		fmt.Println("未登录-(查询/热门小说)功能将无法使用")
	}
	app := &cli.App{
		Name:    "Go轻小说文库下载器",
		Usage:   "在终端实现轻小说的下载",
		Version: "v1.0",
		Action: func(c *cli.Context) error {
			fmt.Println("欢迎使用轻小说文库下载器，本工具源码链接如下：https://github.com/SamsonMXVI/go-wenku-downloader")
			prompt.InitPrompt()
			return nil
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
