package convert

import "strconv"

// Integer is to convert a string to an integer
func Integer(value string) int {

	integer, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return integer

}
