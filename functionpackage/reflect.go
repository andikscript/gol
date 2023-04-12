package functionpackage

import (
	"fmt"
	"reflect"
)

func ReflectTypeValue() {
	var value = "190"

	tipe := reflect.ValueOf(value)

	fmt.Println("type :", tipe.Type())

	if tipe.Kind() == reflect.Int {
		fmt.Println("nilai :", tipe.Int())
	}

	if tipe.Kind() == reflect.String {
		fmt.Println("String :", tipe.String())
	}
}

type car struct {
	Name string
	Age  int
}

func ReflectInterface() {
	car := car{"bmw", 2019}

	var reflectValue, reflectType = reflect.ValueOf(car), reflect.TypeOf(car)

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("Name :", reflectType.Field(i).Name)
		fmt.Println("Type :", reflectValue.Field(i).Type())
		fmt.Println("Value :", reflectValue.Field(i).Interface())
		fmt.Println("=================")
	}
}
