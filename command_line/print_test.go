package command_line_test

import (
	"errors"
	"github.com/gowizzard/compver/v3/command_line"
	"testing"
)

// TestPrint is to test the printing function
func TestPrint(t *testing.T) {

	tests := []struct {
		code   int
		format string
		a      any
	}{
		{
			code:   0,
			format: "%s\n",
			a:      "major update",
		},
		{
			code:   1,
			format: "%s\n",
			a:      errors.New("this is a testing error"),
		},
	}

	for _, value := range tests {
		command_line.Print(value.code, value.format, value.a)
	}

}
