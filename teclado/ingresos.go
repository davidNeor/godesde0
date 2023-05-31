package teclado

import("fmt"
		"os"
		"bufio"
		"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error
func IngresoNumeros() {
	escaner:=bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese número 1: ")	
	if escaner.Scan() {
		numero1,err=strconv.Atoi(escaner.Text())
		if err!=nil {
			panic("El dato ingresado es incorrecto "+err.Error())
		}
	}
	fmt.Println("Ingrese número 2: ")	
	if escaner.Scan() {
		numero2,err=strconv.Atoi(escaner.Text())
		if err!=nil {
			panic("El dato ingresado es incorrecto "+err.Error())
		}
	}
	fmt.Println("Ingrese leyenda: ")	
	if escaner.Scan() {
		leyenda=escaner.Text()
			
	}

	fmt.Println(leyenda,numero1+numero2)
}