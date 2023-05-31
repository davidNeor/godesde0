//esta es la estructura convencional de un fichero go
package main

import ("fmt"
	"github.com/davidNeor/godesde0/variables"
	"github.com/davidNeor/godesde0/condicionales")

func main(){
	variables.MuestroEnteros()
	mibooleano,mitexto:=variables.ConviertoTexto(1677)
	fmt.Println(mibooleano)
	fmt.Println(mitexto)
	condicionales.Micondicion()

}