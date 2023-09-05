package scraper

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func GetCookie() error {
	client := &http.Client{}
	loginUrl := "https://www.wenku8.net/login.php?do=submit&jumpurl=http%3A%2F%2Fwww.wenku8.net%2Findex.php"
	form := url.Values{}
	form.Add("username", "2497360927")
	form.Add("password", "testtest")
	form.Add("usecookie", "315360000")
	form.Add("action", "login")
	form.Add("submit", "%26%23160%3B%B5%C7%26%23160%3B%26%23160%3B%C2%BC%26%23160%3B")
	req, _ := http.NewRequest("POST", loginUrl, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", Cookie)
	req.Header.Add("User-Agent", UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	cookies := resp.Cookies()
	var newCookie string
	for _, cookie := range cookies {
		newCookie += cookie.Name + "=" + cookie.Value + ";"
	}
	if len(newCookie) > 0 {
		Cookie = fmt.Sprintf("%s;%s", Cookie, newCookie)
	}
	return nil
}
