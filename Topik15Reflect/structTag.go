package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	ErrUsername = errors.New("Username kosong")
	ErrEmail    = errors.New("Email kosong")
	ErrPassword = errors.New("Password kosong")
)

type Siswa struct {
	Username string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

func IsValidate(data any) (bool, error) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("validate")
		if strings.Contains(tag, "required") {
			data := (v.Field(i))
			switch strings.ToLower(t.Field(i).Name) {
			case "username":
				if data.Interface() == "" {
					return false, ErrUsername
				}
			case "email":
				if data.Interface() == "" {
					return false, ErrEmail
				}
			case "password":
				if data.Interface() == "" {
					return false, ErrPassword
				}
			}
		}
	}

	return true, nil
}

func main() {
	siswa := Siswa{
		Username: "sa",
		Email:    "",
		Password: "ada",
	}

	result, err := IsValidate(siswa)
	if err != nil {
		fmt.Printf("Semuanya bisa login? %t\nPenyebab Eror > %s\n", result, err)
		return
	}
	fmt.Printf("Semuanya bisa login? %t\n", result)
}
