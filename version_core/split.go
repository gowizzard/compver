package version_core

import (
	"github.com/gowizzard/compver/convert"
	"regexp"
)

var split = Split{}

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
