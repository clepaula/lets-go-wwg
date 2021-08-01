package main

import "fmt"

func main() {
	numero := 13

	if (numero%2) == 0 && numero != 0 {
		fmt.Println(numero, "é múltiplo de 2")
	} else if (numero%3) == 0 && numero != 0 {
		fmt.Println(numero, "é múltiplo de 3")
	} else if numero == 0 {
		fmt.Println(numero, "é zero")
	} else {
		fmt.Println(numero, "é um número diferente de zero e não múltiplo de 2 nem de 3")
	}
}
