package prompt

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func inputNovelId() (int64, error) {
	var novelId int64
	// validate input type number
	validate := func(input string) error {
		mNovelId, err := strconv.ParseInt(input, 10, 64)

		re := regexp.MustCompile(`id=(\d+)`)
		match := re.FindStringSubmatch(input)
		if len(match) > 1 {
			mNovelId, err = strconv.ParseInt(match[1], 10, 64)
		}

		re = regexp.MustCompile(`(\d+)\.htm`)
		match = re.FindStringSubmatch(input)
		if len(match) > 1 {
			mNovelId, err = strconv.ParseInt(match[1], 10, 64)
		}

		if err != nil {
			return errors.New("invalid number or url")
		}

		novelId = mNovelId
		return nil
	}

	c := color.New(color.FgGreen)
	c.Println("支持格式: ")
	c.Println("	数字ID: 1973")
	c.Println("	url格式 1: https://www.wenku8.net/book/2975.htm")
	c.Println("	url格式 2: https://www.wenku8.net/modules/article/articleinfo.php?id=14&char=&charset=big5")

	prompt := promptui.Prompt{
		Label:    "输入小说ID",
		Validate: validate,
	}

	_, err := prompt.Run()

	if err != nil {
		return 0, err
	}

	return novelId, nil
}

func inputCoverIndex() (int, error) {
	var coverIndex int
	// validate input type number
	validate := func(input string) error {
		if input == "" {
			return nil
		}
		mCoverIndex, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("invalid number")
		}
		coverIndex = mCoverIndex
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "输入第几张插图作为封面(默认:0, 使用小说封面)",
		Validate: validate,
	}

	_, err := prompt.Run()

	if err != nil {
		return 0, err
	}

	return coverIndex, nil
}
