package main

import (
	"fmt"
)

func main() {
	var quilometros int16
	quilometros = 150

	fmt.Println(quilometros)
}

//O tipo int8 não comportava o valor da variável quilometros porque corresponde a 8bits e comporta -128 ~ 127
