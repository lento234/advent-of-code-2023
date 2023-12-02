package utils

import (
	"os"
	"strings"

	"github.com/charmbracelet/log"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ParseFile(filename string) []string {
	buf, err := os.ReadFile(filename)
	CheckErr(err)
	return ParseString(string(buf))
}

func ParseString(text string) []string {
	input := strings.Split(strings.TrimSuffix(text, "\n"), "\n")
	return input
}
