package typeslicer

import (
	"math"
	"sort"
	"strconv"
)

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

func (x IntSlice) Sum() (res int) {
	res = 0
	for _, i := range x{
		res += i
	}
	return res
}

func (x IntSlice) ToString() (res StringSlice) {
	res = StringSlice{}
	for _, i := range x {
		res = append(res, strconv.Itoa(i))
	}
	return res
}

func (x IntSlice) Mapper(f func(int) int) (res IntSlice) {
	for _, i := range x {
		res = append(res, f(i))
	}
	return
}

func (x IntSlice) Desc() IntSlice {
	sort.Slice(x, func(i int, j int) bool{ return x[i]>x[j] })
	return x
}

func (x IntSlice) Asc() IntSlice{
	sort.Slice(x, func(i int, j int) bool{ return x[i] < x[j] })
	return x
}
