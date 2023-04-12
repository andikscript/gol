package functionpackage

import "fmt"

type hitung2d interface {
	luas() float32
	keliling() float32
}

type hitung3d interface {
	volume() float32
}

type hitungRuang interface {
	hitung2d
	hitung3d
}

type balok struct {
	panjang float32
	lebar   float32
	tinggi  float32
}

func (b balok) luas() float32 {
	return 2 * ((b.panjang * b.lebar) + (b.panjang * b.tinggi) + (b.lebar * b.tinggi))
}

func (b balok) keliling() float32 {
	return 4 * (b.panjang + b.lebar + b.tinggi)
}

func (b balok) volume() float32 {
	return b.panjang * b.lebar * b.tinggi
}

func AksesEmbededInterface() {
	var balokRuang hitungRuang
	balokRuang = balok{4, 3, 5}

	fmt.Println("luas :", balokRuang.luas())
	fmt.Println("keliling :", balokRuang.keliling())
	fmt.Println("volume :", balokRuang.volume())
}