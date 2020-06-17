package goslicer

import (
	"math"
	"sort"
)

// Int64
type Int64Slice []int64

func (x Int64Slice) MinMax(dummyValue int64, compareF func(int64, int64) bool) (res int64) {
	res = dummyValue
	for _, i64 := range x {
		isOK := compareF(res, i64)
		if isOK {
			res = i64
		}
	}
	return
}

func (x Int64Slice) Min() (res int64) {
	return x.MinMax(math.MaxInt64, func(a int64, b int64)bool{ return a > b })
}

func (x Int64Slice) Max() (res int64) {
	return x.MinMax(math.MinInt64, func(a int64, b int64)bool{ return a < b })
}

func (x Int64Slice) Desc() Int64Slice {
	sort.Slice(x, func(i int, j int) bool{ return x[i]>x[j] })
	return x
}

func (x Int64Slice) Asc() Int64Slice{
	sort.Slice(x, func(i int, j int) bool{ return x[i] < x[j] })
	return x
}

// Int
type IntSlice []int

func (x IntSlice) MinMax(dummyValue int, compareF func(int, int) bool) (res int) {
	res = dummyValue
	for _, i32 := range x {
		isOK := compareF(res, i32)
		if isOK {
			res = i32
		}
	}
	return
}

func (x IntSlice) Min() (res int) {
	return x.MinMax(math.MaxInt32, func(a int, b int)bool{ return a > b })
}

func (x IntSlice) Max() (res int) {
	return x.MinMax(math.MinInt32, func(a int, b int)bool{ return a < b })
}


func (x IntSlice) Desc() IntSlice {
	sort.Slice(x, func(i int, j int) bool{ return x[i]>x[j] })
	return x
}

func (x IntSlice) Asc() IntSlice{
	sort.Slice(x, func(i int, j int) bool{ return x[i] < x[j] })
	return x
}
