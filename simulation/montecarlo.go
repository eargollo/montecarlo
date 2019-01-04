package simulation

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
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
	return Simulation{InputData: &inputData, Future: future, Simulations: simulations}
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
	InputData      *[]float64
	Future         int
	Simulations    int
	ForecastPoints int
	Data           []SimulationData
	Forecasts      []Forecast // One forecast per percentil. If decided increments in 10%, there will be one forecast per each 10%. Each forecast will have a dataset per each future time
}

type SimulationData struct {
	Future    []float64
	SumFuture []float64
}

type Forecast struct {
	Percentil float64
	Forecast  []float64
}

func (s *Simulation) generateData() {
	s.Data = []SimulationData{}
	for i := 0; i < s.Simulations; i++ {
		s.Data = append(s.Data, s.singleMonteCarlo())
	}
}

func (s *Simulation) aggregateFutureData() {
	for i, item := range s.Data {
		s.Data[i].SumFuture = []float64{}
		sum := 0.0
		for _, fut := range item.Future {
			sum += fut
			s.Data[i].SumFuture = append(s.Data[i].SumFuture, sum)
		}
	}
}

func (s *Simulation) calculateForecasts() {
	s.Forecasts = []Forecast{}
	var step float64
	step = 100.0 / (float64(s.ForecastPoints) - 1)
	for i := 0; i < s.ForecastPoints; i++ {
		f := Forecast{Percentil: 100.0 - (float64(i) * step)}
		// Calculate the element in a sorted array that would represent the minimum on the percentil
		point := int((100 - f.Percentil) * float64(s.Simulations-1) / 100)
		for j := 0; j < s.Future; j++ {
			// Sort the array and get the element
			data := s.Data
			sort.Slice(data, func(t, r int) bool { return data[t].SumFuture[j] < data[r].SumFuture[j] })
			f.Forecast = append(f.Forecast, s.Data[point].SumFuture[j])
		}
		s.Forecasts = append(s.Forecasts, f)
	}
}

func (s *Simulation) analyze() {
	s.aggregateFutureData()
	s.calculateForecasts()
}

func (s *Simulation) Run() {
	s.generateData()
	s.analyze()
}

func (s *Simulation) singleMonteCarlo() SimulationData {
	data := SimulationData{}
	totalInput := len(*s.InputData)
	for i := 0; i < s.Future; i++ {
		item := rand.Intn(totalInput)
		data.Future = append(data.Future, (*s.InputData)[item])
	}
	return data
}
