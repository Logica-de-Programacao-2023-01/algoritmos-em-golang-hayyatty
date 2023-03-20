package main

import "fmt"

func main() {
	var (
		dias   float32
		diaria float32
		sal    float32
	)
	fmt.Println("Qual o valor cobrado pela sua diária ?")
	fmt.Scan(&diaria)
	fmt.Println("Quantos dias foram trabalhados ?")
	fmt.Scan(&dias)
	sal = diaria * dias
	fmt.Println("O seu salário será de :", sal, "Reais")

}
