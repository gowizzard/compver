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
func Write(key string, value any) error {

	path := os.Getenv("GITHUB_OUTPUT")
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	output := fmt.Appendf([]byte{}, "%s=\"%v\"\n", key, value)
	_, err = file.Write(output)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil

}
