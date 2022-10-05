// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package version_core

// Core is to save the version core as major, minor,
// patch, prerelease & buildmetadata.
type Core struct {
	Major         int
	Minor         int
	Patch         int
	PreRelease    string
	BuildMetadata string
}

// Block is to save the version core blocks as a pair.
type Block struct {
	Name    string
	Number1 int
	Number2 int
}
