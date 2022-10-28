// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package command_line_test

import (
	"errors"
	"github.com/gowizzard/compver/v5/command_line"
	"testing"
)

// TestPrint is to test the printing function.
func TestPrint(t *testing.T) {

	tests := []struct {
		name   string
		code   int
		format string
		a      any
	}{
		{
			name:   "PRINT=information",
			code:   0,
			format: "%s\n",
			a:      "major update",
		},
		{
			name:   "PRINT=error",
			code:   1,
			format: "%s\n",
			a:      errors.New("this is a testing error"),
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {
			command_line.Print(value.code, value.format, value.a)
		})

	}

}

// BenchmarkPrint is to test the Print function benchmark timing.
func BenchmarkPrint(b *testing.B) {

	for i := 0; i < b.N; i++ {
		command_line.Print(0, "%s\n", "This is a benchmark test")
	}

}
