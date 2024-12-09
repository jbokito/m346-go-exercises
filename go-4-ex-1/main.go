package main

import (
	"errors"
	"fmt"
)

// Zusatzaufgabe
func computeGrade(gotPoints, maxPoints float64) (float64, error) {
	if gotPoints < 0 || maxPoints <= 0 {
		return 0, errors.New("invalid input: points cannot be negative and maxPoints must be positive")
	}
	if gotPoints > maxPoints {
		return 0, errors.New("invalid input: gotPoints cannot exceed maxPoints")
	}

	note := 1.0 + 5.0*(gotPoints/maxPoints)
	if note > 6.0 {
		note = 6.0
	}
	return note, nil
}

func main() {
	grade1, err1 := computeGrade(17.5, 28.0)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Printf("computeGrade(17.5, 28.0) = %.3f\n", grade1)
	}

	grade2, err2 := computeGrade(28.0, 28.0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Printf("computeGrade(28.0, 28.0) = %.3f\n", grade2)
	}

	grade3, err3 := computeGrade(-5.0, 28.0)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Printf("computeGrade(-5.0, 28.0) = %.3f\n", grade3)
	}
}
