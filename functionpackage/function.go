package functionpackage

import (
	"fmt"
)

func FungsiFirst() {
	fmt.Println("ini fungsi pertama")
	fmt.Println(MasukkanParameter(12))
	TanpaKembalian(0)
	fmt.Println(TanpaParameter())
	TanpaParamReturn()
}

func MasukkanParameter(value int) int {
	return value
}

func TanpaKembalian(value int) {
	if value != 0 {
		fmt.Println(value)
	} else {
		return
	}
}

func TanpaParameter() string {
	return "nilai"
}

func TanpaParamReturn() {
	fmt.Println("tanpa parameter dan kembalian")
}

func MultipleReturn(panjang int16, lebar int16) (int16, int16) {
	var keliling int16
	keliling = 2 * (panjang + lebar)

	var luas int16
	luas = panjang * lebar

	return keliling, luas
}

func MultipleReturnSecond(panjang int, lebar int) (keliling int, luas int) {
	keliling = 2 * (panjang + lebar)
	luas = panjang * lebar

	return
}
