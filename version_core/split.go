package version_core

import (
	"github.com/gowizzard/compver/convert"
	"regexp"
)

// split is to save the version core blocks in a struct
var core = Core{
	Major: 0,
	Minor: 0,
	Patch: 0,
}

// Split is to check the given version as a string
// The function checks the string with regex
// And split the value in major, minor and patch
func Split(version string) (Core, error) {

	regex := regexp.MustCompile(`[*]|\d+`)

	find := regex.FindAllString(version, 3)
	for index, value := range find {
		switch index {
		case 0:
			core.Major = convert.Integer(value)
		case 1:
			core.Minor = convert.Integer(value)
		case 2:
			core.Patch = convert.Integer(value)
		}
	}

	return core, nil

}
