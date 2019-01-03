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

// func TestWithSingleDataFile(t *testing.T) {
// 	simulation := New("./testdata/singledata.csv")
// }
