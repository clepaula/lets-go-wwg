package main

import "fmt"

func mapa() map[string]string {
	local := make(map[string]string)

	local["Tóquio"] = "Japão"
	local["Delhi"] = "India"
	local["Xangai"] = "China"
	local["São Paulo"] = "Brasil"
	local["Rio de Janeiro"] = "Brasil"
	local["Cidade do México"] = "México"
	local["Pequin"] = "China"
	local["Cairo"] = "Egito"

	return local
}

func contaPais(local map[string]string) map[string]int {
	contador := make(map[string]int)
	for _, pais := range local {
		contagem, paisListado := contador[pais]
		if paisListado {
			contador[pais] = contagem + 1
		} else {
			contador[pais] = 1
		}
	}

	return contador

}

func main() {
	lista := mapa()
	contagem := contaPais(lista)

	fmt.Println(contagem)

}
