package typeslicer

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestContains(t *testing.T) {
	tmp := StringSlice{`a`, `d`,`c`, `b`, `1`, `abc`, `利根川`, `最上川`, `荒川`, `荒川区`}

	// 含まれない
	actual :=  tmp.Find(`x`)
	expected := false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	// 含まれる
	actual =  tmp.Find(`a`)
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	// 含まれる
	actualStrings := tmp.Contains(`a`)
	expectedStrings := StringSlice{`a`, `abc`}
	if diff :=cmp.Diff(actualStrings, expectedStrings); diff != "" {
		t.Errorf("got != want\n%s\n", diff)
	}

	// endwith
	actualStrings = tmp.EndsWith(`川`)
	if diff :=cmp.Diff(actualStrings, StringSlice{`利根川`, `最上川`, `荒川`}); diff != "" {
		t.Errorf("got != want\n%s\n", diff)
	}

	// startwith
	actualStrings = tmp.StartsWith(`荒`)
	if diff :=cmp.Diff(actualStrings, StringSlice{`荒川`, `荒川区`}); diff != "" {
		t.Errorf("got != want\n%s\n", diff)
	}

	// startwith
	actualStrings = tmp.Contains(`川`)
	if diff :=cmp.Diff(actualStrings, StringSlice{`利根川`, `最上川`, `荒川`, `荒川区`}); diff != "" {
		t.Errorf("got != want\n%s\n", diff)
	}

}