package main

import "fmt"

func main() {
	var (
		n          int
		antecessor int
		sucessor   int
	)
	fmt.Println("Digite um número")
	fmt.Scan(&n)
	antecessor = n - 1
	sucessor = n + 1
	fmt.Println("O antecessor desse número é ", antecessor)
	fmt.Println("O sucessor desse número é ", sucessor)

}
