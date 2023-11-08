/*
Copyright © 2023 Ajith

*/
package server

import (
	"os"
	"io/fs"
	"log"
	"net"
	"strconv"
	"path/filepath"
)

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}

func formatBytes(bytes int64) string {
	const unit = 1024
	
	if bytes < unit {
		return strconv.FormatInt(bytes, 10) + "B"
	}
	div := uint64(0)
	val := float64(bytes)
	for val >= unit {
		val /= unit
		div++
	}
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	
	return strconv.FormatFloat(val, 'f', 2, 64) + units[div]
}

func getContentProperty(entry fs.DirEntry, relativePath string) (*Property, error) {
	property := NewProperty()
	property.Name = entry.Name()
	property.Path = filepath.Join(relativePath, entry.Name())
	
	localFileInfo, err := entry.Info()
	if err != nil {
		return &Property{}, err
	} else {
		var size string
		if (localFileInfo.Size() == 0) {
			size = ""
		} else {
			size = formatBytes(localFileInfo.Size())
		}
		property.Size = size
	}

	return property, nil
}

func getDirectoryContents(path string, relativePath string, entries []fs.DirEntry) (*Content, error) {
	content := NewContent()
	content.Path = filepath.Base(path)
	for _, entry := range entries {
		fileInfo, err := entry.Info()
		if err != nil {
			return &Content{}, err
		}
		if fileInfo.IsDir() {
			property, err := getContentProperty(entry, relativePath)
			if err != nil {
				return &Content{}, err
			}
			content.Directories = append(content.Directories, *property)
		} else {
			property, err := getContentProperty(entry, relativePath)
			if err != nil {
				return &Content{}, err
			}
			content.Files = append(content.Files, *property)
		}
	}
	return content, nil
}

func GetContents(path string, relativePath string) (*Content, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return &Content{}, err
	}

	contents, err := getDirectoryContents(path, relativePath, entries)
	if err != nil {
		return &Content{}, err
	}
	return contents, nil
}
