package version_core

import (
	"github.com/gowizzard/compver/convert"
	"regexp"
)

var split = Split{}

func Get(version string) (Split, error) {

	regex := regexp.MustCompile("(,)?([*]|\\d+)")

	find := regex.FindAllString(version, 3)
	if len(find) >= 3 {
		split.Major = convert.Integer(find[0])
		split.Minor = convert.Integer(find[1])
		split.Patch = convert.Integer(find[2])
	}

	return split, nil

}
