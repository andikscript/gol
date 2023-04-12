package functionpackage

import "fmt"

func InterfaceKosong() {
	var terserah interface{}

	terserah = 12
	fmt.Println(terserah)

	terserah = "andik"
	fmt.Println(terserah)

	terserah = []string{"andik", "sap"}
	fmt.Println(terserah)
}

func InterfaceAnyAkses() {
	var terserah any

	terserah = 12
	fmt.Println(terserah)

	terserah = "andik"
	fmt.Println(terserah)

	terserah = []string{"andik", "sap"}
	fmt.Println(terserah)

}

func CastingInterface() {
	var terserah interface{}

	terserah = 10
	fmt.Println(terserah.(int) + 12)

	terserah = "123"
	fmt.Println(terserah.(string) + "23")
}

func CombineInterface() {
	var combine = map[string]interface{}{
		"satu": []string{"andik"},
		"dua":  123,
		"tiga": "str",
	}

	fmt.Println(combine)

	var sliceInterface = []interface{}{
		map[string]interface{}{
			"satu": []string{"andik"},
			"dua":  123,
			"tiga": "str",
		},
		"string",
		12,
		true,
	}

	fmt.Println(sliceInterface)
}

func CastingInterfaceToObject() {
	type person struct {
		name string
		age  int
	}

	var sec interface{} = &person{name: "andik", age: 123}

	fmt.Println(sec)
}
