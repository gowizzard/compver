// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package output_test

import (
	"github.com/gowizzard/compver/v5/output"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// TestWrite is to test the write function.
func TestWrite(t *testing.T) {

	tests := []struct {
		name     string
		path     string
		key      string
		value    any
		data     []byte
		error    bool
		expected []byte
	}{
		{
			name:     "CREATE_DATA",
			path:     filepath.Join(t.TempDir(), "output"),
			key:      "key",
			value:    "value",
			data:     []byte(""),
			error:    false,
			expected: []byte("key=\"value\"\n"),
		},
		{
			name:     "ADD_DATA",
			path:     filepath.Join(t.TempDir(), "GITHUB_OUTPUT"),
			key:      "key_2",
			value:    "value",
			data:     []byte("key_1=\"value\"\n"),
			error:    false,
			expected: []byte("key_1=\"value\"\nkey_2=\"value\"\n"),
		},
		{
			name:     "FILE_ERROR",
			path:     "",
			key:      "",
			value:    nil,
			data:     nil,
			error:    true,
			expected: nil,
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			t.Setenv("GITHUB_OUTPUT", value.path)

			err := output.Write(value.key, value.value, value.data)
			if err != nil && !value.error {
				t.Error(err)
			}

			if !value.error {

				read, err := os.ReadFile(value.path)
				if err != nil {
					t.Error(err)
				}

				if !reflect.DeepEqual(value.expected, read) {
					t.Errorf("expected: \"%v\", got \"%v\"", value.expected, read)
				}

			}

		})

	}

}

// BenchmarkWrite is to test the Write function benchmark timing.
func BenchmarkWrite(b *testing.B) {

	path, key, value, data := filepath.Join(b.TempDir(), "GITHUB_OUTPUT"), "key", "value", []byte("")

	b.Setenv("GITHUB_OUTPUT", path)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = output.Write(key, value, data)
	}

}
