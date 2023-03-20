package main

import "fmt"

func main() {
	var (
		sal float32
	)
	fmt.Println("Você ganhou um aumento, digite o seu salário:")
	fmt.Scan(&sal)
	nsal := sal + (sal * 0.15)
	fmt.Println("O seu novo salário será:", nsal, "Reais")
}
