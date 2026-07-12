package main

import (
	"fmt"
	"reflect"
)

type Product struct {
	Name  string
	Price int
	Stock int
}

func nameStruct(field any) string {
	return reflect.TypeOf(field).Name()
}

func kindStruct(field any) reflect.Kind {
	return reflect.TypeOf(field).Kind()
}

func jumlahField(data any) int {
	return reflect.ValueOf(data).NumField()
}

func cekStruct(data any) {
	fmt.Printf("Nama Struct : %s\n", nameStruct(data))
	fmt.Printf("Jenis : %s\n", kindStruct(data))
	fmt.Printf("Jumlah Field : %d\n", jumlahField(data))
}

func tampilkanField(data any) {
	fmt.Printf("%s : %s\n", nameStruct(data), kindStruct(data))
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("Field %d \n", i)
		fmt.Printf("Nama : %s\n", f.Name)
		fmt.Printf("Tipe : %v\n", f.Type)
		fmt.Printf("Nilai Field %d : %v\n", i, v.Field(i).Interface())
		fmt.Println()
	}
}

func main() {
	product := Product{
		Name:  "Laptop",
		Price: 15000000,
		Stock: 5,
	}

	fmt.Printf("Nama Type : %s\n", nameStruct(product))
	fmt.Printf("Kind : %s\n\n", kindStruct(product))

	v := reflect.ValueOf(product)
	t := reflect.TypeOf(product)

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("Field %d \n", i)
		fmt.Printf("Nama : %s\n", f.Name)
		fmt.Printf("Tipe : %v\n", f.Type)
		fmt.Printf("Nilai Field %d : %v\n", i, v.Field(i).Interface())
		fmt.Println()
	}

	cekStruct(product)
	fmt.Println("\nBagian tampilkan field")
	tampilkanField(product)
}
