package day05

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	fmt.Print(t.Name() + "\n")
	// Example binary numbers (4-bit)
	a := [4]bool{false, false, true, false} // 0010 (10 in decimal)
	b := [4]bool{false, true, false, false} // 0100 (12 in decimal)

	// Perform binary addition
	sum, carry := BinaryAdder(a, b)

	// Print result
	fmt.Print("Sum: ")
	for i := len(sum) - 1; i >= 0; i-- {
		fmt.Print(boolToInt(sum[i]))
	}
	fmt.Printf("\nCarry: %d\n", boolToInt(carry))
}

// Helper function to convert bool to int for display
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// 4-bit Binary Adder
func BinaryAdder(a, b [4]bool) (sum [4]bool, carry bool) {
	var c bool
	for i := 0; i < 4; i++ {
		sum[i], c = FullAdder(a[i], b[i], c)
	}
	carry = c
	return
}

// Half Adder
func HalfAdder(a, b bool) (sum, carry bool) {
	sum = XOR(a, b)
	carry = AND(a, b)
	return
}

// Full Adder
func FullAdder(a, b, cin bool) (sum, carry bool) {
	sum1, carry1 := HalfAdder(a, b)
	sum, carry2 := HalfAdder(sum1, cin)
	carry = OR(carry1, carry2)
	return
}

// AND gate
func AND(a, b bool) bool {
	return a && b
}

// OR gate
func OR(a, b bool) bool {
	return a || b
}

// NOT gate
func NOT(a bool) bool {
	return !a
}

// NAND gate
func NAND(a, b bool) bool {
	return !(a && b)
}

// NOR gate
func NOR(a, b bool) bool {
	return !(a || b)
}

// XOR gate
func XOR(a, b bool) bool {
	return a != b
}

// XNOR gate
func XNOR(a, b bool) bool {
	return a == b
}
