package simulation

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	expected := []float64{5, 6, 7, 4, 5, 12, 1, 3, 5, 7, 8}
	data, err := readDataFile("testdata/multidata.csv")
	if err != nil {
		t.Errorf("Failed reading data: %v", err)
	}

	if !reflect.DeepEqual(expected, data) {
		t.Errorf("Arrays are different. Expected %v, got %v", expected, data)
	}
}

func TestRunWithSingleData(t *testing.T) {
	sim := NewWithData([]float64{10}, 12, 100, 21)
	sim.Run()
	res := sim.Data
	if len(res) != 100 {
		t.Errorf("Simulation did not return the 100 expected results. It returned only %v", len(res))
	}
	for _, item := range res {
		if len(item.Future) != 12 {
			t.Errorf("Simulation item has %v future instead of 12", len(item.Future))
			return
		}
		for _, fut := range item.Future {
			if fut != 10 {
				t.Errorf("Simulation future item is %v when it should be 10", fut)
				return
			}
		}
	}
}

func TestDistribution(t *testing.T) {
	sim := NewWithData([]float64{0, 1}, 1, 1000000, 21)
	sim.Run()
	res := sim.Data
	distribution := []int{0, 0}
	for _, item := range res {
		for _, fut := range item.Future {
			distribution[int(fut)]++
		}
	}
	if distribution[0]+distribution[1] != 1000000 {
		t.Errorf("Distribution does not have 1 million elements. It has %v", distribution[0]+distribution[1])
	}
	if distribution[0] < 499000 || distribution[0] > 501000 || distribution[1] < 499000 || distribution[1] > 501000 {
		t.Errorf("Distribution is not balanced %v", distribution)
	}
}

func TestAggregateFutureData(t *testing.T) {
	sim := Simulation{
		Future:      10,
		Simulations: 2,
		Data: []SimulationData{
			SimulationData{Future: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			SimulationData{Future: []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		},
	}
	sim.aggregateFutureData()
	expected := []float64{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	if !reflect.DeepEqual(expected, sim.Data[0].SumFuture) {
		t.Errorf("Arrays are different. Expected %v, got %v", expected, sim.Data[0].SumFuture)
	}
	expected = []float64{10, 19, 27, 34, 40, 45, 49, 52, 54, 55}
	if !reflect.DeepEqual(expected, sim.Data[1].SumFuture) {
		t.Errorf("Arrays are different. Expected %v, got %v", expected, sim.Data[1].SumFuture)
	}
}

func TestForecasts(t *testing.T) {
	sim := Simulation{
		Future:         3,
		Simulations:    20,
		ForecastPoints: 6,
		Data: []SimulationData{
			SimulationData{Future: []float64{1, 10, 11}},
			SimulationData{Future: []float64{2, 11, 10}},
			SimulationData{Future: []float64{3, 12, 12}},
			SimulationData{Future: []float64{4, 13, 9}},
			SimulationData{Future: []float64{5, 14, 13}},
			SimulationData{Future: []float64{6, 15, 8}},
			SimulationData{Future: []float64{7, 16, 14}},
			SimulationData{Future: []float64{8, 17, 7}},
			SimulationData{Future: []float64{9, 18, 15}},
			SimulationData{Future: []float64{10, 19, 6}},
			SimulationData{Future: []float64{11, 20, 16}},
			SimulationData{Future: []float64{12, 1, 5}},
			SimulationData{Future: []float64{13, 2, 17}},
			SimulationData{Future: []float64{14, 3, 4}},
			SimulationData{Future: []float64{15, 4, 18}},
			SimulationData{Future: []float64{16, 5, 3}},
			SimulationData{Future: []float64{17, 6, 19}},
			SimulationData{Future: []float64{18, 7, 2}},
			SimulationData{Future: []float64{19, 8, 20}},
			SimulationData{Future: []float64{20, 9, 1}},
		},
	}
	sim.aggregateFutureData()
	sim.calculateForecasts()
	expected := []Forecast{
		Forecast{Percentil: 100, Forecast: []float64{1, 11, 18}},
		Forecast{Percentil: 80, Forecast: []float64{4, 15, 23}},
		Forecast{Percentil: 60, Forecast: []float64{8, 19, 27}},
		Forecast{Percentil: 40, Forecast: []float64{12, 23, 32}},
		Forecast{Percentil: 20, Forecast: []float64{16, 27, 37}},
		Forecast{Percentil: 0, Forecast: []float64{20, 31, 47}},
	}
	if !reflect.DeepEqual(expected, sim.Forecasts) {
		t.Errorf("Forecasts are different. Expected %v, got %v", expected, sim.Forecasts)
	}
}
