package goroutines

import(
	"fmt"
	"strings"
	"time"
)

func MiNombreLentoo(nombre string) {

	letras:= strings.Split(nombre,"")
	for _, letra:= range letras{
		time.Sleep(1000 * time.Millisecond)
	}
}