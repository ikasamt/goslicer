package clefs

import (
	"fmt"
	goslicer2 "goslicer/goslicer"
	"reflect"
	"sort"
)




type Anything struct {
	Name string
}
type Anythings []Anything


//Select[AREA],
//Select[int],
//DistinctBy,

// First is ..
func (anythings Anythings) First(f func(anything Anything) bool) (Anything, error){
	for _, a := range anythings {
		if f(a){
			return a, nil
		}
	}
	return Anything{}, fmt.Errorf(`Not found`)
}

// Where is ..
func (anythings Anythings) Where(f func(anything Anything) bool) (result Anythings){
	for _, a := range anythings {
		if f(a){
			result = append(result, a)
		}
	}
	return
}

// Count is ..
func (anythings Anythings) Count() (counter int) {
	return len(anythings)
}

// CountIf is ..
func (anythings Anythings) CountIf(f func(anything Anything) bool) (counter int) {
	for _, a := range anythings {
		if f(a){
			counter++
		}
	}
	return
}

// Select is ..
func (anythings Anythings) Select(fieldName string) (result goslicer2.InterfaceSlice) {
	for _, a := range anythings {
		i := reflect.ValueOf(a).FieldByName(fieldName).Interface()
		result = append(result, i)
	}
	return
}

// SortBy is ..
func (anythings Anythings) SortBy(sortFunc func(int, int) bool) (result Anythings) {
	sort.Slice(anythings, sortFunc)
	return anythings
}

// DistinctBy is ..
func (anythings Anythings) DistinctBy(f func(anything Anything) interface{}) (result Anythings) {
	tmp := map[interface{}]Anything{}
	for _, a := range anythings {
		tmp[f(a)] = a
	}
	for _, t := range tmp {
		result = append(result, t)
	}
	return
}
