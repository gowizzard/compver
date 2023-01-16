// Copyright 2023 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package output

import "os"

// Read is to read the content from the github output file.
func Read() ([]byte, error) {

	path := os.Getenv("GITHUB_OUTPUT")
	read, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return read, nil

}
