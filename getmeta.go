package main

import (
	"encoding/json"
	"fmt"

	"github.com/barasher/go-exiftool"
)

type CustomString string

func (cb CustomString) UnmarshalJSON(data []byte) error {
	cb = CustomString(data)
	fmt.Print(cb, "\u266A")
	return nil
	//if reflect.TypeOf(data).Kind() == reflect.String {
	//	cb = CustomString(data)
	//} else if reflect.TypeOf(data).Kind() == reflect.Slice {
	//	cb = strings.Join(data, " ")
	//}
	//switch reflect.TypeOf(data).Kind() {
	//case reflect.String, reflect.Bool, reflect.Float64, reflect.Int:
	//	cb = CustomString(data)
	//	fmt.Print(cb, "\u266A")
	//	return nil
	//case reflect.Slice:
	//	//cb = CustomString(data[1:len(data)-1])
	//	//fmt.Print(cb, "\u266A")
	//	fmt.Println(len(data))
	//	return nil
	//default:
	//	return errors.New("CustomString: parsing \"" + string(data) + "\": unknown value")
	//}
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

	fileInfos := et.ExtractMetadata("climate-change-ma.pdf")

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}
		// fmt.Println(fileInfo.Fields)

		jsonString, _ := json.Marshal(fileInfo.Fields)
		// fmt.Println(string(jsonString))

		s := MetaStruct{}
		err := json.Unmarshal(jsonString, &s)
		if err != nil {
			fmt.Println("fdsf", err)
			return
		}

		fmt.Println()

		//v := reflect.ValueOf(s)
		//
		//values := make([]interface{}, v.NumField())
		//
		//for i := 0; i < v.NumField(); i++ {
		//	values[i] = v.Field(i).Interface()
		//}

		//fmt.Println(values)
		//for i, t := range values {
		//	if i == 2 || i == 28 || i == 41 {
		//		if reflect.TypeOf(s.Creator).Kind() == reflect.String {
		//			fmt.Print(s.Creator, "\u266A")
		//		} else if reflect.TypeOf(s.Creator).Kind() == reflect.Slice {
		//			//toString := strings.Join(s.Creator, " ")
		//			//fmt.Println(toString, "\u266A")
		//		}
		//	} else {
		//		fmt.Print(t, "\u266A")
		//	}
		//}
		//fmt.Println()

		//for _, t := range values {
		//	fmt.Print(t, "\u266A")
		//}

		//if reflect.TypeOf(s.Creator).Kind() == reflect.String {
		//	fmt.Println("String")
		//} else if reflect.TypeOf(s.Creator).Kind() == reflect.Slice {
		//	result1 := strings.Join(s.Creator, " ")
		//	fmt.Println(result1)
		//}

		//for k, v := range fileInfo.Fields {
		//	fmt.Printf("[%v] %v\n", k, v)
		//}
	}
}
