package main

import "fmt"

func main() {

	ePar := false
	for !ePar {
		var x int
		fmt.Scanln(&x)
		if x%2 == 0 {
			fmt.Println("Agora sim, podemos dividir igualmente entre mim e você!")
			ePar = true
		}

	}

}
