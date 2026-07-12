package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {

	user := User{
		Name: "Ikhlas",
		Age:  19,
	}

	/*
		=========================
		REFLECT TYPE
		=========================

		Type digunakan untuk mengetahui
		informasi tentang tipe data.
	*/

	fmt.Println("===== TYPE OF =====")

	t := reflect.TypeOf(user)

	// A. Name()
	// Mengambil nama tipe
	fmt.Println("Name :", t.Name())

	// B. Kind()
	// Mengambil kategori tipe
	fmt.Println("Kind :", t.Kind())

	// C. NumField()
	// Menghitung jumlah field struct
	fmt.Println("Jumlah Field :", t.NumField())

	// D. Field()
	// Mengambil informasi field berdasarkan index
	field := t.Field(0)

	fmt.Println("Field ke-0 :", field)
	fmt.Println("Nama Field :", field.Name)
	fmt.Println("Tipe Field :", field.Type)


	/*
		=========================
		REFLECT VALUE
		=========================

		Value digunakan untuk
		mengakses nilai/data.
	*/

	fmt.Println("\n===== VALUE OF =====")

	value := reflect.ValueOf(user)

	// A. Interface()
	// Mengambil kembali nilai asli
	fmt.Println("Interface :", value.Interface())

	// B. Kind()
	// Mengambil kategori tipe data
	fmt.Println("Kind :", value.Kind())

	// C. Field()
	// Mengambil isi field berdasarkan index
	fmt.Println("Field ke-0 :", value.Field(0))

}