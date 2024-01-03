package util

import (
	"encoding/xml"
	"os"
)

func XMLDecode[T any](date []byte, mapped T) error {
	return xml.Unmarshal(date, &mapped)
}

func FolderExists(path string) bool {
	stat, err := os.Stat(path)
	if err == nil && stat.IsDir() {
		return true
	}
	return false
}
