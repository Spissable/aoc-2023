package util

import (
	"fmt"
	"os"
)

func ReadInput(day int) string {
	folder, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/../day%d/input.txt", folder, day)
	content, err := os.ReadFile(filePath)

	if err != nil {
		panic(fmt.Sprintf("File not found for path: %s", err))
	}

	return string(content)
}
