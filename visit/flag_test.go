package visit_test

import (
	"flag"
	"github.com/gowizzard/compver/v2/visit"
	"testing"
)

// TestFlag is to check the Flag function
// We simulate flags with name, value & usage and test it table driven
func TestFlag(t *testing.T) {

	tests := []struct {
		name  string
		value string
		usage string
	}{
		{
			name:  "version1",
			value: "1.1.0",
			usage: "Set the first version number",
		},
		{
			name:  "version2",
			value: "1.0.5",
			usage: "Set the second version number",
		},
	}

	for _, value := range tests {

		_ = flag.String(value.name, value.value, value.usage)
		err := flag.Set(value.name, value.value)
		if err != nil {
			t.Fatalf("%v\n", err)
		}
		flag.Parse()

		visit.Flag()
		if !visit.Set[value.name] {
			t.Fatalf("got \"%v\" for flag \"%s\"", visit.Set[value.name], value.name)
		}

	}

}
