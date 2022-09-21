// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package statement

import "github.com/gowizzard/compver/v3/version_core"

// Core is to get the version number, split it to core block and return them.
func Core(version, block string) (int, error) {

	core, err := version_core.Split(version)
	if err != nil {
		return 0, err
	}

	var number int
	switch block {
	case "major":
		number = core.Major
	case "minor":
		number = core.Minor
	case "patch":
		number = core.Patch
	}

	return number, nil

}
