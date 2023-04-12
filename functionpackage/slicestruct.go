package functionpackage

import "fmt"

type Person struct {
	Name string
	Age  int
}

func SliceStruct() {
	var allPerson = []Person{
		{"andik", 24},
		{"mindai", 23},
		{"saputro", 22},
	}

	fmt.Println(allPerson)

	for _, person := range allPerson {
		fmt.Println("Name :", person.Name+", Age :", person.Age)
	}
}

func SliceAnonStruct() {
	var person = []struct {
		Name string
		Age  int
	}{}

	person = append(person, struct {
		Name string
		Age  int
	}{Name: "andik", Age: 21})

	fmt.Println(person)

	var person2 = []struct {
		Name string
		Age  int
	}{
		{"andik", 24},
		{"mindai", 23},
		{"saputro", 22},
	}

	fmt.Println(person2)
}
