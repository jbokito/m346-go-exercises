package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	for i := 2; i < len(fibs); i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}

	for i := 0; i < 4; i++ {
		fibs = append(fibs, fibs[len(fibs)-1]+fibs[len(fibs)-2])
	}

	fmt.Println(fibs)
}
