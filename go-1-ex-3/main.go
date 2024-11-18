package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	var eyes = rand.Intn(6) + 1
	var when = time.Now()

	eyesFilePath := "eyes.txt"
	diceLogFilePath := "dice.log"

	eyesFile, err := os.Create(eyesFilePath)
	if err != nil {
		fmt.Println("Fehler beim Erstellen von eyes.txt:", err)
		return
	}
	defer eyesFile.Close()

	diceLogFile, err := os.Create(diceLogFilePath)
	if err != nil {
		fmt.Println("Fehler beim Erstellen von dice.log:", err)
		return
	}
	defer diceLogFile.Close()

	fmt.Fprintln(eyesFile, "The dice shows", eyes, "eyes")
	fmt.Fprintln(diceLogFile, "The dice was rolled at", when)

	fmt.Println("Die Augenzahl wurde in", eyesFilePath, "gespeichert.")
	fmt.Println("Der Zeitpunkt wurde in", diceLogFilePath, "gespeichert.")
}
