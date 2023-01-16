// Copyright 2023 Jonas Kwiedor. All rights reserved.
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

// TestRead is to test the Read function with table driven tests.
func TestRead(t *testing.T) {

	tests := []struct {
		name     string
		path     string
		data     []byte
		perm     os.FileMode
		error    bool
		expected []byte
	}{
		{
			name:     "READ_DATA",
			path:     filepath.Join(t.TempDir(), "GITHUB_OUTPUT"),
			data:     []byte("key=\"value\"\n"),
			perm:     os.ModePerm,
			error:    false,
			expected: []byte("key=\"value\"\n"),
		},
		{
			name:     "FILE_ERROR",
			path:     "",
			data:     nil,
			perm:     0,
			error:    true,
			expected: nil,
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			if !value.error {
				err := os.WriteFile(value.path, value.data, value.perm)
				if err != nil {
					t.Error(err)
				}
			}

			t.Setenv("GITHUB_OUTPUT", value.path)

			content, err := output.Read()
			if err != nil && !value.error {
				t.Error(err)
			}

			if !reflect.DeepEqual(value.expected, content) && !value.error {
				t.Errorf("expected: \"%s\", got \"%s\"", value.expected, content)
			}

		})

	}

}

// BenchmarkRead is to test the Read function benchmark timing.
func BenchmarkRead(b *testing.B) {

	path, data, perm := filepath.Join(b.TempDir(), "GITHUB_OUTPUT"), []byte("key=\"value\"\n"), os.ModePerm

	err := os.WriteFile(path, data, perm)
	if err != nil {
		b.Error(err)
	}

	b.Setenv("GITHUB_OUTPUT", path)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = output.Read()
	}

}
