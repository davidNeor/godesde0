package variables

import ("fmt"
		"time"
		"strconv")

var lkoque string
var fecha time.Time
var estado bool
var sueldo float32
func RestoVariables() {

	lkoque="David"
	fecha=time.Now()
	estado=true
	sueldo=1600.56
	
	fmt.Println(lkoque)
	fmt.Println(fecha)
	fmt.Println(estado)
	fmt.Println(sueldo)

}

func ConviertoTexto(numero int)(bool,string){
	
	texto:=strconv.Itoa(numero)
	return true,texto
}

