// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package output is there to provide data for the github output.
package output

import (
	"fmt"
	"os"
)

// Write is to write an environment variable to the github output file.
func Write(key string, value any, read []byte) error {

	path := os.Getenv("GITHUB_OUTPUT")
	data := fmt.Appendf(read, "%s=\"%v\"\n", key, value)

	err := os.WriteFile(path, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil

}
