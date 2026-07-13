package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Login struct {
	Username string `validate:"required,max=20,min=5"`
	Email    string `validate:"required,max=20,min=5"`
	Password string `validate:"required,max=20,min=10"`
}

func IsValid(data any) (bool, error) {

	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	for i := 0; i < t.NumField(); i++ {
		value := v.Field(i).Kind()
		field := t.Field(i)
		tag := field.Tag.Get("validate")
		rules := strings.Split(tag, ",")

		for _, rule := range rules {
			switch {
			case rule == "required":
				if value.String() == "" {
					return false, fmt.Errorf("%s kosong", field.Name)
				}
			case strings.HasPrefix(rule, "min="):
				min := rule[4:]
				conv, _ := strconv.Atoi(min)
				if len(value.String()) < conv {
					return false, fmt.Errorf("%s minimal %s karakter", field.Name, min)
				}
			case strings.HasPrefix(rule, "max="):
				max := rule[4:]
				conv, _ := strconv.Atoi(max)
				if len(value.String()) > conv {
					return false, fmt.Errorf("%s maximal %s karakter", field.Name, max)
				}
			}
		}
	}
	return true, nil
}

func main() {
	login := Login{
		Username: "lufyoo",
		Email:    ".com",
		Password: "12345678910",
	}
	result, err := IsValid(login)
	if err != nil {
		fmt.Printf("Semuanya bisa login? %t\nPenyebab Eror > %s\n", result, err)
		return
	}
	fmt.Printf("Semuanya bisa login? %t\n", result)

}
