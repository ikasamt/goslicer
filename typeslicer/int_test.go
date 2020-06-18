package typeslicer

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInt64(t *testing.T) {
	tmp := Int64Slice{1,2,3,4,5,6,11,-199}

	// Max
	actual :=  tmp.Max()
	expected := int64(11)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	// Min
	actual =  tmp.Min()
	expected = int64(-199)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	// Asc
	actualSlice := tmp.Asc()
	expectedSlice := Int64Slice{-199, 1,2,3,4,5,6,11}
	if diff :=cmp.Diff(actualSlice, expectedSlice); diff != "" {
		fmt.Printf("got != want\n%s\n", diff)
	}

	// Desc
	actualSlice = tmp.Desc()
	expectedSlice = Int64Slice{11,6,5,4,3,2,1,-199}
	if diff :=cmp.Diff(actualSlice, expectedSlice); diff != "" {
		fmt.Printf("got != want\n%s\n", diff)
	}

	// Sum
	actualSum := tmp.Sum()
	expectedSum := int64(-167)
	if actualSum != expectedSum{
		t.Errorf("got %v\nwant %v", actualSum, expectedSum)
	}

}


func TestInt32(t *testing.T) {
	tmp := IntSlice{1,2,3,4,5,6,11,-199}

	// Max
	actual :=  tmp.Max()
	expected := 11
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	// Min
	actual =  tmp.Min()
	expected = -199
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	// Asc
	actualSlice := tmp.Asc()
	expectedSlice := IntSlice{-199, 1,2,3,4,5,6,11}
	if diff :=cmp.Diff(actualSlice, expectedSlice); diff != "" {
		fmt.Printf("got != want\n%s\n", diff)
	}

	// Desc
	actualSlice = tmp.Desc()
	expectedSlice = IntSlice{11,6,5,4,3,2,1,-199}
	if diff :=cmp.Diff(actualSlice, expectedSlice); diff != "" {
		fmt.Printf("got != want\n%s\n", diff)
	}

	// Sum
	actualSum := tmp.Sum()
	expectedSum := -167
	if actualSum != expectedSum{
		t.Errorf("got %v\nwant %v", actualSum, expectedSum)
	}


}

