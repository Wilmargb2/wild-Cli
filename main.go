package main

import (
	"Wild/Flags"
	"Wild/Actions"
	"fmt"
	"os"
)

func main() {

	dir, _ := os.Getwd()
	Flags := flags.InitFlags()

	switch {
		case Flags.Make && Flags.CompName != "":
			actions.CreateComp(dir, &Flags.CompName)

		case Flags.Make && Flags.HookName != "":
			actions.CreateHook(dir, &Flags.HookName)
		
		case Flags.Make && Flags.ServiceName != "":
			actions.CreateService(dir, &Flags.ServiceName)
			
		default:
			fmt.Println("Comando no reconocido usa --help para ver la ayuda")
	}

}
