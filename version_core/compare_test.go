package version_core_test

import (
	"github.com/gowizzard/compver/version_core"
	"testing"
)

// TestCompare is to check the version core compare function
// We simulate the major, minor & patch values
// And log the returned value
func TestCompare(t *testing.T) {

	majorNumber1 := 1
	majorNumber2 := 1

	minorNumber1 := 3
	minorNumber2 := 2

	patchNumber1 := 0
	patchNumber2 := 5

	var blocks []version_core.Block
	blocks = append(
		blocks, version_core.Block{
			Name:    "major",
			Number1: majorNumber1,
			Number2: majorNumber2,
		}, version_core.Block{
			Name:    "minor",
			Number1: minorNumber1,
			Number2: minorNumber2,
		}, version_core.Block{
			Name:    "patch",
			Number1: patchNumber1,
			Number2: patchNumber2,
		},
	)

	compare := version_core.Compare(blocks)
	t.Logf("The compare function returns the following value: \"%s\".\n", compare)

}
