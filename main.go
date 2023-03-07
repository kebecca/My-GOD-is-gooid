package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var fileContent []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		fileContent = append(fileContent, sc.Text())
	}

	text := os.Args[1]
	textLines := strings.Split(text, "\\n")
	fmt.Print(stylizedText(textLines, fileContent))
}

func stylizedText(textLines []string, fileContent []string) string {
	var sb strings.Builder
	for _, letter := range textLines {
		for j := 0; j < 8; j++ {
			for _, char := range letter {
				pos := 1 + (int(char)-32)*9 + j
				sb.WriteString(fileContent[pos])
			}
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
