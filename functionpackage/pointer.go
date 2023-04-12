package functionpackage

import "fmt"

func Pointer() {
	var nama string = "andik"
	var point *string = &nama // referencing
	var pointSecond *string = point

	fmt.Println(*point) // deferencing
	fmt.Println(*pointSecond)

	nama = "sap"

	fmt.Println(*point) // deferencing
	fmt.Println(*pointSecond)

	*point = "min"

	fmt.Println(*point) // deferencing
	fmt.Println(*pointSecond)
	fmt.Println(nama)
}

func ChangeNumber(originalValue *int, value int) {
	*originalValue = value
}
