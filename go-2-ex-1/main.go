package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   byte
	Month byte
	Year  int
}

type Profile struct {
	Name             FullName
	Birth            BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {

	var me = Profile{
		Name: FullName{
			FirstName: "Remo",
			LastName:  "Albisser",
		},
		Birth: BirthDate{
			Day:   1,
			Month: 4,
			Year:  2007,
		},
		NumberOfSiblings: 2,
		ZodiacSign:       '\u2648',
	}

	fmt.Println("Steckbrief:", me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
