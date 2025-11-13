package actions

import (
	"fmt"
	"path/filepath"
	"Wild/Helpers"
)

// --- Estructura principal ---

func CreateComp(baseDir string, compName *string) {
	root := filepath.Join(baseDir, *compName)

	dirs := []string{
		"Comp",
		"Container",
		"Hooks",
		"Services",
		"Styles",
		"Types",
	}

	// Crear carpetas
	for _, d := range dirs {
		helpers.CreateDir(filepath.Join(root, d))
	}

	// --- Archivos a crear ---
	files := map[string]string{
		filepath.Join(root, "Comp", *compName+"Comp.tsx"): compTemplate(*compName),
		filepath.Join(root, "Container", *compName+"Container.tsx"): containerTemplate(*compName),
		filepath.Join(root, "Styles", *compName+".css"):             "",
		filepath.Join(root, "Types", *compName+"ContainerTypes.tsx"): containerTypesTemplate(*compName),
		filepath.Join(root, "Types", *compName+"CompTypes.tsx"):      compTypesTemplate(*compName),
	}

	for path, content := range files {
		helpers.CreateFile(path, content)
	}
}

// --- Templates ---

func compTemplate(name string) string {
	return fmt.Sprintf(`//Types
import { %[1]sCompProps } from "../Types/%[1]sCompTypes";

//Css
import "../Styles/%[1]s.css";

export function %[1]sComp({}: %[1]sCompProps) {
	return (
		<></>
	);
}
`, name)
}

func containerTemplate(name string) string {
	return fmt.Sprintf(`//Types
import { %[1]sContainerProps } from "../Types/%[1]sContainerTypes";

//Component
import { %[1]sComp } from "../Comp/%[1]sComp";

export function %[1]sContainer({}: %[1]sContainerProps) {
	return (
		< %[1]sComp />
	);
}
`, name)
}

func containerTypesTemplate(name string) string {
	return fmt.Sprintf(`export interface %[1]sContainerProps {

}
`, name)
}

func compTypesTemplate(name string) string {
	return fmt.Sprintf(`export interface %[1]sCompProps {

}
`, name)
}
