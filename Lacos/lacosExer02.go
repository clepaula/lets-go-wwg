package main

import "fmt"

type Apes struct {
	morador    string
	numero     int
	temGaragem bool
}

func main() {
	imovel1 := Apes{
		morador:    "Julieta",
		numero:     340,
		temGaragem: true,
	}

	imovel2 := Apes{
		morador:    "Fermina",
		numero:     242,
		temGaragem: true,
	}

	imovel3 := Apes{
		morador:    "Jucelia",
		numero:     343,
		temGaragem: false,
	}
	imovel4 := Apes{
		morador:    "Jucelia",
		numero:     244,
		temGaragem: true,
	}

	imoveis := []Apes{imovel1, imovel2, imovel3, imovel4}

	for _, imovel := range imoveis {
		temGaragem := "Não"
		if imovel.temGaragem {
			temGaragem = "Sim"
		}
		fmt.Printf("O apartamento número %d é habitado por %s. O imóvel tem vaga de garagem? %s\n", imovel.numero, imovel.morador, temGaragem)
	}
}
