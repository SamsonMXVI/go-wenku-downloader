package scraper

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func getTableGridNovel(doc *goquery.Document) ([]*Novel, error) {
	var err error
	var novelArray []*Novel

	doc.Find("table.grid td").Children().Each(func(i int, divEle *goquery.Selection) {
		divHtml, mErr := divEle.Html()
		if err != nil {
			err = mErr
			return
		}

		var (
			novelName    string
			author       string
			tag          string
			catalogueUrl string
		)

		novelLinkRe := regexp.MustCompile(`href="(https://www\.wenku8\.net/book/\d+\.htm)"`)
		novelLinkMatch := novelLinkRe.FindStringSubmatch(divHtml)
		if len(novelLinkMatch) > 1 {
			catalogueUrl = novelLinkMatch[1]
		}

		novelNameRe := regexp.MustCompile(`title="(.*?)"`)
		novelNameMatch := novelNameRe.FindStringSubmatch(divHtml)
		if len(novelNameMatch) > 1 {
			novelName = novelNameMatch[1]
		}

		authorRe := regexp.MustCompile(`作者:(.*?)/`)
		authorMatch := authorRe.FindStringSubmatch(divHtml)
		if len(authorMatch) > 1 {
			author = authorMatch[1]
		}

		tagsRe := regexp.MustCompile(`Tags:<span.*?>(.*?)</span>`)
		tagsMatch := tagsRe.FindStringSubmatch(divHtml)
		if len(tagsMatch) > 1 {
			tag = tagsMatch[1]
		}

		novel := &Novel{
			NovelName:    novelName,
			Author:       author,
			Tag:          tag,
			CatalogueUrl: catalogueUrl,
		}

		novelArray = append(novelArray, novel)
	})

	if err != nil {
		return nil, err
	}

	return novelArray, nil
}

func getTotalPage(doc *goquery.Document) string {
	return doc.Find("#pagelink .last").Text()
}
