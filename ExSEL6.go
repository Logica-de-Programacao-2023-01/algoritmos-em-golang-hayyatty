package main

import "fmt"

func main() {
	var (
		n1 float32
		n2 float32
	)
	fmt.Println("Digite um número :")
	fmt.Scan(&n1)
	fmt.Println("Digite outro número :")
	fmt.Scan(&n2)
	mult := n1 * n2
	if n1 > 0 && n2 > 0 {
		fmt.Println("A multiplicação dos números é :", mult)
	} else if n1 < 0 || n2 < 0 {
		fmt.Println("A soma dos número é :", n1+n2)
	}
}
