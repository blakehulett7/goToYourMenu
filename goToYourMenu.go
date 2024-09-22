package goToYourMenu

import (
	"fmt"
	"os"
	"os/exec"
)

const cyan = "\033[36m"
const lightGreen = "\033[38;5:10m"
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

func MoveCursorDown(lines int) {
	fmt.Println(fmt.Sprintf("\033[%vB", lines))
}

func RunWorkaround(command exec.Cmd) {
	command.Stdout = os.Stdout
	command.Run()
}

func DrawMenu(options []MenuOption, currentIdx int) error {
	if currentIdx < 0 || currentIdx > len(options)-1 {
		return errIndexOutofRange
	}
	for idx, option := range options {
		if idx == currentIdx {
			fmt.Println(" " + lightGreen + "> " + option.Name + reset)
			continue
		}
		fmt.Println("  ", option.Name)
	}
	return nil
}

func GetUserInput() string {
	hideCliEcho := exec.Command("stty", "-F", "/dev/tty", "-echo")
	hideCliCursor := exec.Command("tput", "civis")
	showCliEcho := exec.Command("stty", "-F", "/dev/tty", "echo")
	showCliCursor := exec.Command("tput", "cnorm")
	hideCliEcho.Run()
	RunWorkaround(*hideCliCursor)
	defer showCliEcho.Run()
	defer RunWorkaround(*showCliCursor)
	readNextKeyPress := exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1")
	normalPrompt := exec.Command("stty", "-F", "/dev/tty", "-cbreak")
	readNextKeyPress.Run()
	defer normalPrompt.Run()
	var input []byte = make([]byte, 1)
	for {
		os.Stdin.Read(input)
		//fmt.Println("I got the byte", input, "("+string(input)+")")
		pressedKey := string(input)
		return pressedKey
	}
}

func WrapIndex(options []MenuOption, currentIndex int) int {
	arrayLength := len(options)
	if currentIndex < 0 {
		return arrayLength - 1
	}
	if currentIndex > arrayLength-1 {
		return 0
	}
	return currentIndex
}

func Menu(options []MenuOption) string {
	currentIdx := 0
	for {
		DrawMenu(options, currentIdx)
		MoveCursorUp(len(options) + 1)
		pressedKey := GetUserInput()
		if pressedKey == "j" {
			currentIdx++
			currentIdx = WrapIndex(options, currentIdx)
			continue
		}
		if pressedKey == "k" {
			currentIdx--
			currentIdx = WrapIndex(options, currentIdx)
			continue
		}
		if pressedKey == "\n" {
			MoveCursorDown(len(options) - 1)
			options[currentIdx].Command()
			return options[currentIdx].Name
		}
		return ""
	}
}
