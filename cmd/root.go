package cmd

import (
	"aoc2023/day01"
	"aoc2023/day02"
	"aoc2023/day03"
	"aoc2023/day04"
	"aoc2023/day05"
	"aoc2023/day06"
	"aoc2023/day07"
	"aoc2023/day08"
	"aoc2023/day09"
	"aoc2023/day10"
	"aoc2023/day11"
	"aoc2023/day12"
	"aoc2023/utils"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var day int

var challenges = []interface{}{
	day01.Solve,
	day02.Solve,
	day03.Solve,
	day04.Solve,
	day05.Solve,
	day06.Solve,
	day07.Solve,
	day08.Solve,
	day09.Solve,
	day10.Solve,
	day11.Solve,
	day12.Solve,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc2023",
	Short: "🎅🎄 Advent of Code 2023 ☃️❄️",
}

var xmasCmd = &cobra.Command{
	Use:   "xmas",
	Short: "Is it christmas?",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		xmasDay := time.Date(2023, 12, 25, 0, 0, 0, 0, time.Local)

		if now.Before(xmasDay) {
			fmt.Printf("It's %d days till christmas!\n", int(xmasDay.Sub(now).Hours()/24))
		} else {
			fmt.Println("Jingle bells, Jingle bells...")
		}
	},
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new day",
	RunE: func(cmd *cobra.Command, args []string) error {
		if day <= 0 || day > 24 {
			return fmt.Errorf("day %d is not an advent day", day)
		}
		err := initDay(day)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the challenge of the day",
	RunE: func(cmd *cobra.Command, args []string) error {
		if day <= 0 || day > 24 {
			return fmt.Errorf("day %d is not an advent day", day)
		}
		// Calling challenge of the day
		if day > len(challenges) {
			log.Fatalf("Day %d not solved!", day)
		}
		f := reflect.ValueOf(challenges[day-1])

		// Solve challenge
		fmt.Println(utils.FormatMagenta(fmt.Sprintf("Challenge day %d:", day)))
		f.Call([]reflect.Value{})

		return nil
	},
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	// Easter egg
	rootCmd.AddCommand(xmasCmd)

	// Initizalize the day
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().IntVar(&day, "day", 0, "Day to initialize")
	err := initCmd.MarkFlagRequired("day")
	if err != nil {
		log.Fatal(err)
	}

	// Run the challenge of the day
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVar(&day, "day", 0, "Day to initialize")
	err = runCmd.MarkFlagRequired("day")
	if err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
