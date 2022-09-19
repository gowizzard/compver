// Package main is used for reading and checking the tags also the
// different statements are called over it and the return value is passed on.
package main

import (
	"flag"
	"github.com/gowizzard/compver/v3/build_information"
	"github.com/gowizzard/compver/v3/command_line"
	"github.com/gowizzard/compver/v3/statement"
	"os"
)

// version, compare, core is to save the boolean for the compare statement
// version1, version2 is to save the version numbers from the flags
var (
	version, compare, core    bool
	version1, version2, block string
	visit                     = make(map[string]bool)
	action                    = os.Getenv("GITHUB_ACTIONS") == "true"
)

// init is to parse the versions from the flags and check all visited flags.
func init() {

	flag.BoolVar(&version, "version", false, "Get the current version")
	flag.BoolVar(&compare, "compare", false, "Set the statement to compare the version numbers")
	flag.BoolVar(&core, "core", false, "Set the statement to get a block from version core")
	flag.StringVar(&version1, "version1", "1.1.0", "Set the first version number")
	flag.StringVar(&version2, "version2", "1.0.5", "Set the second version number")
	flag.StringVar(&block, "block", "major", "Set the desired block")

	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		visit[f.Name] = true
	})

}

// main is to check the flags from the command line interface
// and execute the statements or return the no statement message.
func main() {

	if version {
		command_line.Print(0, "version: %s\n", build_information.Version)
	}

	if compare && visit["version1"] && visit["version2"] {

		result, err := statement.Compare(version1, version2)
		if err != nil {
			command_line.Print(1, "%s\n", err)
		}

		switch {
		case action:
			command_line.Output(map[string]any{"compare_result": result})
		default:
			command_line.Print(0, "%s\n", result)
		}

	}

	if core && visit["version1"] {

		number, err := statement.Core(version1, block)
		if err != nil {
			command_line.Print(1, "%s\n", err)
		}

		switch {
		case action:
			command_line.Output(map[string]any{"core_number": number})
		default:
			command_line.Print(0, "%d\n", number)
		}

	}

	command_line.Print(1, "Error! No statement flags were specified.\n")

}
