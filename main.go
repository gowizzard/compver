package main

import (
	"fmt"
	"github.com/gowizzard/compver/version_core"
	"log"
	"os"
)

func main() {

	os.Args = os.Args[:1]
	if len(os.Args) >= 2 {

		version1, err := version_core.Get(os.Args[1])
		if err != nil {
			log.Fatalln(err)
		}

		version2, err := version_core.Get(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}

		var blocks []version_core.Block
		blocks = append(
			blocks, version_core.Block{
				Name:    "major",
				Number1: version1.Major,
				Number2: version2.Major,
			}, version_core.Block{
				Name:    "minor",
				Number1: version1.Minor,
				Number2: version2.Minor,
			}, version_core.Block{
				Name:    "patch",
				Number1: version1.Patch,
				Number2: version2.Patch,
			},
		)

		compare := version_core.Compare(blocks)

		fmt.Printf("%s\n", compare)

	}

	fmt.Printf("Error! Not enoght arguments.\n")

}
