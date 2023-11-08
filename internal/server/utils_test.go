/*
Copyright © 2023 Ajith

*/
package server

import (
	"time"
	"io/fs"
	"testing"
	"path/filepath"

	"github.com/stretchr/testify/assert"
)

type MockFileInfo struct {
	FIName		string
	FISize		int64
	FIMode		fs.FileMode
	FIModTime	time.Time
	FIIsDir		bool
}

func (fi MockFileInfo) Name() string {
	return fi.FIName
}

func (fi MockFileInfo) Size() int64 {
	return fi.FISize
}

func (fi MockFileInfo) Mode() fs.FileMode {
	return fi.FIMode
}

func (fi MockFileInfo) ModTime() time.Time {
	return fi.FIModTime
}

func (fi MockFileInfo) IsDir() bool {
	return fi.FIIsDir
}

func (fi MockFileInfo) Sys() any {
	return nil
}

func TestGetContentProperty(t *testing.T) {

	mockFileInfo := MockFileInfo{
		"test_file.txt",
		512,
		fs.FileMode(0644),
		time.Now(),
		false,
	}
	dirEntry := fs.FileInfoToDirEntry(mockFileInfo)

	relativePath := "/path/to"

	property, err := getContentProperty(dirEntry, relativePath)

	assert.Equal(t, err, nil)
	
	expectedProperty := &Property{
		Name: "test_file.txt",
		Path: filepath.Clean("/path/to/test_file.txt"),
		Size: "512B",
	}
	assert.Equal(t, expectedProperty, property)
}

func TestGetDirectoryContents(t *testing.T) {

	mockEntries := []fs.DirEntry{
		fs.FileInfoToDirEntry(MockFileInfo{"test_file1.txt", 512, fs.FileMode(0644), time.Now(), false}),
		fs.FileInfoToDirEntry(MockFileInfo{"test_file2.txt", 256, fs.FileMode(0644), time.Now(), false}),
		fs.FileInfoToDirEntry(MockFileInfo{"test_folder", 0, fs.FileMode(0644), time.Now(), true}),
	}

	path := "/home/test_user/path/to"
	relativePath := "/path/to"

	contents, err := getDirectoryContents(path, relativePath, "", mockEntries)

	assert.Equal(t, err, nil)

	expectedContents := &Content{
		Path: "to",
		PreviousPath: filepath.Clean("/path"),
		Files: []Property{
				Property{
					Name:	"test_file1.txt",
					Path:	filepath.Clean("/path/to/test_file1.txt"),
					Size:	"512B",
				},
				Property{
					Name:	"test_file2.txt",
					Path:	filepath.Clean("/path/to/test_file2.txt"),
					Size:	"256B",
				},
			},
		Directories: []Property{
			Property{
				Name:	"test_folder",
				Path:	filepath.Clean("/path/to/test_folder"),
				Size:	"",
			},
		},
	}

	assert.Equal(t, expectedContents, contents)

}

func TestGetDirectoryContentsWithSearch(t *testing.T) {

	mockEntries := []fs.DirEntry{
		fs.FileInfoToDirEntry(MockFileInfo{"test_file1.txt", 512, fs.FileMode(0644), time.Now(), false}),
		fs.FileInfoToDirEntry(MockFileInfo{"test_file2.txt", 256, fs.FileMode(0644), time.Now(), false}),
		fs.FileInfoToDirEntry(MockFileInfo{"test_folder", 0, fs.FileMode(0644), time.Now(), true}),
		fs.FileInfoToDirEntry(MockFileInfo{"new_file1.txt", 512, fs.FileMode(0644), time.Now(), false}),
		fs.FileInfoToDirEntry(MockFileInfo{"new_file2.txt", 256, fs.FileMode(0644), time.Now(), false}),
		fs.FileInfoToDirEntry(MockFileInfo{"new_folder", 0, fs.FileMode(0644), time.Now(), true}),
	}

	path := "/home/test_user/path/to"
	relativePath := "/path/to"

	contents, err := getDirectoryContents(path, relativePath, "test", mockEntries)

	assert.Equal(t, err, nil)

	expectedContents := &Content{
		Path: "to",
		PreviousPath: filepath.Clean("/path"),
		Files: []Property{
				Property{
					Name:	"test_file1.txt",
					Path:	filepath.Clean("/path/to/test_file1.txt"),
					Size:	"512B",
				},
				Property{
					Name:	"test_file2.txt",
					Path:	filepath.Clean("/path/to/test_file2.txt"),
					Size:	"256B",
				},
			},
		Directories: []Property{
			Property{
				Name:	"test_folder",
				Path:	filepath.Clean("/path/to/test_folder"),
				Size:	"",
			},
		},
	}

	assert.Equal(t, expectedContents, contents)

}

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{0, "0B"},
		{1023, "1023B"},
		{1024, "1.00KB"},
		{1048576, "1.00MB"},
		{1073741824, "1.00GB"},
		{1099511627776, "1.00TB"},
		{1125899906842624, "1.00PB"},
		{1152921504606846976, "1.00EB"},
		{123456789, "117.74MB"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, formatBytes(test.input))
	}
}
