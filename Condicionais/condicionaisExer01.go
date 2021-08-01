package main

import "fmt"

func main() {
	a, b, c := 5, 13, 9

	if a > b && a > c {
		fmt.Printf("A variável %v é maior que %v e %v", a, b, c)
	}
	if b > a && b > c {
		fmt.Printf("A variável %v é maior que %v e %v", b, c, a)
	}
	if c > b && c > a {
		fmt.Printf("A variável %v é maior que %v e %v", c, b, a)
	}
	if a == b || b == c {
		fmt.Printf("Não há uma variável maior")
	}

}
