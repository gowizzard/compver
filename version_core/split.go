package version_core

import (
	"github.com/gowizzard/compver/convert"
	"regexp"
)

var body = Body{}

func Get(version string) (Body, error) {

	regex := regexp.MustCompile("(\\\\,)?(\\\\.)?(\\*|\\d+)")

	find := regex.FindAllString(version, 3)
	if len(find) >= 3 {
		body.Major = convert.Integer(find[0])
		body.Minor = convert.Integer(find[1])
		body.Patch = convert.Integer(find[2])
	}

	return body, nil

}
