package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/corpix/uarand"
	"github.com/fatih/color"
	"github.com/samsonmxvi/go-wenku-downloader/prompt"
	"github.com/samsonmxvi/go-wenku-downloader/scraper"
	"github.com/urfave/cli"
)

func main() {
	f, err := os.OpenFile("scraper.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
	}()
	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	scraper.UserAgent = uarand.GetRandom()
	err = scraper.GetCookie()
	if err != nil {
		c := color.New(color.FgRed)
		c.Printf("登陆失败 %v \n", err)
		c.Println("未登录-(查询/热门小说)功能将无法使用")
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
