package util

import (
	"fmt"
	"log"
	"strings"

	"github.com/bmaupin/go-epub"
)

func AddSection(epub *epub.Epub, title string, article string) error {
	body := strings.ReplaceAll(article, "\\n", "<br/>")
	xhtml := fmt.Sprintf(`<h1>%v</h1>%v`, title, body)
	_, err := epub.AddSection(xhtml, title, title+".xhtml", "")
	if err != nil {
		return err
	}
	return nil
}

func CreateSectionXhtml(title string, article string) string {
	body := strings.ReplaceAll(article, "\\n", "<br/>")
	xhtml := fmt.Sprintf(`<h1>%v</h1>%v`, title, body)
	return xhtml
}

func AddSectionXhtml(epub *epub.Epub, title string, xhtml string) error {
	_, err := epub.AddSection(xhtml, title, "", "")
	if err != nil {
		return err
	}
	return nil
}

func AddImage(epub *epub.Epub, filePath string) (string, error) {
	imgPath, err := epub.AddImage(filePath, GetUrlLastString(filePath))
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return imgPath, nil
}

func AddImageToXhtml(internalPath string, xhtml string) string {
	return fmt.Sprintf("%v<img src='%v'/>", xhtml, internalPath)
}
