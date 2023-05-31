package variables

import "fmt"


func MuestroEnteros(){

	var entero int = 10
	intde32 := int32(100)
	intde64 := int64(90)

	fmt.Println("Primera variable",entero)
	fmt.Println("Segunda variable",intde32)
	fmt.Println("Tercera variable",intde64)

}