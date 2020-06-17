package goslicer_test

import (
	"goslicer/goslicer"
	"testing"
)

func TestContains(t *testing.T) {
	tmp := goslicer.StringSlice{`a`, `d`,`c`, `b`, `1`}

	// 含まれない
	actual :=  tmp.Contains(`x`)
	expected := false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	// 含まれる
	actual =  tmp.Contains(`a`)
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

}