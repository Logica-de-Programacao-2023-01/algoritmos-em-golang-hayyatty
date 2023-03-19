package main

import "fmt"

func main() {
	var (
		idade float32
	)
	fmt.Print("Escreva a sua idade em anos")
	fmt.Scan(&idade)
	emdias := idade * 365.3
	fmt.Println("A sua idade em dias é ", emdias)
}
