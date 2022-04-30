/*
	Los packages se organizan en carpetas, dos packages no pueden estar en la misma carpeta.
	El archivo go.mod indica el nombre del modulo. Este modulo, al que pertenecen todos los packages,
	se llama: Testing, por lo que para importar un package de este modulo es necesario importarlo como:
	Testing/greeting, por ejemplo, ya que este package es greeting... tambien es como referirse a las carpetas
	desde la ubicacion del modulo que esta en el top del directorio.
*/

package greeting

import "fmt"

func Greeting() {
	fmt.Println("From a diferent package")
}

//tipos de variables
func sum(x int, y int) int {
	return x + y
}

func compare(x int, y int) bool {
	return x > y
}
