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

func Exists(path string) bool {
    _, err := os.Stat(path)
    return err == nil
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

func getContentProperty(entry fs.DirEntry, relativePath string) *Property {
	property := NewProperty()
	property.Name = entry.Name()
	property.Path = filepath.Join(relativePath, entry.Name())
	
	localFileInfo, err := entry.Info()
	if err != nil {
		log.Panic(err)
	} else {
		var size string
		if (localFileInfo.Size() == 0) {
			size = ""
		} else {
			size = formatBytes(localFileInfo.Size())
		}
		property.Size = size
	}

	return property
}

func GetDirectoryContents(path string, relativePath string) *Content {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Panic(err)
	}
	
	content := NewContent()
	content.Path = filepath.Base(path)
	for _, entry := range entries {
		fileInfo, _ := os.Stat(filepath.Join(path, entry.Name()))
		if fileInfo.IsDir() {
			content.Directories = append(content.Directories, *getContentProperty(entry, relativePath))
		} else {
			content.Files = append(content.Files, *getContentProperty(entry, relativePath))
		}
	}
	return content
}
