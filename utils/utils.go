package utils

import (
	"fmt"
	"io"
	"os"
)

func ReadFromFile(fileLocation string) string {
	wd, _ := os.Getwd()
	file, fileOpenError := os.Open(wd + fileLocation)

	if fileOpenError != nil {
		fmt.Println("Error while opening file")
		panic(fileOpenError)
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}

	return string(content)
}
