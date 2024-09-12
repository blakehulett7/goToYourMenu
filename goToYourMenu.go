package goToYourMenu

import (
	"fmt"
)

const cyan = "\033[36m"
const reset = "\033[0m"
const clearLine = "\033[2K"

var errIndexOutofRange = fmt.Errorf("Index out of slice range")

type menuOption struct {
	name        string
	description string
	command     func()
}

func MoveCursorUp(lines int) {
	fmt.Println(fmt.Sprintf("\033[%vA", lines))
}

func DrawMenu(options []menuOption, currentIdx int) error {
	if currentIdx < 0 || currentIdx > len(options)-1 {
		return errIndexOutofRange
	}
	for idx, option := range options {
		if idx == currentIdx {
			fmt.Println(" " + cyan + "> " + option.name + reset)
			continue
		}
		fmt.Println("  ", option.name)
	}
	return nil
}

func GetUserInput(currentIdx int) int {

}

func Menu(options []menuOption) string {
	currentIdx := 0
	for {
		DrawMenu(options, currentIdx)
		MoveCursorUp(len(options) + 1)
		return ""
	}
}
