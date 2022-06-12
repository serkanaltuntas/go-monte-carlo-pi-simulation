package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	size := 10_000_000
	fmt.Println("Calculating the π...")
	inside := 0
	for i := 0; i < size; i++ {
		x := rand.Float64()
		y := rand.Float64()
		distance := x*x + y*y
		if distance < 1 {
			inside += 1
		}
	}
	result := 4 * float64(inside) / float64(size)
	fmt.Println("π:", result)
}
