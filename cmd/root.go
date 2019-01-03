package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "montecarlo",
	Short: "MonteCarlo applies a Monte Carlo estimation to a sequence of data",
	Long: `MonteCarlo: Given a sequence of measurements, and the amount of future data,
it applies MonteCarlo to predict future results based on past ones.
`,
}

// Execute runs the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
