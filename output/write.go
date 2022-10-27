// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package output is there to provide data for the github output.
package output

import (
	"flag"
	"fmt"
	"github.com/gowizzard/compver/v5/command_line"
	"os"
)

// output is to define the variable for the format specifier.
var output []byte

// Write is to write an environment variable to the github output file.
func Write(key string, value any) {

	path := os.Getenv("GITHUB_OUTPUT")
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		command_line.Print(1, "the output file cannot be opened\n")
	}

	output := fmt.Appendf(output, "%s=\"%v\"\n", key, value)
	_, err = file.Write(output)
	if err != nil {
		command_line.Print(1, "the string cannot be written to the output\n")
	}

	err = file.Close()
	if err != nil {
		command_line.Print(1, "the output file cannot be closed")
	}

	if flag.Lookup("test.v") == nil {
		defer os.Exit(0)
	}

}
