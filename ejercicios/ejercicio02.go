package ejercicios

import(
	"fmt"
	"bufio"
	"os"
	"strconv"

)
var err error
var numero1 int
var texto string

func Tabla() string { 
	escaner:=bufio.NewScanner(os.Stdin)

	
	for{
		fmt.Println("Introduce un numero por teclado: ")
		if escaner.Scan(){ 
			numero1,err=strconv.Atoi(escaner.Text())
			if err==nil {
				break
			}
		}else{
			fmt.Println("Vuelva a introducir un numero por teclado: ")
		}
	}
	for i:=1;i<=10;i++{
		texto+=fmt.Sprintf("%d x %d = %d \n",i,numero1,(i*numero1))
	}
	return texto

}
