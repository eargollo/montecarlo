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
	sim := NewWithData([]float64{10}, 12, 100)
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
	sim := NewWithData([]float64{0, 1}, 1, 1000000)
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
