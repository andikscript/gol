package functionpackage

import (
	"fmt"
	"math"
)

type hitung interface {
	keliling() float32
	luas() float32
}

// persegi
type persegi struct {
	sisi float32
}

func (p persegi) keliling() float32 {
	return 4 * p.sisi
}

func (p persegi) luas() float32 {
	return p.sisi * p.sisi
}

// persegi panjang
type persegip struct {
	panjang float32
	lebar   float32
}

func (p persegip) keliling() float32 {
	return 2 * (p.panjang + p.lebar)
}

func (p persegip) luas() float32 {
	return p.panjang * p.lebar
}

func (p persegip) diagonal() float32 {
	var diagonal = (p.lebar * p.lebar) + (p.panjang * p.panjang)
	return float32(math.Sqrt(float64(diagonal)))
}

func AksesInterface() {
	var hitungPersegi hitung
	hitungPersegi = persegi{4}

	fmt.Println("----Persegi")
	fmt.Println("luas :", hitungPersegi.luas())
	fmt.Println("keliling :", hitungPersegi.keliling())

	var hitungPersegiP hitung
	hitungPersegiP = persegip{panjang: 12, lebar: 10}
	fmt.Println("-----Persegi Panjang")
	fmt.Println("luas :", hitungPersegiP.luas())
	fmt.Println("keliling :", hitungPersegiP.keliling())
	fmt.Println("diagonal :", hitungPersegiP.(persegip).diagonal())
}
