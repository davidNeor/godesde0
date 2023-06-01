package defer_panic

import (
	"os"
	"log"
	"fmt"
)

func VemosDefer() {

	fmt.Println("Este es un primer mensaje")
	defer archivo.Close()
}