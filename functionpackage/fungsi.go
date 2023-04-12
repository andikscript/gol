package functionpackage

import "fmt"

var value int

func SetValue(val int) {
	value = val
}

func GetValue() int {
	return value
}

func demo() {
	fmt.Println("demo")
}
