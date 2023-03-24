package main

import "fmt"

func main() {
	var (
		n int
	)
	fmt.Println("Digite um número")
	fmt.Scan(&n)
	for i := n; i < n*10; i = i + n {
		fmt.Println(i)
	}
}
