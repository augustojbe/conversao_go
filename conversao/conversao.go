package main

import "fmt"

func main() {

	fmt.Println("Digite o valor que você gostaria de transformar de Kelvin para Celsius")
	var temperaturaK int

	fmt.Scan(&temperaturaK)

	temperaturaC := temperaturaK - 273

	fmt.Println("A temperatura do valor escolhido em Celcius é: ", temperaturaC)

}
