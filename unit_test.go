package goToYourMenu

import (
	"fmt"
	"testing"
)

func tcommand() {
	fmt.Println("command executed")
}

func TestDrawMenu(t *testing.T) {
	options := []menuOption{
		{
			name:        "add",
			description: "adddesc",
			command:     tcommand,
		},
		{
			name:        "view",
			description: "viewdesc",
			command:     tcommand,
		},
		{
			name:        "exit",
			description: "exitdesc",
			command:     tcommand,
		},
	}
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
