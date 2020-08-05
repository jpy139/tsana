package file

import (
	"fmt"
	"os"
)

// Open open file according to name string
func Open(name string) (*os.File, error) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("Open file failed!, err:", err)
		return nil, err
	}

	return file, err
}

// func Close(handle *file)
