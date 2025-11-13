package flags

import "flag"

func InitFlags() *string {

	CompName := flag.String("comp", "", "crear un componente --comp=Nombre-componente que se va a crear")

	flag.Parse()

	return CompName
}
