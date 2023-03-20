package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3 int

	fmt.Println("Digite um número:")
	fmt.Scan(&n1)
	fmt.Println("Digite outro número")
	fmt.Scan(&n2)
	fmt.Println("Digite o terceito número:")
	fmt.Scan(&n3)

	mp := ((2 * n1) + (3 * n2) + (5 * n3)) / 10
	println("A media ponderada é", mp)
}
