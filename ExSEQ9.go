package main

import "fmt"

func main() {
	var (
		preco float32
		pdes  float32
	)
	fmt.Println("Estamos dando 10% de desconto na loja toda")
	fmt.Println("Digite o preço do produto para ver ele com desconto")
	fmt.Scan(&preco)
	pdes = preco - (preco * 0.1)
	fmt.Println("O valor desse item será de :", pdes, "Reais")
}
