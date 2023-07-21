package prompt

import (
	"errors"
	"strconv"

	"github.com/manifoldco/promptui"
)

func inputNovelId() (int64, error) {
	var novelId int64
	// validate input type number
	validate := func(input string) error {
		mNovelId, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return errors.New("invalid number")
		}
		novelId = mNovelId
		return nil
	}

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
