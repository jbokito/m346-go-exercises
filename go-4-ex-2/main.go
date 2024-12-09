package main

import (
	"fmt"
	"math"
)

// computeHypotenuse berechnet die Länge der Hypotenuse gegeben die Längen der beiden Katheten a und b.
func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	// Testfall 1: Beide Katheten sind 3 und 4
	hypotenuse1 := computeHypotenuse(3, 4)
	fmt.Printf("computeHypotenuse(3, 4) = %.2f\n", hypotenuse1)

	// Testfall 2: Beide Katheten sind 5 und 12
	hypotenuse2 := computeHypotenuse(5, 12)
	fmt.Printf("computeHypotenuse(5, 12) = %.2f\n", hypotenuse2)

	// Testfall 3: Beide Katheten sind 8 und 15
	hypotenuse3 := computeHypotenuse(8, 15)
	fmt.Printf("computeHypotenuse(8, 15) = %.2f\n", hypotenuse3)

	// Testfall 4: Beide Katheten sind 7 und 24
	hypotenuse4 := computeHypotenuse(7, 24)
	fmt.Printf("computeHypotenuse(7, 24) = %.2f\n", hypotenuse4)
}
