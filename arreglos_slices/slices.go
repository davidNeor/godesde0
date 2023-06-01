package arreglos_slices

import "fmt"

var tablaS []int=[]int{2,5,4}
var arreglo [10]int=[10]int{5,7,12,323,43,123,765,8912,1213}
func MuestroSlice(){

	porcion:=arreglo[3:] //slice de un vector desde la posicion 3
	porcion2:=arreglo[:5]//slice desde la posicion 0 hasta la 5
	porcion3:=arreglo[3:9]//slice desde la posicion 3 hasta la 9

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {

	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos),cap(elementos))	
	nums:= make([]int,0,0)
	for i:=0;i<100;i++ {
		nums=append(nums, i)
	}
	fmt.Printf("\nLargo %d,Capacidad %d", len(nums), cap(nums))
}
