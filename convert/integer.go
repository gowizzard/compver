// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package convert contains functions with which the
// type of value can be changed. If a value is not converted
// due to an error, the default value of the type to be converted is used.
package convert

import "strconv"

// Integer is to convert a string to an integer.
func Integer(value string) int {

	integer, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return integer

}
