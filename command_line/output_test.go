// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package command_line_test

import (
	"github.com/gowizzard/compver/v5/command_line"
	"testing"
)

// TestOutput is to test the output function.
func TestOutput(t *testing.T) {

	tests := []struct {
		name  string
		key   string
		value any
	}{
		{
			name:  "KEY=compare_result",
			key:   "compare_result",
			value: "major update",
		},
		{
			name:  "KEY=core_result",
			key:   "core_result",
			value: 5,
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {
			command_line.Output(value.key, value.value)
		})

	}

}
