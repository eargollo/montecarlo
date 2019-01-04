package simulation

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// New initializes a simulation object
func New(inputFile string, future int) (Simulation, error) {
	inputData, err := readDataFile(inputFile)
	if err != nil {
		return Simulation{}, err
	}
	return NewWithData(inputData, future, 1000000), err
}

func NewWithData(inputData []float64, future int, simulations int) Simulation {
	return Simulation{inputData: &inputData, future: future, simulations: simulations}
}

func readDataFile(fileName string) ([]float64, error) {
	data := []float64{}
	// Open CSV file
	f, err := os.Open(fileName)
	if err != nil {
		return data, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return data, err
	}

	// Loop through lines & turn into object
	for _, line := range lines {
		fl, err := strconv.ParseFloat(line[0], 32)
		if err == nil {
			data = append(data, fl)
		} else {
			fmt.Printf("Data %v is not a float number and will be ignored.", line[0])
		}
	}
	return data, nil
}

// Simulation represents a MonteCarlo simulation
type Simulation struct {
	inputData   *[]float64
	future      int
	simulations int
}

type SimulationData struct {
	future []float64
}

func (s *Simulation) Run() []SimulationData {
	data := []SimulationData{}
	for i := 0; i < s.simulations; i++ {

		data = append(data, s.singleMonteCarlo())
	}
	return data
}

func (s *Simulation) singleMonteCarlo() SimulationData {
	data := SimulationData{}
	totalInput := len(*s.inputData)
	for i := 0; i < s.future; i++ {
		item := rand.Intn(totalInput)
		data.future = append(data.future, (*s.inputData)[item])
	}
	return data
}
