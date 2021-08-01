package main

import "fmt"

func listaCompras() {

	var cesta = []string{"abacaxi", "laranja", "banana", "carambola", "mamão", "uva", "abacate"}
	cesta = append(cesta, "morango", "melancia")

	fmt.Println("Minha lista de compras:")

	for i := 0; i < len(cesta); i++ {
		fmt.Println(i, ">>", cesta[i])
	}

}

func main() {
	listaCompras()
}
