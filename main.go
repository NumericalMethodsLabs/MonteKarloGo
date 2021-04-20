package main

import (
	"fmt"
	"math"
	"math/rand"
)

func f(x, y float64) float64 {
	return math.Sin(math.Pi*x) * math.Cos(y*math.Pi/2)
}
func main() {
	n := 1000
	I := 0.405284734569351 // Двойной интеграл от sin(pi * x) * cos(pi*y/2)dxdy
	fmt.Println("I:", I)

	S := 0.0
	for i := 0; i < n; i++ {
		S += f(rand.Float64(), rand.Float64())
	}
	S = S / float64(n)
	fmt.Println("1. Sn(f):", S)
	fmt.Println("погрешность:", math.Abs(I-S))

	S = 0.0
	for i := 0; i < n; i++ {
		x, y, z := rand.Float64(), rand.Float64(), rand.Float64()
		if z < f(x, y) {
			S += 1
		}
	}
	S = S / float64(n)
	fmt.Println("2. Sn(f):", S)
	fmt.Println("погрешность:", math.Abs(I-S))
}
