package main

import "fmt"

func main() {
	var (
		peso  float32
		pesol float32
	)
	fmt.Println("Digite o seu peso em quilos, para ver em libras")
	fmt.Scan(&peso)
	pesol = peso * 2.2046
	fmt.Println("O seu peso em libras é :", pesol)
}
