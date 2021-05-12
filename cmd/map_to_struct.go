package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MyStruct struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	UserId    string `json:"user_id"`
	CreatedAt int64  `json:"created_at"`
	DataTest  string `json:"dsad"`
}

func main() {
	m := make(map[string]interface{})
	m["id"] = "2"
	m["name"] = "jack"
	m["user_id"] = "123"
	m["created_at"] = 5
	fmt.Println(m)

	// convert map to json
	jsonString, _ := json.Marshal(m)
	fmt.Println(string(jsonString))

	// convert json to struct
	s := MyStruct{}
	err := json.Unmarshal(jsonString, &s)
	if err != nil {
		return
	}

	v := reflect.ValueOf(s)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	fmt.Println(values)
	for _, t := range values {
		fmt.Print(t, "\u266A")
	}
	fmt.Println()
}

//package main
//
//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	x := struct {
//		Foo string
//		Bar int
//	}{"foo", 2}
//
//	v := reflect.ValueOf(x)
//
//	values := make([]interface{}, v.NumField())
//
//	for i := 0; i < v.NumField(); i++ {
//		values[i] = v.Field(i).Interface()
//	}
//
//	fmt.Println(values)
//}
