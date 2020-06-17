package typeslicer

import (
	"testing"
)

func TestContains(t *testing.T) {
	tmp := StringSlice{`a`, `d`,`c`, `b`, `1`}

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