package functionpackage

import "fmt"

func Closure(sisi int) (keliling int, luas int) {
	var getLuasKeliling = func(sisi int) (keliling int, luas int) {
		keliling = 2 * sisi
		luas = sisi * sisi

		return
	}

	keliling, luas = getLuasKeliling(sisi)
	return
}

func Iife(name []string) {
	var find string
	var findNama = func(nama string) string {
		for _, s := range name {
			if s == nama {
				find = "find"
				break
			}

			find = "not find"
		}

		return find
	}("and")

	fmt.Println(findNama)
}

func NilaiKembalian(value int) func() (keliling int, luas int) {
	return func() (keliling int, luas int) {
		keliling = 2 * value
		luas = value * value

		return
	}
}

func NilaiKembalianSecond(value int) (func() (keliling int), func() (luas int)) {
	return func() (keliling int) {
			keliling = 2 * value

			return
		}, func() (luas int) {
			luas = value * value

			return
		}
}
