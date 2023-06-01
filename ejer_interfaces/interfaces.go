package ejer_interfaces

import(
	"github.com/davidNeor/godesde0/interfaces"
	"fmt"

)

func HumanosRespirando(hu interfaces.Humano){
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n",hu.Sexo())
}

