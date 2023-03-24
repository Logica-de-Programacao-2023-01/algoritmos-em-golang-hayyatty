package main

import "fmt"

func main() {
	var (
		soma int
		num  int
	)
	cont := 0
	for {
		fmt.Println("Digite um número (0 para parar)")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		soma += num
		cont++
	}
	fmt.Println("A média é :", soma/cont)

}
