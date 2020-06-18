package typeslicer

import (
	"math"
	"sort"
	"strconv"
)

// Int
type IntSlice []int

func (x IntSlice) Any(f func(int,int)bool, n int) (found bool){
	for _, i := range x{
		if f(i, n) {
			return true
		}
	}
	return false
}

func (x IntSlice) Where(f func(int)bool) (res IntSlice){
	for _, i := range x{
		if  f(i) {
			res = append(res, i)
		}
	}
	return
}

func (x IntSlice) Find(n int) (found bool){
	return x.Any(func(a int, b int)bool{return a == b }, n)
}

func (x IntSlice) MinMax(defaultValue int, compareF func(int, int) bool) (res int) {
	res = defaultValue
	for _, i := range x {
		isOK := compareF(res, i)
		if isOK {
			res = i
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

