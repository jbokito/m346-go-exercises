package main

import "fmt"

func main() {

	firstName := "Remo"
	lastName := "Albisser"
	dayOfBirth := 01 - 04 - 2024
	monthOfBirth := 04 - 2024
	yearOfBirth := 2024
	numberOfSiblings := 3
	heightInMeters := 1.85
	zodiacSign := '\u264E'

	// Ausgabe der Informationen
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
