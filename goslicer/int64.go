package goslicer

import (
	"math"
	"sort"
	"strconv"
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

func (x Int64Slice) Sum() (res int64) {
	res = 0
	for _, i64 := range x{
		res += i64
	}
	return res
}

func (x Int64Slice) ToString() (res StringSlice) {
	res = StringSlice{}
	for _, i64 := range x {
		res = append(res, strconv.FormatInt(i64, 10))
	}
	return res
}

func (x Int64Slice) Desc() Int64Slice {
	sort.Slice(x, func(i int, j int) bool{ return x[i]>x[j] })
	return x
}

func (x Int64Slice) Asc() Int64Slice{
	sort.Slice(x, func(i int, j int) bool{ return x[i] < x[j] })
	return x
}

