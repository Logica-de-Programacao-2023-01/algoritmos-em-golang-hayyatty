package main

import "fmt"

func main() {
	var peso, altura float64
	fmt.Print("Qual é o seu peso ?")
	fmt.Scanln(&peso)
	fmt.Print("Qual a sua altura?")
	fmt.Scanln(&altura)
	imc := peso / (altura * altura)
	println("O seu imc é :", imc)
}
