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
	"os/exec"
)

// Output is to build the github action output.
func Output(key string, value any) {

	output := fmt.Sprintf("\"%s=%v\" >> $GITHUB_OUTPUT", key, value)

	command := exec.Command("echo", output)
	err := command.Run()
	if err != nil {
		os.Exit(1)
	}

	if flag.Lookup("test.v") == nil {
		defer os.Exit(0)
	}

}
