package main

import (
	"Wild/Helpers"
	"Wild/Flags"
	"fmt"
	"os"
	"strings"
)

func main() {

	dir, _ := os.Getwd()
	CompName := flags.InitFlags()

	if strings.TrimSpace(*CompName) == "" {
		fmt.Println("El nombre del componente es obligatorio y no puede estar vacio")
		return
	}

	helpers.CreateComp(dir, CompName)

}
