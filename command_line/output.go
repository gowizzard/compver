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

// Output is to build the attribute output.
func Output(attributes map[string]any) {

	var build []byte
	for index, value := range attributes {
		build = fmt.Append(build, "::set-output")
		build = fmt.Appendf(build, " name=%s::%v\n", index, value)
	}
	fmt.Printf("%s", build)

	if flag.Lookup("test.v") == nil {
		defer os.Exit(0)
	}

}
