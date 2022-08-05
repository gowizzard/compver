package version_core_test

import (
	"github.com/gowizzard/compver/version_core"
	"log"
	"testing"
)

func TestSplit(t *testing.T) {

	version := "v1.2.4"

	core, err := version_core.Split(version)
	if err != nil {
		log.Fatalln(err)
	}

	t.Logf("The split function returns the following values. For the major: \"%d\", the minor: \"%d\" and the patch: \"%d\".\n", core.Major, core.Minor, core.Patch)

}
