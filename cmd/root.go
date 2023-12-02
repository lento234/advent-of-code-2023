package cmd

import (
	"aoc2023/day01"
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
	func() error {
		return fmt.Errorf("Not yet solved!")
	},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc2023",
	Short: "Advent of Code 2023",
	Long:  "Solutions to Advent of Code",
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
			return fmt.Errorf("Day %d is not an advent day.", day)
		}
		// fmt.Printf("Initializing day %d\n", day)
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
			return fmt.Errorf("Day %d is not an advent day.", day)
		}
		// Calling challenge of the day
		f := reflect.ValueOf(challenges[day-1])
		res := f.Call([]reflect.Value{})
		err := res[0].Interface()
		if err != nil {
			log.Fatal(err)
		}

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
	initCmd.MarkFlagRequired("day")

	// Run the challenge of the day
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVar(&day, "day", 0, "Day to initialize")
	runCmd.MarkFlagRequired("day")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
