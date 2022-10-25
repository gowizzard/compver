// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package version_core

import (
	"errors"
	"github.com/gowizzard/compver/v5/convert"
	"regexp"
)

// regex is to save the compiled expression.
var regex = regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

// Split is to check the given version as a string. The function checks
// the string with regex and split the value in major, minor, patch, prerelease & buildmetadata.
func Split(version string) (Core, error) {

	core := Core{
		Major:         0,
		Minor:         0,
		Patch:         0,
		PreRelease:    "",
		BuildMetadata: "",
	}

	if !regex.Match([]byte(version)) {
		return core, errors.New("it is not a semantic version number")
	}

	match := regex.FindStringSubmatch(version)
	for index, value := range regex.SubexpNames() {
		switch value {
		case "major":
			core.Major = convert.Integer(match[index])
		case "minor":
			core.Minor = convert.Integer(match[index])
		case "patch":
			core.Patch = convert.Integer(match[index])
		case "prerelease":
			core.PreRelease = match[index]
		case "buildmetadata":
			core.BuildMetadata = match[index]
		}
	}

	return core, nil

}
