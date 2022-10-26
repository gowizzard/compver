// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package command_line is to print data to the
// command line interface. The package contains functions
// to ensure the output in the console and the exit it.
package command_line

import (
	"flag"
	"fmt"
	"os"
)

// Print is to print the formatted message and
// return an exit code to the command line.
func Print(code int, format string, a ...any) {
	if flag.Lookup("test.v") == nil {
		defer os.Exit(code)
	}
	fmt.Printf(format, a...)
}
