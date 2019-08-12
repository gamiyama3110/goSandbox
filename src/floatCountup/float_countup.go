package main

import (
	"fmt"
)

func main() {
	float64Countup()
	fmt.Println()
	float32Countup()
	fmt.Println()
	floatQuiz1()
	fmt.Println()
	floatQuiz2()
}

func float64Countup() (f64 float64) {
	for i := 0; i < 10; i++ {
		f64 += 0.1
		fmt.Println(f64)
	}
	return f64
}

func float32Countup() (f32 float32) {
	for i := 0; i < 10; i++ {
		f32 += float32(0.1)
		fmt.Println(f32)
	}
	return f32
}

func floatQuiz1() bool {
	fmt.Println("quiz 1.\n  a := 0.3 - 0.2\n  b := 0.2 - 0.1")
	a := 0.3 - 0.2
	b := 0.2 - 0.1
	c := a == b
	fmt.Printf("a == b : %t\n", c)
	return c
}

func floatQuiz2() bool {
	fmt.Println("quiz 2.\n  a := float64(0.3) - float64(0.2)\n  b := float64(0.2) - float64(0.1)")
	a := float64(0.3) - float64(0.2)
	b := float64(0.2) - float64(0.1)
	c := a == b
	fmt.Printf("a == b : %t\n", c)
	return c
}
