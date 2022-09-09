package main

import (
	"flag"
	"github.com/gowizzard/compver/v3/command_line"
	"github.com/gowizzard/compver/v3/statement"
)

// compare is to save the boolean for the compare statement
// version1, version2 is to save the version numbers from the flags
var (
	compare            *bool
	version1, version2 *string
	visit              = make(map[string]bool)
)

// init is to parse the versions from the flags
// And check all visited flags
func init() {

	compare = flag.Bool("compare", false, "Set the statement to compare the version numbers")
	version1 = flag.String("version1", "1.1.0", "Set the first version number")
	version2 = flag.String("version2", "1.0.5", "Set the second version number")
	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		visit[f.Name] = true
	})

}

// main is to check the flags from the cli
// And execute the statements or return the no statement message
func main() {

	if *compare && visit["version1"] && visit["version2"] {

		compare, err := statement.Compare(version1, version2)
		if err != nil {
			command_line.Print(1, "%s\n", err)
		}
		command_line.Print(0, "%s\n", compare)

	}

	command_line.Print(1, "Error! No statement flags were specified.\n")

}
