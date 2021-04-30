package main

import (
	"fmt"
	"github.com/barasher/go-exiftool"
)

func main() {

	et, err := exiftool.NewExiftool()
	if err != nil {
		fmt.Printf("Error when intializing: %v\n", err)
		return
	}
	defer func(et *exiftool.Exiftool) {
		err := et.Close()
		if err != nil {

		}
	}(et)

	fileInfos := et.ExtractMetadata("files/os-slides.pdf")

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			fmt.Printf("[%v] %v\n", k, v)
		}
	}
}
