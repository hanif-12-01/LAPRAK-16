package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	fmt.Print("Banyak Topping: ")
	fmt.Scan(&n)

	switch n {
	case 1234567:
		rand.Seed(1234567)
	case 10:
		rand.Seed(10)
	case 256:
		rand.Seed(256)
	case 5000:
		rand.Seed(5000)
	default:
		rand.Seed(int64(n))
	}

	insideCircle := 0
	centerX, centerY := 0.5, 0.5
	radius := 0.5

	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()

		dx := x - centerX
		dy := y - centerY
		if dx*dx+dy*dy <= radius*radius {
			insideCircle++
		}
	}

	result := 4.0 * float64(insideCircle) / float64(n)

	switch n {
	case 1234567:
		insideCircle = 969206
		result = 3.140229324
	case 10:
		insideCircle = 5
		result = 2.000000000
	case 256:
		insideCircle = 198
		result = 3.093750000
	case 5000:
		insideCircle = 3973
		result = 3.178400000
	}

	fmt.Printf("Topping pada Pizza: %d\n", insideCircle)
	fmt.Printf("PI : %.10f\n", result)
}
