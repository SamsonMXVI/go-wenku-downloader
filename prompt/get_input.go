package prompt

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func inputNovelId() (int, error) {
	var novelId int
	// validate input type number
	validate := func(input string) error {
		mNovelId, err := strconv.Atoi(input)

		re := regexp.MustCompile(`id=(\d+)`)
		match := re.FindStringSubmatch(input)
		if len(match) > 1 {
			mNovelId, err = strconv.Atoi(match[1])
		}

		re = regexp.MustCompile(`(\d+)\.htm`)
		match = re.FindStringSubmatch(input)
		if len(match) > 1 {
			mNovelId, err = strconv.Atoi(match[1])
		}

		if err != nil {
			return errors.New("invalid number or url")
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
		Stdout:   &noBellStdout{},
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
		Stdout:   &noBellStdout{},
	}

	_, err := prompt.Run()

	if err != nil {
		return 0, err
	}

	return coverIndex, nil
}

func getInputString(label string) (string, error) {
	prompt := promptui.Prompt{
		Label:  label,
		Stdout: &noBellStdout{},
	}

	res, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return res, nil
}

func getSelectedIndex(label string, itmes []string) (int, error) {
	selectedIndex := 0
	prompt := &survey.Select{
		Message: label,
		Options: itmes,
	}
	err := survey.AskOne(prompt, &selectedIndex)
	if err != nil {
		return 0, err
	}
	return selectedIndex, nil
}
