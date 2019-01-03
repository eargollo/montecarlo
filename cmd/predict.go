package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var baseFile string
var future int

func init() {
	rootCmd.AddCommand(estimateCmd)
	estimateCmd.PersistentFlags().StringVar(&baseFile, "input", "./input.csv", "Input data, one value per line")
	estimateCmd.PersistentFlags().IntVar(&future, "future", 12, "Future data points")
}

var estimateCmd = &cobra.Command{
	Use:   "estimate",
	Short: "Simulate future estimation",
	Long:  `Estimate towards the future based on past data`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("I did run!")
	},
}
