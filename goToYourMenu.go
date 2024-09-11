package goToYourMenu

import "fmt"

type menuOption struct {
	name        string
	description string
	command     func()
}

func DrawMenu(options []menuOption, currentIdx int) error {
	for idx, option := range options {
		fmt.Println(idx, option)
	}
	return nil
}

func Menu(options []menuOption) string {
	return ""
}
