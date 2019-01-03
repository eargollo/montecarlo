package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var baseFile string
var future int

var rootCmd = &cobra.Command{
	Use:   "montecarlo",
	Short: "MonteCarlo applies a Monte Carlo estimation to a sequence of data",
	Long: `MonteCarlo: Given a sequence of measurements, and the amount of future data,
it applies MonteCarlo to predict future results based on past ones.
`,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&baseFile, "input", "./input.csv", "Input data, one value per line")
	rootCmd.PersistentFlags().Int(&future, "future", 12, "Future data points")
}

// Execute runs the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
