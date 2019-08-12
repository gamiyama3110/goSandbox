package main

import (
	"testing"
)

func TestFloat64(t *testing.T) {
	expected := 1.0
	result := float64Countup()

	// 1にならない
	if result == expected {
		t.Fatal("failed test")
	}
}

func TestFloat32(t *testing.T) {
	expected := float32(1.0)
	result := float32Countup()

	// 1にならない
	if result == expected {
		t.Fatal("faild test")
	}
}

func TestFloatQuiz1(t *testing.T) {
	expected := true
	result := floatQuiz1()

	if result != expected {
		t.Fatal("faild test")
	}
}

func TestFloatQuiz2(t *testing.T) {
	expected := false
	result := floatQuiz2()

	if result != expected {
		t.Fatal("faild test")
	}
}
