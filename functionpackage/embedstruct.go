package functionpackage

import "fmt"

type Animal struct {
	Nama    string
	Makanan string
}

type Kucing struct {
	Kesukaan string
	Animal
}

func AksesEmbededStruct() {
	var hewan Kucing
	hewan.Nama = "meow"
	hewan.Makanan = "Ikan Asin"
	hewan.Kesukaan = "Berpetualang"

	fmt.Println(hewan)
	fmt.Println(hewan.Nama)
	fmt.Println(hewan.Makanan)
	fmt.Println(hewan.Kesukaan)
	fmt.Println(hewan.Animal.Nama)
}

type Komputer struct {
	MotherBoard string
	Processor   string
	Harddisk    string
}

type Laptop struct {
	Komputer
	MotherBoard string
	Cdrom       string
}

func AksesEmbededStructSameProperty() {
	var kom = Komputer{"gigabit", "intel", "sata"}
	var laptop = Laptop{kom, "ami", "rom"}

	fmt.Println(laptop.MotherBoard)
	fmt.Println(laptop.Processor)
	fmt.Println(laptop.Harddisk)
	fmt.Println(laptop.Cdrom)
	laptop.Komputer.MotherBoard = "msi - premium"
	fmt.Println(laptop.Komputer.MotherBoard)

	laptop.MotherBoard = "msi"
	laptop.Processor = "intel"
	laptop.Harddisk = "sata"
	laptop.Cdrom = "seagate"

	fmt.Println(laptop)
	fmt.Println(laptop.MotherBoard)
	fmt.Println(laptop.Processor)
	fmt.Println(laptop.Harddisk)
	fmt.Println(laptop.Cdrom)
	laptop.Komputer.MotherBoard = "msi - premium"
	fmt.Println(laptop.Komputer.MotherBoard)
	fmt.Println(laptop)

}
