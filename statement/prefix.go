// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package statement

import "strings"

// Prefix is to get the version, check if the prefix
// exists & return the version without the prefix.
func Prefix(version, prefix string) string {

	if strings.HasPrefix(version, prefix) {
		version = strings.TrimPrefix(version, prefix)
	}

	return version

}
