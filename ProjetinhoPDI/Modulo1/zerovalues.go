package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("Type: %T - Value: %v\n", i, i) // int zero value is 0
	fmt.Printf("Type: %T - Value: %v\n", f, f) // float64 zero value is 0.0
	fmt.Printf("Type: %T - Value: %v\n", b, b) // bool zero value is false
	fmt.Printf("Type: %T - Value: %v\n", s, s) // string zero value is "" (empty string)
}
