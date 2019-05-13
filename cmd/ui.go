package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uiCmd)
}

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Opens UI interface",
	Long:  `Upens UI interface with default settings for estimation`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Abacaxi...\n")
		http.HandleFunc("/", homePage)
		http.HandleFunc("/all", returnResults)
		log.Fatal(http.ListenAndServe(":8081", nil))
	},
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// SimulationResult12 is plain 12 future datapoints of results for a simulation
// the goal is to simplify processing of JavaScript at a first version where UI
// will not be able to differ on simulation parameters
type SimulationResult12 struct {
	Conf     float64
	Future1  float64
	Future2  float64
	Future3  float64
	Future4  float64
	Future5  float64
	Future6  float64
	Future7  float64
	Future8  float64
	Future9  float64
	Future10 float64
	Future11 float64
	Future12 float64
}

// SimulationResults12 is a list of results
type SimulationResults12 []SimulationResult12

func returnResults(w http.ResponseWriter, r *http.Request) {
	articles := SimulationResults12{
		SimulationResult12{Conf: 100, Future1: 1, Future2: 2},
		SimulationResult12{Conf: 90, Future1: 3, Future2: 4},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}
