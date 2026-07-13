package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Login struct {
	Username string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

func IsValidated(data any) (bool, error) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("validate")
		if strings.Contains(tag, "required") {
			if v.Field(i).Interface() == "" {
				return false, fmt.Errorf("%s kosong", t.Field(i).Name)
			}
		}
	}

	return true, nil
}

func main() {
	login := Login{
		Username: "luffy",
		Email:    "luffy@gmail.com",
		Password: "1234321",
	}

	result, err := IsValidated(login)
	if err != nil {
		fmt.Printf("Semuanya bisa login? %t\nPenyebab Eror > %s\n", result, err)
		return
	}
	fmt.Printf("Semuanya bisa login? %t\n", result)
}
