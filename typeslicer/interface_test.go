package typeslicer

import (
	"reflect"
	"testing"
)


func TestToInt(t *testing.T) {
	tmp := InterfaceSlice{1,2,3}
	for _, i := range tmp{
		if reflect.TypeOf(i).Kind() != reflect.Int{
			t.Fatal("is not int")
		}
	}
}

func TestToString(t *testing.T) {
	tmp := InterfaceSlice{`1`,`2`,`3`}
	for _, i := range tmp{
		if reflect.TypeOf(i).Kind() != reflect.String{
			t.Fatal("is not string")
		}
	}
}