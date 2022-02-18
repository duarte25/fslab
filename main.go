package main

import "fmt"

func main() {
	var n1, n2 float64

	fmt.Scanf("%f", &n1)
	fmt.Scanf("%f", &n2)

	fmt.Println("soma:", soma(n1, n2))

	fmt.Println("Média aritmética:", mediaAritmetica(n1, n2))

	fmt.Println("Média ponderada:", mediaPonderada(n1, n2))

}

func soma(n1, n2 float64) float64 {
	return n1 + n2
}

func mediaAritmetica(n1, n2 float64) float64 {
	return (n1 + n2) / 2
}

func mediaPonderada(n1, n2 float64) float64 {
	return ((1 * n1) + (2 * n2)) / 3
}
