package main

import "fmt"

func main() {
	var (
		n1 float32
		n2 float32
		n3 float32
	)
	fmt.Println("Digite o primeiro número")
	fmt.Scan(&n1)
	fmt.Println("Digite o segundo número")
	fmt.Scan(&n2)
	fmt.Println("Digite o terceiro número")
	fmt.Scan(&n3)

	if n1 > n2 && n1 > n3 {
		fmt.Println("O maior deles é ", n1)
	} else if n2 > n1 && n2 > n3 {
		fmt.Println("O maior deles é ", n2)
	} else if n3 > n1 && n3 > n2 {
		fmt.Println("O maior deles é ", n3)
	}

}
