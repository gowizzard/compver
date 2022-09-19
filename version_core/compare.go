// Package version_core is used for the functions around the
// core blocks of a function. With regex the different core
// blocks like major, minor and patch are captured also the
// different blocks are compared.
package version_core

// Compare is to check the different version core blocks. The function returns
// a string with the information like: `no changes`, `... update` & `... downgrade`.
func Compare(blocks []Block) string {

	for _, value := range blocks {

		switch {
		case value.Number1 > value.Number2:
			return value.Name + " update"
		case value.Number1 == value.Number2:
			break
		case value.Number1 < value.Number2:
			return value.Name + " downgrade"
		}

	}

	return "no changes"

}
