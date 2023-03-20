package main

import "fmt"

func main() {
	var (
		n1 float32
		n2 float32
	)
	fmt.Println("Digite o primeiro número:")
	fmt.Scan(&n1)
	fmt.Println("Digite o segundo número:")
	fmt.Scan(&n2)
	if n1 > n2 {
		fmt.Println("O maior número é o :", n1)
	} else if n1 == n2 {
		fmt.Println("Os dois são iguais")
	} else if n1 < n2 {
		fmt.Println("O maior número é o :", n2)
	}

}
