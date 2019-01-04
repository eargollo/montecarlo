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
	res := sim.Run()
	if len(res) != 100 {
		t.Errorf("Simulation did not return the 100 expected results. It returned only %v", len(res))
	}
	for _, item := range res {
		if len(item.future) != 12 {
			t.Errorf("Simulation item has %v future instead of 12", len(item.future))
			return
		}
		for _, fut := range item.future {
			if fut != 10 {
				t.Errorf("Simulation future item is %v when it should be 10", fut)
				return
			}
		}
	}
}

// func TestWithSingleDataFile(t *testing.T) {
// 	simulation := New("./testdata/singledata.csv")
// }
