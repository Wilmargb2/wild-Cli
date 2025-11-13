package helpers

import (
	"fmt"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func CreateComp(dir string, CompName *string) {
	

	os.Mkdir(dir+"/"+*CompName, 0755)
	os.Mkdir(dir+"/"+*CompName+"/Comp", 0755)
	os.Mkdir(dir+"/"+*CompName+"/Container", 0755)
	os.Mkdir(dir+"/"+*CompName+"/Hooks", 0755)
	os.Mkdir(dir+"/"+*CompName+"/Services", 0755)
	os.Mkdir(dir+"/"+*CompName+"/Styles", 0755)
	os.Mkdir(dir+"/"+*CompName+"/Types", 0755)

	Comp, err := os.Create(dir + "/" + *CompName + "/Comp/" + *CompName + "Comp.tsx")
	checkErr(err)

	Container, err := os.Create(dir + "/" + *CompName + "/Container/" + *CompName + "Container.tsx")
	checkErr(err)

	Css, err := os.Create(dir + "/" + *CompName + "/Styles/" + *CompName + ".css")
	checkErr(err)

	ContainerTypes, err := os.Create(dir + "/" + *CompName + "/Types/" + *CompName + "ContainerTypes.tsx")
	checkErr(err)
	
	CompTypes, err := os.Create(dir + "/" + *CompName + "/Types/" + *CompName + "CompTypes.tsx")
	checkErr(err)

	defer Comp.Close()
	defer Container.Close()
	defer Css.Close()
	defer ContainerTypes.Close()
	defer CompTypes.Close()

Comp.WriteString(fmt.Sprintf(
`//Types
import { %[1]sCompProps } from "../Types/%[1]sCompTypes";

//Css
import "../Styles/%[1]s.css";

export function %[1]sComp({}: %[1]sCompProps) {
	return (
		<></>
	);
}`,
*CompName))

Container.WriteString(fmt.Sprintf(
`//Types
import { %[1]sContainerProps } from "../Types/%[1]sContainerTypes";

//Component
import { %[1]sComp } from "../Comp/%[1]sComp";

export function %[1]sContainer({}: %[1]sContainerProps) {
	return (
		<%[1]sComp />
	);
}`,
*CompName))

ContainerTypes.WriteString(fmt.Sprintf(
`export interface %[1]sContainerProps {

}`,
*CompName))

CompTypes.WriteString(fmt.Sprintf(
`export interface %[1]sCompProps {

}`,
*CompName))


}