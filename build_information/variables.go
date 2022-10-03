// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package build_information is to save variables from the build
// process. The variables are stored as public variables.
package build_information

// Version is to save the git tag from makefile.
var (
	Version string
)
