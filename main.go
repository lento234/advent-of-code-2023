package main

import (
	"aoc2023/cmd"
	"aoc2023/utils"
	"os"

	"github.com/charmbracelet/lipgloss"
	_log "github.com/charmbracelet/log"
)

var log = _log.New(os.Stderr)

func initLogger() {
	styles := _log.DefaultStyles()
	styles.Levels[_log.InfoLevel] = lipgloss.NewStyle().
		Bold(true).
		SetString("[INFO]").
		Foreground(utils.Green)

	styles.Levels[_log.FatalLevel] = lipgloss.NewStyle().
		Bold(true).
		SetString("[FATAL]").
		Background(utils.Red)

	_log.SetReportTimestamp(false)

	_log.SetStyles(styles)
}

func main() {
	initLogger()
	cmd.Execute()
}
