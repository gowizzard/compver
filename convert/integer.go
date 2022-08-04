package convert

import "strconv"

func Integer(value string) int {

	integer, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return integer

}
