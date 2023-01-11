// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package statement

import (
	"errors"
	"github.com/gowizzard/compver/v5/version_core"
)

// Core is to get the version number, split it to core block and return them.
func Core(version, block string) (any, error) {

	core, err := version_core.Split(version)
	if err != nil {
		return 0, err
	}

	var result any
	switch block {
	case "major":
		result = core.Major
	case "minor":
		result = core.Minor
	case "patch":
		result = core.Patch
	case "prerelease":
		result = core.PreRelease
	case "buildmetadata":
		result = core.BuildMetadata
	default:
		return nil, errors.New("block not found")
	}

	return result, nil

}
