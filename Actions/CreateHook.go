package actions

import (
	"Wild/Helpers"
	"fmt"
	"path/filepath"
	"strings"
)



func CreateHook(baseDir string, hookName *string) {

	if strings.Contains(strings.ToLower(*hookName), "hook") {
		fmt.Println("El nombre del hook debe ser sin 'use' y sin terminar en 'Hook'")
		return
	}

	if strings.Contains(strings.ToLower(*hookName), "use") {
		fmt.Println("El nombre del hook debe ser sin 'use' y sin terminar en 'Hook'")
		return
	}

	//Creamos la carpeta de tipos.
	helpers.CreateDir(filepath.Join(baseDir, "../Types")) //--> Crea la carpeta de tipos si no existe.

	//Creamos los archivos
	helpers.CreateFile(filepath.Join(baseDir, "use"+*hookName+"Hook.tsx"), HookTemplate(*hookName))
	helpers.CreateFile(filepath.Join(baseDir, "../Types", "use"+*hookName+"HookTypes.tsx"), HookTypesTemplate(*hookName))
}

func HookTemplate(name string) string {
	return fmt.Sprintf(`//Types
import { %[1]sCompProps } from "../Types/use%[1]sHookTypes";

export function use%[1]sHook({}: %[1]sCompProps) {
	return (
		<></>
	);
}
`, name)
}

func HookTypesTemplate(name string) string {
	return fmt.Sprintf(`export interface use%[1]sHookProps {

}
`, name)
}