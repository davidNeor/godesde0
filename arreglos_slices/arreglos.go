package arreglos_slices

import "fmt"


var tabla [10]int=[10]int{10,0,67,89}
func MuestroArreglos() {

	tabla2 := [10]string{"David","Alberto"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i:=0;i<len(tabla);i++{
		fmt.Println(tabla[i])
	}

	
}