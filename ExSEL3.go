package main

import "fmt"

func main() {
	var (
		n int
	)
	fmt.Println("Digite o número")
	fmt.Scan(&n)
	div := n % 2
	if div == 0 {
		fmt.Println("O número", n, "é par")
	} else if div != 0 {
		fmt.Println("O número", n, "é impar")
	}
}
