package statement

import (
	"github.com/gowizzard/compver/v2/version_core"
)

// Compare is to compare the versions & return the result
func Compare(version1, version2 *string) (string, error) {

	core1, err := version_core.Split(*version1)
	if err != nil {
		return "", err
	}

	core2, err := version_core.Split(*version2)
	if err != nil {
		return "", err
	}

	var blocks []version_core.Block
	blocks = append(
		blocks, version_core.Block{
			Name:    "major",
			Number1: core1.Major,
			Number2: core2.Major,
		}, version_core.Block{
			Name:    "minor",
			Number1: core1.Minor,
			Number2: core2.Minor,
		}, version_core.Block{
			Name:    "patch",
			Number1: core1.Patch,
			Number2: core2.Patch,
		},
	)

	return version_core.Compare(blocks), nil

}
