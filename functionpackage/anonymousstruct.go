package functionpackage

import "fmt"

func AnonymousStruct() {
	var siswa = struct {
		Nama  string
		Nilai int
	}{}

	siswa.Nama = "andik"
	siswa.Nilai = 89

	fmt.Println(siswa)
}

func AnonymousStructValue() {
	var siswa = struct {
		Nama  string
		Nilai int
	}{
		Nama:  "Andik",
		Nilai: 90,
	}

	fmt.Println(siswa)
}
