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

	switch os2:=runtime.GOOS; os2{
	case "linux": 
		fmt.Println("Esto es",os2)		
	case "darwin":
		fmt.Println("Esto es",os2)		
	default:
		fmt.Printf("%s \n",os2)
		


	}
}