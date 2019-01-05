package cmd

import (
	"fmt"
	"os"

	"github.azc.ext.hp.com/eduardo/montecarlo/simulation"
	"github.com/spf13/cobra"
)

var baseFile string
var future, simulations int
var increment float64
var csv bool

func init() {
	rootCmd.AddCommand(estimateCmd)
	estimateCmd.PersistentFlags().StringVar(&baseFile, "input", "input.txt", "Input data, one value per line")
	estimateCmd.PersistentFlags().IntVar(&future, "future", 12, "Future data points")
	estimateCmd.PersistentFlags().IntVar(&simulations, "simulations", 1000000, "Amount of MonteCarlo simulations used (the bigger the number, the better the precision but it may take longer to simulate).")
	estimateCmd.PersistentFlags().Float64Var(&increment, "increment", 5, "Percentual increment for each confidence data point. Default is 5, i.e. one data point for each 5%: 100%, 95%, 90%,...0%")
	estimateCmd.PersistentFlags().BoolVar(&csv, "csv", false, "Output data as a csv table")
}

var estimateCmd = &cobra.Command{
	Use:   "estimate",
	Short: "Simulate future estimation",
	Long:  `Estimate towards the future based on past data`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(baseFile); os.IsNotExist(err) {
			// File does not exist
			fmt.Printf("File %v could not be found. Make sure either the default input.txt file exists or pass your file with the option '--input my-file.txt'", baseFile)
			return
		}
		sim, err := simulation.New(baseFile, future, simulations, increment)
		if err != nil {
			panic(err)
		}
		sim.Run()
		sim.ForecastStdout(csv)
	},
}
