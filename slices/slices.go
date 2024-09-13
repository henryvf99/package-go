package slices

import "fmt"

// los identificadores con letra mayuscula pueden ser exportados
// y los que usan minusculas solo dentro del paquete se pueden usar
func sum(a, b int) {
	fmt.Println(a + b)
}

func Res(a, b int) {
	fmt.Println(a + b)
}
