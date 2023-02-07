// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package build_information is to save variables from the build
// process. The variables are stored as public variables.
package build_information

import "runtime/debug"

// Version is to save the git tag from makefile.
var (
	Version string
)

// init is to add the version if the binary is installed via `go install`.
func init() {

	if len(Version) == 0 {
		info, ok := debug.ReadBuildInfo()
		if ok {
			switch info.Main.Version {
			case "(devel)":
				Version = "unavailable"
			default:
				Version = info.Main.Version
			}
		}
	}

}
