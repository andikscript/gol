package functionpackage

import "fmt"

type Student struct {
	Nama  string
	Nilai int
}

func AksesStruct() (string, int) {
	var siswa Student
	siswa.Nama = "andik"
	siswa.Nilai = 90

	var sis = Student{"an", 123}
	fmt.Println(sis)

	return siswa.Nama, siswa.Nilai
}

func StructPointer() {
	var siswa = Student{"andik", 89}
	var siswaSecond *Student

	siswaSecond = &siswa

	fmt.Println(siswaSecond.Nama)
	fmt.Println(siswaSecond.Nilai)
}

func AliasStruct() {
	type People = Student

	var people People
	people.Nama = "andik"
	people.Nilai = 90

	fmt.Println(people)
}
