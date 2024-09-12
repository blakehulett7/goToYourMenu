package goToYourMenu

import (
	"fmt"
	"testing"
)

var options = []MenuOption{
	{
		Name:        "add",
		Description: "adddesc",
		Command:     tcommand,
	},
	{
		Name:        "view",
		Description: "viewdesc",
		Command:     tcommand,
	},
	{
		Name:        "exit",
		Description: "exitdesc",
		Command:     tcommand,
	},
}

func tcommand() {
	fmt.Println("command executed")
}

func TestDrawMenu(t *testing.T) {
	tests := map[string]struct {
		options    []menuOption
		currentIdx int
		want       error
	}{
		"simple": {
			options:    options,
			currentIdx: 1,
			want:       nil,
		},
		"idx too low": {
			options:    options,
			currentIdx: -1,
			want:       errIndexOutofRange,
		},
		"idx too high": {
			options:    options,
			currentIdx: 3,
			want:       errIndexOutofRange,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := DrawMenu(test.options, test.currentIdx)
			if got != test.want {
				t.Fatalf("test failed: expected %v, got %v", test.want, got)
			}
		})
	}
}

func TestGetUserInput(t *testing.T) {
	fmt.Println(GetUserInput())
}
