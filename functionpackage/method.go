package functionpackage

type Tumbuhan struct {
	Nama string
	Age  int
}

// hanya bisa get tidak bisa untuk setter
func (t *Tumbuhan) GetTumbuhan() (string, int) {
	return t.Nama, t.Age
}
