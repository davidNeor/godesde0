package ejercicios

import(
		"strconv"
)

func Primera(parametro string)(int,string,error) {

	miEntero,mierror:=strconv.Atoi(parametro)

	
	
	if miEntero>100 {		
		return miEntero,"Es mayor a  100",mierror
	} else{
		return miEntero,"Es menor a 100",mierror
	}
}