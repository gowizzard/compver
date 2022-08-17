package main

import (
	"fmt"
	"github.com/gowizzard/compver/version_core"
	"log"
	"os"
)

// main is to check the arguments from the cli
// Split the version core in major, minor & patch
// And check the compare of the single blocks
func main() {

	switch arguments := os.Args[1:]; {
	case len(arguments) >= 2:

		core1, err := version_core.Split(os.Args[1])
		if err != nil {
			log.Fatalln(err)
		}

		core2, err := version_core.Split(os.Args[2])
		if err != nil {
			log.Fatalln(err)
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

	default:

		fmt.Printf("Error! Not enoght arguments.\n")

	}

}
