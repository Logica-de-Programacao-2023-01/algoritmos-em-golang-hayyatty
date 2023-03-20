package main

//Faça um algoritmo que leia um número inteiro e mostre o seu dobro, triplo e quadruplo.
import (
	"fmt"
)

func main() {
	var (
		n1 int64
	)
	fmt.Println("Digite um número :")
	fmt.Scanln(&n1)
	x2 := n1 * 2
	fmt.Println("O dobro desse número é ", x2)
	x3 := n1 * 3
	fmt.Println("O triplo desse número é ", x3)
	x4 := n1 * 4
	fmt.Println("O quadruplo desse número é", x4)
}
