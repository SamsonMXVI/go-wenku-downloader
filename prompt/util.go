package prompt

import (
	"regexp"
	"strconv"
)

func getNovelIdFromUrl(url string) (int, error) {
	var (
		mNovelId int
		err      error
	)

	re := regexp.MustCompile(`id=(\d+)`)
	match := re.FindStringSubmatch(url)
	if len(match) > 1 {
		mNovelId, err = strconv.Atoi(match[1])
	}

	re = regexp.MustCompile(`(\d+)\.htm`)
	match = re.FindStringSubmatch(url)
	if len(match) > 1 {
		mNovelId, err = strconv.Atoi(match[1])
	}

	return mNovelId, err
}
