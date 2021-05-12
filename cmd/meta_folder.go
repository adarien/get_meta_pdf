package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/barasher/go-exiftool"
)

type CustomString string

func (cb CustomString) UnmarshalJSON(data []byte) error {
	cb = CustomString(data)
	fmt.Print(cb, "\u266A")
	return nil
}

// var sung = []byte("\u266A")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type MetaStruct struct {
	Author                        CustomString `json:"Author"`
	CreateDate                    CustomString `json:"CreateDate"`
	Creator                       CustomString `json:"Creator"`
	CreatorTool                   CustomString `json:"CreatorTool"`
	DerivedFromDocumentID         CustomString `json:"DerivedFromDocumentID"`
	DerivedFromInstanceID         CustomString `json:"DerivedFromInstanceID"`
	DerivedFromOriginalDocumentID CustomString `json:"DerivedFromOriginalDocumentID"`
	DerivedFromRenditionClass     CustomString `json:"DerivedFromRenditionClass"`
	Description                   CustomString `json:"Description"`
	Directory                     CustomString `json:"Directory"`
	DocumentID                    CustomString `json:"DocumentID"`
	ExifToolVersion               CustomString `json:"ExifToolVersion"`
	FileAccessDate                CustomString `json:"FileAccessDate"`
	FileInodeChangeDate           CustomString `json:"FileInodeChangeDate"`
	FileModifyDate                CustomString `json:"FileModifyDate"`
	FileName                      CustomString `json:"FileName"`
	FilePermissions               CustomString `json:"FilePermissions"`
	FileSize                      CustomString `json:"FileSize"`
	FileType                      CustomString `json:"FileType"`
	FileTypeExtension             CustomString `json:"FileTypeExtension"`
	Format                        CustomString `json:"Format"`
	HasXFA                        CustomString `json:"HasXFA"`
	HistoryAction                 CustomString `json:"HistoryAction"`
	HistoryChanged                CustomString `json:"HistoryChanged"`
	HistoryParameters             CustomString `json:"HistoryParameters"`
	HistorySoftwareAgent          CustomString `json:"HistorySoftwareAgent"`
	HistoryWhen                   CustomString `json:"HistoryWhen"`
	InstanceID                    CustomString `json:"InstanceID"`
	Keywords                      CustomString `json:"Keywords"`
	Language                      CustomString `json:"Language"`
	Linearized                    CustomString `json:"Linearized"`
	MIMEType                      CustomString `json:"MIMEType"`
	MetadataDate                  CustomString `json:"MetadataDate"`
	ModifyDate                    CustomString `json:"ModifyDate"`
	OriginalDocumentID            CustomString `json:"OriginalDocumentID"`
	PDFVersion                    CustomString `json:"PDFVersion"`
	PageCount                     CustomString `json:"PageCount"`
	PageLayout                    CustomString `json:"PageLayout"`
	Producer                      CustomString `json:"Producer"`
	RenditionClass                CustomString `json:"RenditionClass"`
	SourceFile                    CustomString `json:"SourceFile"`
	Subject                       CustomString `json:"Subject"`
	TaggedPDF                     CustomString `json:"TaggedPDF"`
	Title                         CustomString `json:"Title"`
	Trapped                       CustomString `json:"Trapped"`
	XMPToolkit                    CustomString `json:"XMPToolkit"`
}

func main() {
	f, err := os.Create("./info/data2.txt")
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

	file, err := os.Open("./info/href.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	scanner := bufio.NewScanner(file)
	n := 1
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "♪")
		fmt.Println(s[2], s[5])
		fileInfos := et.ExtractMetadata(fmt.Sprintf("./files/%s", s[5]))

		for _, fileInfo := range fileInfos {
			if fileInfo.Err != nil {
				fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
				continue
			}
			jsonString, _ := json.Marshal(fileInfo.Fields)
			// fmt.Println(string(jsonString))

			s := MetaStruct{}
			err := json.Unmarshal(jsonString, &s)
			if err != nil {
				fmt.Println("fdsf", err)
				return
			}

			fmt.Println()
		}
		n++
		if n%1000 == 0 {
			time.Sleep(30 * time.Second)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
