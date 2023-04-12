package functionpackage

import "fmt"

func VaridicNama(name string, nama ...string) {
	fmt.Println(name)

	for i := 0; i < len(nama); i++ {
		fmt.Printf("%s ", nama[i])
	}

	fmt.Println()
}

func VaridicValue(value ...int) (total int) {
	for i := 0; i < len(value); i++ {
		total += value[i]
	}

	return
}
