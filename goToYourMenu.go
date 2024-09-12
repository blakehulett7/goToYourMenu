package goToYourMenu

import (
	"fmt"
	"os"
	"os/exec"
)

const cyan = "\033[36m"
const reset = "\033[0m"
const clearLine = "\033[2K"

var errIndexOutofRange = fmt.Errorf("Index out of slice range")

type MenuOption struct {
	Name        string
	Description string
	Command     func()
}

func MoveCursorUp(lines int) {
	fmt.Println(fmt.Sprintf("\033[%vA", lines))
}

func DrawMenu(options []MenuOption, currentIdx int) error {
	if currentIdx < 0 || currentIdx > len(options)-1 {
		return errIndexOutofRange
	}
	for idx, option := range options {
		if idx == currentIdx {
			fmt.Println(" " + cyan + "> " + option.Name + reset)
			continue
		}
		fmt.Println("  ", option.Name)
	}
	return nil
}

func GetUserInput() string {
	hideCliEcho := exec.Command("stty", "-F", "/dev/tty", "-echo")
	hideCliEcho.Run()
	showCliEcho := exec.Command("stty", "-F", "/dev/tty", "echo")
	defer showCliEcho.Run()
	readNextKeyPress := exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1")
	readNextKeyPress.Run()
	var input []byte = make([]byte, 1)
	for {
		os.Stdin.Read(input)
		//fmt.Println("I got the byte", input, "("+string(input)+")")
		pressedKey := string(input)
		return pressedKey
	}
}

func Menu(options []MenuOption) string {
	currentIdx := 0
	for {
		DrawMenu(options, currentIdx)
		MoveCursorUp(len(options) + 1)
		pressedKey := GetUserInput()
		if pressedKey == "j" {
			currentIdx++
			continue
		}
		if pressedKey == "k" {
			currentIdx--
			continue
		}
		if pressedKey == "\n" {
			fmt.Println("pick this one")
		}
		return ""
	}
}
