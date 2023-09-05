package scraper

import (
	"context"
	"fmt"

	cu "github.com/Davincible/chromedp-undetected"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func GetCFCookie() error {
	ctx, cancel, err := cu.New(cu.NewConfig(
	// Remove this if you want to see a browser window.
	// cu.WithHeadless(),
	))
	if err != nil {
		return err
	}
	defer cancel()

	// var outerHTML string
	var cookies []*network.Cookie
	if err != nil {
		return err
	}
	chromedp.UserAgent(UserAgent)
	err = chromedp.Run(ctx,
		network.Enable(),
		chromedp.Navigate("https://www.wenku8.net/book/1973.htm"),
		chromedp.WaitVisible(`#content > div:nth-child(1) > table:nth-child(1) > tbody > tr:nth-child(1) > td > table > tbody > tr > td:nth-child(1) > span > b`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 获取 Cookie
			c, err := network.GetCookies().Do(ctx)
			if err != nil {
				return err
			}
			cookies = c
			return nil
		}),
	)
	if err != nil {
		return err
	}

	// 输出结果
	for _, cookie := range cookies {
		// fmt.Printf("- %s: %s\n", cookie.Name, cookie.Value)
		Cookie = fmt.Sprintf("%s=%s;%s", cookie.Name, cookie.Value, Cookie)
	}
	return nil
}
