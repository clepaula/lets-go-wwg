package main

import "fmt"

func main() {
	timeAmarelo := make(map[int]string)
	timeAmarelo[1] = "Fernando"
	timeAmarelo[2] = "João"
	timeAmarelo[3] = "Lúcia"
	timeAmarelo[4] = "Mariana"
	timeAmarelo[5] = "Ana"

	timeVermelho := make(map[int]string)
	timeVermelho[1] = "Helena"
	timeVermelho[2] = "Jonas"
	timeVermelho[3] = "José"
	timeVermelho[4] = "Juliana"

	//fmt.Println("Time Amarelo:", timeAmarelo)
	fmt.Println("Time Amarelo:", timeAmarelo[1], timeAmarelo[2], timeAmarelo[3], timeAmarelo[4], timeAmarelo[5])

	//fmt.Println("Time Vermelho:", timeVermelho)
	fmt.Println("Time Vermelho:", timeVermelho[1], timeVermelho[2], timeVermelho[3], timeVermelho[4])

}
