package functionpackage

type Motor struct {
	merk  string
	tahun uint16
}

// bisa ada method set jika tipe pointer
func (m *Motor) SetMotor(merk string, tahun uint16) {
	m.merk = merk
	m.tahun = tahun
}

func (m *Motor) GetMotor() (string, uint16) {
	return m.merk, m.tahun
}
