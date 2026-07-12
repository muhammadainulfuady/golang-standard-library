package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name     string `required:"true" max:"100"`
	Email    string `required:"true" max:"50"`
	Password string `required:"true" max:"12"`
}

func readField(value any) {
	typeSample := reflect.TypeOf(value)
	fmt.Println("Type of", typeSample.Name())

	for i := 0; i < typeSample.NumField(); i++ {
		structField := typeSample.Field(i)
		fmt.Println(structField.Name)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))

	}
}

func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	sample := Sample{
		Name:     "Muhammad Ainul Fuady",
		Email:    "ainulfuadi1234@gmail.com",
		Password: "",
	}

	fmt.Println(IsValid(sample))

}
