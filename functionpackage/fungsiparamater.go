package functionpackage

import "fmt"

// fungsi yang dijadikan tipe di parameter harus ada nilai kembaliannya di fungsi yang dijadikan tipe data

func FungsiSebagaiParameter(fun func() int) {
	fmt.Println(fun())
}

func FungsiSebagaiParameter2(value int, fun func(value int) int) {
	value *= value
	fmt.Println(fun(value))
}
