package statement

import "github.com/gowizzard/compver/v3/version_core"

// Core is to get a core number as result
func Core(version, block *string) (int, error) {

	core, err := version_core.Split(*version)
	if err != nil {
		return 0, err
	}

	var number int
	switch *block {
	case "major":
		number = core.Major
	case "minor":
		number = core.Minor
	case "patch":
		number = core.Patch
	}

	return number, nil

}
