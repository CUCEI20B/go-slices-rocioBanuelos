package main

import "fmt"

func main() {
	var cantidadElemento uint64
	var numero int64
	var resultado int64

	//Se obtiene el número "n" de elementos del slice
	fmt.Scanf("%d\n", &cantidadElemento)
	//Creación del slice
	s := make([]int64, 0, cantidadElemento)
	//Leer valores del slice
	for i := 0; i < int(cantidadElemento); i++ {
		fmt.Scanf("%d\n", &numero)
		s = append(s,numero)
	}
	//Realizar suma de elementos
	for _, v := range s {
		resultado += v
	}
	//Imprimir resultado de la suma
	fmt.Println(resultado)	
}