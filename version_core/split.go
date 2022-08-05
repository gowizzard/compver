package version_core

import (
	"github.com/gowizzard/compver/convert"
	"regexp"
)

// split is to save the version core blocks in a struct
var split = Split{}

// Get is to check the given version as a string
// The function checks the string with regex
// And split the value in major, minor and patch
func Get(version string) (Split, error) {

	regex := regexp.MustCompile("(,)?([*]|\\d+)")

	find := regex.FindAllString(version, 3)
	for index, value := range find {
		switch index {
		case 0:
			split.Major = convert.Integer(value)
		case 1:
			split.Minor = convert.Integer(value)
		case 2:
			split.Patch = convert.Integer(value)
		}
	}

	return split, nil

}
