package main

import (
	"flag"
	"fmt"
	"github.com/gowizzard/compver/version_core"
	"os"
)

// version1, version2 is to save the version numbers from the flags
var (
	version1 *string
	version2 *string
)

// init is to parse the versions from the flags
func init() {
	version1 = flag.String("version1", "1.0.0", "Set the first version number")
	version2 = flag.String("version2", "1.0.0", "Set the second version number")
	flag.Parse()
}

// main is to check the arguments from the cli
// Split the version core in major, minor & patch
// And check the compare of the single blocks
func main() {

	switch arguments := os.Args[1:]; {
	case len(arguments) >= 4:

		core1, err := version_core.Split(*version1)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		core2, err := version_core.Split(*version2)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
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

		compare := version_core.Compare(blocks)

		fmt.Printf("%s\n", compare)
		os.Exit(0)

	default:

		fmt.Printf("Error! Not enoght arguments.\n")
		os.Exit(1)

	}

}
