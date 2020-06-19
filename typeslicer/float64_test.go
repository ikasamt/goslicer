package typeslicer

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFloat64(t *testing.T) {
	tmp := Float64Slice{1.0, 2.0, 3.0, 4.1, 5.5, 6.11, 3.14, -199.001}

	// Max
	actual := tmp.Max()
	expected := float64(6.11)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	// Min
	actual = tmp.Min()
	expected = float64(-199.001)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	// Asc
	actualSlice := tmp.Asc()
	expectedSlice := Float64Slice{-199.001, 1.0, 2.0, 3.0, 3.14, 4.1, 5.5, 6.11}
	if diff := cmp.Diff(actualSlice, expectedSlice); diff != "" {
		fmt.Printf("got != want\n%s\n", diff)
	}

	// Desc
	actualSlice = tmp.Desc()
	expectedSlice = Float64Slice{6.11, 5.5, 4.1, 3.14, 3.0, 2.0, 1.0, -199.001}
	if diff := cmp.Diff(actualSlice, expectedSlice); diff != "" {
		fmt.Printf("got != want\n%s\n", diff)
	}

	// Sum
	actualSum := tmp.Sum()
	expectedSum := float64(-174.151)
	if actualSum != expectedSum {
		t.Errorf("got %v\nwant %v", actualSum, expectedSum)
	}

}
