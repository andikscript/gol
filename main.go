package main

import (
	fp "Chapter1/functionpackage"
	"fmt"
)

var first string
var number float32
var on bool

const PI float32 = 3.14

func main() {
	//var luas, keliling = fp.MultipleReturn(12, 5)
	//var luasSecond, kelilingSecond = fp.MultipleReturnSecond(13, 6)
	//
	//fmt.Println(luas, keliling)
	//fmt.Println(luasSecond, kelilingSecond)

	//var nama []string
	//nama = append(nama, "1")
	//nama = append(nama, "1")
	//nama = append(nama, "1")
	//
	//fmt.Println(nama)
	//
	//var name = make([]string, 10)
	//name[0] = "ad"
	//fmt.Println(name)
	//
	//var n = []string{"a", "n"}
	//fmt.Println(n)
	//
	//fmt.Println(.VaridicValue(1, 2, 3))
	//.VaridicNama("andik", "and", "an")
	//.VaridicNama("andik", nama...)
	//
	//const a int = 123
	//
	//fmt.Println(a)

	//var keliling, luas = fp.Closure(12)
	//fmt.Println(keliling)
	//fmt.Println(luas)
	//
	//var nama = []string{"andik", "mindai", "and"}
	//
	//fp.Iife(nama)
	//
	//var keluas = fp.NilaiKembalian(12)
	//fmt.Println(keluas())
	//
	//var kel, luass = fp.NilaiKembalianSecond(13)
	//fmt.Println(kel())
	//fmt.Println(luass())

	//fp.Pointer()
	//
	//var num int = 12
	//fmt.Println(num)
	//
	//fp.ChangeNumber(&num, 18)
	//fmt.Println(num)

	//var nama, nilai = fp.AksesStruct()
	//fmt.Println(nama)
	//fmt.Println(nilai)
	//
	//var sis fp.Student
	//sis.Nama = "mindai"
	//sis.Nilai = 91

	//fmt.Println(sis.Nama)
	//fmt.Println(sis.Nilai)

	//fp.StructPointer()

	//fp.AksesEmbededStruct()

	//fp.AksesEmbededStructSameProperty()

	//fp.AnonymousStruct()

	//fp.AnonymousStructValue()

	//fp.SliceStruct()

	//fp.SliceAnonStruct()

	//fp.AliasStruct()

	//var tumbuhan fp.Tumbuhan
	//tumbuhan.Nama = "Jeruk"
	//tumbuhan.Age = 9
	//var nama, age = tumbuhan.GetTumbuhan()
	//
	//fmt.Println("nama tumbuhan :", nama+", umur tumbuhan :", age)
	//
	//fp.SetValue(123)
	//fmt.Println("value :", fp.GetValue())

	//var motor fp.Motor
	//motor.SetMotor("Honda", 2003)
	//var merk, tahun = motor.GetMotor()
	//
	//fmt.Println("Merk :", merk+", Tahun :", tahun)

	//fp.AksesInterface()

	//fp.AksesEmbededInterface()

	//fp.InterfaceKosong()

	//fp.InterfaceAnyAkses()

	//fp.CastingInterface()

	//fp.CombineInterface()

	//fp.CastingInterfaceToObject()

	//fp.FungsiSebagaiParameter(func() int {
	//	return 123
	//})

	//fp.FungsiSebagaiParameter2(12, func(value int) int {
	//	return value
	//})

	//fp.ReflectTypeValue()

	fp.ReflectInterface()
}

func sliceMap() {
	var slicemap = []map[string]int{
		{"senin": 1, "selasa": 2},
		{"senin": 1, "selasa": 2},
	}

	for i, m := range slicemap {
		fmt.Println(i, m)
	}
}

func mapGo() {
	var buah = map[string]int{}

	buah["mangga"] = 10
	buah["pepaya"] = 15

	fmt.Println(buah)

	for s, i := range buah {
		fmt.Println(s, i)
	}

	delete(buah, "mangga")

	fmt.Println(buah)

	// mencari nilai dalam map
	var value, isExist = buah["pepaya"]

	if isExist {
		fmt.Println(value)
	}

	if !isExist {
		fmt.Println(value)
	}
}

func panjang(arr []string) int {
	var val int
	for i := 0; i < len(arr); i++ {
		val += i
	}

	return val
}

func slice3() {
	// convert slice to array
	var buah = []string{"nanas", "melon", "semangka"}
	var buahh = make([]string, len(buah))

	for i := 0; i < len(buah); i++ {
		buahh[i] = buah[i]
	}

	fmt.Println(buah)
	fmt.Println(buahh)
}

func slice2() {
	// convert array to slice
	var buah [4]string
	buah[0] = "nanas"
	buah[1] = "pepaya"
	buah[2] = "melon"
	buah[3] = "semangka"

	var sliceBuah = make([]string, len(buah))

	fmt.Println(buah)

	for i := 0; i < len(buah); i++ {
		sliceBuah[i] = buah[i]
	}

	fmt.Println(sliceBuah)
}

func slice() {
	var buah = []string{"nanas", "mangga"}
	var arrbuah []string
	arrbuah = buah[0:2]
	arbuah := make([]string, len(arrbuah))

	fmt.Println(arrbuah)

	arrbuah[0] = "pepaya"
	fmt.Println(arrbuah)
	fmt.Println(buah)

	copy(arbuah, arrbuah)

	fmt.Println(arrbuah)
}

func changeOn() {
	on = true
}

func changeNumber() {
	number = 12.0001
}

func change() {
	first = "gen"
}

func percabangan(value uint8) {
	if value == 0 {
		fmt.Println("kosong")
	}

	if value != 0 {
		fmt.Println("isi")
	}
}

func switchCase(name string) {
	switch name {
	case "a":
		fmt.Println("1")
	case "b":
		fmt.Println("2")
	default:
		fmt.Println(name)
	}
}

func perulangan() {
outerLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if i == 5 && j == 2 {
				break outerLoop
			}

			fmt.Print(j, " ")
		}
	}
}

func array() {
	var arr [2]string
	arr[0] = "andik"
	arr[1] = "mindai"

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for _, ar := range arr {
		fmt.Println(ar, " ")
	}

	var no [3]uint8
	no[0] = 1
	no[1] = 2
	no[2] = 3

	fmt.Println(no)
}

func test() {
	var arr [3]string
	arr[0] = "1"
	arr[1] = "2"
	arr[2] = "3"

	fmt.Println(arr)

	var slicee = []string{"1", "2", "3"}

	fmt.Println(slicee)

	var mapG = map[int]string{}
	mapG[1] = "senin"
	mapG[2] = "selasa"

	fmt.Println(mapG)

	var sliceMap = []map[string]int{
		{"value": 98, "bool": 1},
		{"value": 90, "bool": 0},
	}

	fmt.Println(sliceMap)

	for _, m := range sliceMap {
		fmt.Println(m["value"], m["bool"])
	}

	var ar = make([]int, 5)
	ar[0] = 10

	fmt.Println(ar)

	var as = [...]int{1, 2, 3, 4, 5}

	fmt.Println(as)
}
