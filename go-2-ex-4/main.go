package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Name     string
	Students []Student
}

func main() {
	student1 := Student{FirstName: "Remo", LastName: "Albisser"}
	student2 := Student{FirstName: "Danial", LastName: "Najafi"}
	student3 := Student{FirstName: "Dylan", LastName: "Rodriguez"}
	student4 := Student{FirstName: "Lenny", LastName: "Amstutz"}
	student5 := Student{FirstName: "Domenic", LastName: "Troxler"}
	student6 := Student{FirstName: "Ricardo", LastName: "Grecco"}

	class1 := Class{
		Name: "Klasse INA-23bL",
		Students: []Student{
			student1, student2, student3,
		},
	}

	class2 := Class{
		Name: "Klasse INA-23aD",
		Students: []Student{
			student4, student5, student6,
		},
	}

	modules := map[int][]Class{
		346: {class1, class2},
		202: {class1},
		303: {class2},
	}

	fmt.Println("Klassen und ihre Schüler:")
	fmt.Printf("%s: %+v\n", class1.Name, class1.Students)
	fmt.Printf("%s: %+v\n", class2.Name, class2.Students)

	fmt.Println("\nModule und ihre Klassen:")
	for module, attendingClasses := range modules {
		fmt.Printf("Modul %d wird besucht von: \n", module)
		for _, class := range attendingClasses {
			fmt.Printf("  - %s\n", class.Name)
		}
	}
}
