package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
	"github.com/ChickenWhisky/fmtDB/internal/processor"
)

var rootCmd = &cobra.Command{
    Use:   "csv-frequency-counter [input files...]",
    Short: "A CLI tool to compute frequency counts from CSV files",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 0 {
            fmt.Println("Please provide at least one CSV file as an argument.")
            os.Exit(1)
        }
        output := "output.csv"
        if err := processor.ProcessCSVFiles(args, output); err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(1)
        }
        fmt.Printf("Frequency counter CSV created: %s\n", output)
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
