// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package output_test

import (
	"github.com/gowizzard/compver/v5/output"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestWrite(t *testing.T) {

	tests := []struct {
		name     string
		path     string
		data     []byte
		flag     int
		perm     os.FileMode
		key      string
		value    any
		expected []byte
	}{
		{
			name:     "WRITE=new",
			path:     filepath.Join(os.TempDir(), "output"),
			data:     []byte(""),
			flag:     os.O_RDONLY,
			perm:     os.ModePerm,
			key:      "core_result",
			value:    4,
			expected: []byte("core_result=\"4\"\n"),
		},
		{
			name:     "WRITE=add",
			path:     filepath.Join(os.TempDir(), "GITHUB_OUTPUT"),
			data:     []byte("DEMO=\"data\"\n"),
			flag:     os.O_RDONLY,
			perm:     os.ModePerm,
			key:      "compare_result",
			value:    "major update",
			expected: []byte("DEMO=\"data\"\ncompare_result=\"major update\"\n"),
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			err := os.WriteFile(value.path, value.data, value.perm)
			if err != nil {
				t.Error(err)
			}

			err = os.Setenv("GITHUB_OUTPUT", value.path)
			if err != nil {
				t.Error(err)
			}

			output.Write(value.key, value.value)

			file, err := os.OpenFile(value.path, value.flag, value.perm)
			if err != nil {
				t.Error(err)
			}

			read, err := io.ReadAll(file)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(value.expected, read) {
				t.Errorf("expected: \"%v\", got \"%v\"", value.expected, read)
			}

		})

	}

}
