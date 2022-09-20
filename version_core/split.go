package version_core

import (
	"errors"
	"github.com/gowizzard/compver/v3/convert"
	"regexp"
)

// regex is to save the compiled expression.
var regex = regexp.MustCompile(`(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)`)

// Split is to check the given version as a string. The function checks
// the string with regex and split the value in major, minor and patch.
func Split(version string) (Core, error) {

	core := Core{
		Major: 0,
		Minor: 0,
		Patch: 0,
	}

	if !regex.Match([]byte(version)) {
		return core, errors.New("it is not a version number")
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
		}
	}

	return core, nil

}
