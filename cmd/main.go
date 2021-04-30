package main

import (
	"bufio"
	"fmt"
	"github.com/barasher/go-exiftool"
	"log"
	"os"
	"strings"
)

//var sung = []byte("\u266A")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Create("./info/data.txt")
	check(err)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)

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

	file, err := os.Open("./info/read.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		fmt.Println(s[2], s[5])
		fileInfos := et.ExtractMetadata(fmt.Sprintf("./files/%s", s[5]))

		for _, fileInfo := range fileInfos {
			if fileInfo.Err != nil {
				fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
				continue
			}
			_, err = f.WriteString(fmt.Sprintf("{\"id\": %v,", s[2]))
			for k, v := range fileInfo.Fields {
				text := fmt.Sprintf("\"%v\": \"%v\",", k, v)
				_, err = f.WriteString(text)
			}
			_, err = f.WriteString("}\n")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
