package main

import "fmt"

//Faça um algoritmo que leia três números inteiros e mostre a soma entre eles
func main() {
	var (
		n1, n2, n3 int64
	)
	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&n1)
	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&n2)
	fmt.Println("Digite o terceiro e útimo número:")
	fmt.Scanln(&n3)
	soma := n1 + n2 + n3
	println("A soma dos 3 números é :", soma)
}
