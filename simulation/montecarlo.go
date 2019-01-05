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
func New(inputFile string, future int, simulations int, increment float64) (Simulation, error) {
	inputData, err := readDataFile(inputFile)
	if err != nil {
		return Simulation{}, err
	}
	return NewWithData(inputData, future, simulations, increment), err
}

// NewWithData initializes a simulation object with all its details
func NewWithData(inputData []float64, future int, simulations int, increment float64) Simulation {
	forecastPoints := int(100/increment) + 1
	return Simulation{InputData: &inputData, Future: future, Simulations: simulations, ForecastPoints: forecastPoints}
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
	Data           []Data
	Forecasts      []Forecast // One forecast per percentil. If decided increments in 10%, there will be one forecast per each 10%. Each forecast will have a dataset per each future time
}

// Data is a single random simultaion datapoint
type Data struct {
	Future    []float64
	SumFuture []float64
}

// Forecast represents one forecast for all the future points of a percentil
type Forecast struct {
	Percentil float64
	Forecast  []float64
}

func (s *Simulation) generateData() {
	fmt.Println("Generating randomized data...")
	s.Data = []Data{}
	for i := 0; i < s.Simulations; i++ {
		s.Data = append(s.Data, s.singleMonteCarlo())
	}
}

func (s *Simulation) aggregateFutureData() {
	fmt.Println("Aggregating future data...")
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
	fmt.Println("Calculating forecasts...")
	s.Forecasts = []Forecast{}
	var step float64
	step = 100.0 / (float64(s.ForecastPoints) - 1)
	var points []int

	for i := 0; i < s.ForecastPoints; i++ {
		// Calculate the element in a sorted array that would represent the minimum on the percentil
		point := int(float64(i) * step * float64(s.Simulations-1) / 100)
		points = append(points, point)
		// Initialize the forecast for this percentil
		f := Forecast{Percentil: 100.0 - (float64(i) * step)}
		s.Forecasts = append(s.Forecasts, f)
	}

	// Now calculate
	data := s.Data
	fmt.Print("   ")
	for j := 0; j < s.Future; j++ {
		fmt.Printf(" %v...", j)
		// Sort the array and get the points in it
		sort.Slice(data, func(t, r int) bool { return data[t].SumFuture[j] < data[r].SumFuture[j] })
		for i, point := range points {
			s.Forecasts[i].Forecast = append(s.Forecasts[i].Forecast, data[point].SumFuture[j])
		}
	}
	fmt.Println("Done.")
}

func (s *Simulation) analyze() {
	s.aggregateFutureData()
	s.calculateForecasts()
}

func (s *Simulation) Run() {
	s.generateData()
	s.analyze()
}

func (s *Simulation) singleMonteCarlo() Data {
	data := Data{}
	totalInput := len(*s.InputData)
	for i := 0; i < s.Future; i++ {
		item := rand.Intn(totalInput)
		data.Future = append(data.Future, (*s.InputData)[item])
	}
	return data
}

// ForecastStdout prints out a report on the standard output (it is a tab separated report)
func (s *Simulation) ForecastStdout() {
	fmt.Printf("FuturePoints\t%v\tSimulations\t%v\n", s.Future, s.Simulations)
	fmt.Printf("Confidence%%")
	for i := 0; i < s.Future; i++ {
		fmt.Printf("\t%v", i+1)
	}
	fmt.Printf("\n")
	for i := 0; i < s.ForecastPoints; i++ {
		fmt.Printf("%v%%", s.Forecasts[i].Percentil)
		for j := 0; j < s.Future; j++ {
			fmt.Printf("\t%v", s.Forecasts[i].Forecast[j])
		}
		fmt.Printf("\n")
	}
}
