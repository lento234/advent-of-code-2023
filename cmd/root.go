package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var day int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc2023",
	Short: "Advent of Code 2023",
	Long:  "Solutions to Advent of Code",
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(xmasCmd)

	rootCmd.AddCommand(initCmd)
	initCmd.Flags().IntVar(&day, "day", 0, "Day to initialize")
	initCmd.MarkFlagRequired("day")
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

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
