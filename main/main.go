package main

import (
	"TSAnalysis/file"
	"fmt"
)

func main() {
	// fmt.Println("nihao")
	fileHandle, err := file.Open("F:/projects/m7-op/record stream/streams/Test_11_TP3209.ts")
	if err != nil {
		fmt.Println("Open file error:", err)
	}
	fileHandle.Read()
}
