package main

import "fmt"

func main() {
	var (
		n int
	)
	fmt.Println("Digite um número")
	fmt.Scan(&n)

	if n%3 == 0 && n%5 == 0 {
		fmt.Println("Esse número é divisivel por 3 e 5")
	} else {
		fmt.Println("Esse número NÃO é divisivel por 3 e 5")
	}

}
