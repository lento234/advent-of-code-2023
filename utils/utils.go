package utils

import (
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
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

const (
	Blue    = lipgloss.Color("4")
	Green   = lipgloss.Color("6")
	Grey    = lipgloss.Color("0")
	Red     = lipgloss.Color("9")
	Magenta = lipgloss.Color("5")
	None    = lipgloss.Color("#FFFFFF")
)

func FormatGreen(text string) string {
	style := lipgloss.NewStyle().Bold(true)
	style = style.Foreground(Green)
	return style.Render(text)
}
func FormatMagenta(text string) string {
	style := lipgloss.NewStyle().Bold(true)
	style = style.Foreground(Magenta)
	return style.Render(text)
}
