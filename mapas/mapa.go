package mapas

import "fmt"


func MostrarMapas() {
	paises := make(map[string]string)
	paises["mexico"]="D.F"
	paises["argentina"]="buenos aires"
	fmt.Println(paises)
	fmt.Println(paises["argentina"])

	campeonato:=map[string]int{
		"Real Madrid":20,
		"Valencia":15,
		"Sevilla":12,
	}
	campeonato["Varcelona"]=10

	for equipo,puntos:=range campeonato{

		fmt.Printf("\n Equipo: %s puntos: %d",equipo,puntos)
	}

}