package command_line

import (
	"flag"
	"fmt"
	"os"
)

// Output is to build the attribute output
func Output(attributes map[string]any) {

	var build []byte
	for index, value := range attributes {
		build = fmt.Append(build, "::set-output")
		build = fmt.Appendf(build, " name=%s::%v\n", index, value)
	}
	fmt.Printf("%s", build)

	if flag.Lookup("test.v") == nil {
		defer os.Exit(0)
	}

}