package archivos

import(
	"fmt"
	// "bufio"
	"os"
	"github.com/davidNeor/godesde0/ejercicios"
	// "ioutil"
)


func GrabaTabla(){

	
	texto:=ejercicios.Tabla()
	archivo,err:=os.Create("./archivos/txt/Tablas.txt")
	if err!=nil {
		println("Error al crear el archivo")
		return
	}
	fmt.Fprintln(archivo,texto)
	archivo.Close()
}

