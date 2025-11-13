package helpers


import (
	"os"
)

func CreateDir(path string) {
	err := os.MkdirAll(path, 0755)
	CheckErr(err)
}