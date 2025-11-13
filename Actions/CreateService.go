package actions

import (
	"Wild/Helpers"
	"fmt"
	"path/filepath"
	"strings"
)



func CreateService(baseDir string, serviceName *string) {

	if strings.Contains(strings.ToLower(*serviceName), "service") {
		fmt.Println("El nombre del hook debe ser sin 'use' y sin terminar en 'Hook'")
		return
	}

	if strings.Contains(strings.ToLower(*serviceName), "use") {
		fmt.Println("El nombre del hook debe ser sin 'use' y sin terminar en 'Hook'")
		return
	}

	//Creamos la carpeta de tipos.
	helpers.CreateDir(filepath.Join(baseDir, "../Types")) //--> Crea la carpeta de tipos si no existe.

	//Creamos los archivos
	helpers.CreateFile(filepath.Join(baseDir, "use"+*serviceName+"service.tsx"), ServiceTemplate(*serviceName))
	helpers.CreateFile(filepath.Join(baseDir, "../Types", "use"+*serviceName+"ServiceTypes.tsx"), ServiceTypesTemplate(*serviceName))
}

func ServiceTemplate(name string) string {
	return fmt.Sprintf(`//Types
import { %[1]sCompProps } from "../Types/use%[1]sHookTypes";

export function use%[1]sHook({}: %[1]sCompProps) {
	return (
		<></>
	);
}
`, name)
}

func ServiceTypesTemplate(name string) string {
	return fmt.Sprintf(`export interface use%[1]sHookProps {

}
`, name)
}