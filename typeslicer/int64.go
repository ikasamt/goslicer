package typeslicer

import (
	"math"
	"sort"
	"strconv"
)

// Int64
type Int64Slice []int64

func (x Int64Slice) Any(f func(int64,int64)bool, n int64) (found bool){
	for _, i := range x{
		if f(i, n) {
			return true
		}
	}
	return false
}

func (x Int64Slice) Where(f func(int64)bool) (res Int64Slice){
	for _, i := range x{
		if  f(i) {
			res = append(res, i)
		}
	}
	return
}

func (x Int64Slice) Find(n int64) (found bool){
	return x.Any(func(a int64, b int64)bool{return a == b }, n)
}

func (x Int64Slice) MinMax(defaultValue int64, compareF func(int64, int64) bool) (res int64) {
	res = defaultValue
	for _, i := range x {
		isOK := compareF(res, i)
		if isOK {
			res = i
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

func (x Int64Slice) Mapper(f func(int64) int64) (res Int64Slice) {
	for _, i := range x {
		res = append(res, f(i))
	}
	return
}

func (x Int64Slice) Desc() Int64Slice {
	sort.Slice(x, func(i int, j int) bool{ return x[i]>x[j] })
	return x
}

func (x Int64Slice) Asc() Int64Slice{
	sort.Slice(x, func(i int, j int) bool{ return x[i] < x[j] })
	return x
}

