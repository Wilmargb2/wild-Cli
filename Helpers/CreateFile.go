package helpers

import (
	"os"
)

func CreateFile(path, content string) {
	file, err := os.Create(path)
	CheckErr(err)
	defer file.Close()

	if content != "" {
		_, err = file.WriteString(content)
		CheckErr(err)
	}
}