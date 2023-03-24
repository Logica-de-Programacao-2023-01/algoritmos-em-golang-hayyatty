package main

import "fmt"

/*Faça um algoritmo que leia a altura e o sexo de
uma pessoa e mostre se ela está abaixo,
dentro ou acima do peso ideal, utilizando as
fórmulas do exercício 3 da lista de algoritmos sequenciais.*/

func main() {
	var (
		alt  float32
		peso float32
		sex  int
	)
	fmt.Println("Qual a sua altura (em metros)?")
	fmt.Scan(&alt)
	fmt.Println("Qual o seu peso (em quilos)?")
	fmt.Scan(&peso)
	fmt.Println("Qual o seu sexo ?")
	fmt.Println("[1]-masculino")
	fmt.Println("[2]-feminino")
	fmt.Scan(&sex)
	imc := peso / (alt * alt)
	switch sex {
	case 1:
		fmt.Println("Seu imc é", imc)
		var class string
		if imc <= 18.5 {
			class = "Abaixo do peso normal"
			println(class)
		} else if imc >= 18.5 && imc <= 24.9 {
			class = "Peso normal"
			println(class)
		} else if imc >= 25.0 && imc <= 29.9 {
			class = "Excesso de peso"
			println(class)
		} else if imc >= 30.0 && imc <= 34.9 {
			class = "Obesidade GRAU 1"
			println(class)
		} else if imc >= 35.0 && imc <= 39.9 {
			class = "Obesidade GRAU 2 (Severa)"
			println(class)
		} else {
			class = "Obesidade GRAU 3 (Morbida)"
		}
	case 2:
		fmt.Println("Seu imc é", imc)
		var class string
		if imc <= 18.5 {
			class = "Abaixo do peso normal"
			println(class)
		} else if imc >= 18.5 && imc <= 24.9 {
			class = "Peso normal"
			println(class)
		} else if imc >= 25.0 && imc <= 29.9 {
			class = "Excesso de peso"
			println(class)
		} else if imc >= 30.0 && imc <= 34.9 {
			class = "Obesidade GRAU 1"
			println(class)
		} else if imc >= 35.0 && imc <= 39.9 {
			class = "Obesidade GRAU 2 (Severa)"
			println(class)
		} else {
			class = "Obesidade GRAU 3 (Morbida)"
		}

	}
}
