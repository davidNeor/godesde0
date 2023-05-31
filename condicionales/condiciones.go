package condicionales

import("fmt"
		"runtime"
)

func Micondicion() { 
	
	if os:= runtime.GOOS; os=="Linux" ||os=="OS X."{
		fmt.Println("Tu sistema no es windows")
	}else{
		fmt.Println("Tu sistema es ",os)
	}

}