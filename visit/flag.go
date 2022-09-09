package visit

import "flag"

// Set is to share the visited flags
var Set = make(map[string]bool)

// Flag is to check the visited flags
func Flag() {
	flag.Visit(func(f *flag.Flag) {
		Set[f.Name] = true
	})
}
