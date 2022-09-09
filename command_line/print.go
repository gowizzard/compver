package command_line

import (
	"flag"
	"fmt"
	"os"
)

// Print is to print the formatted message
// And return an exit code to the command line
func Print(code int, format string, a ...any) {
	if flag.Lookup("test.v") == nil {
		defer os.Exit(code)
	}
	fmt.Printf(format, a...)
}
