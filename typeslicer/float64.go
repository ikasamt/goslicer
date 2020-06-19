package typeslicer

import (
	"math"
	"sort"
	"strconv"
)

// Int
type Float64Slice []float64

func (x Float64Slice) Any(f func(float64, float64) bool, n float64) (found bool) {
	for _, i := range x {
		if f(i, n) {
			return true
		}
	}
	return false
}

func (x Float64Slice) Where(f func(float64) bool) (res Float64Slice) {
	for _, i := range x {
		if f(i) {
			res = append(res, i)
		}
	}
	return
}

func (x Float64Slice) Inject(defaultValue float64, f func(float64, float64) float64) (res float64) {
	res = defaultValue
	for _, i := range x {
		res = f(res, i)
	}
	return
}

func (x Float64Slice) Collect(f func(float64) float64) (res Float64Slice) {
	for _, i := range x {
		res = append(res, f(i))
	}
	return
}

func (x Float64Slice) CollectInterface(f func(float64) interface{}) (res InterfaceSlice) {
	for _, i := range x {
		res = append(res, f(i))
	}
	return
}

func (x Float64Slice) Find(n float64) (found bool) {
	return x.Any(func(a float64, b float64) bool { return a == b }, n)
}

func (x Float64Slice) Min() (res float64) {
	return x.Inject(math.MaxInt32, func(a float64, b float64) float64 {
		if a < b {
			return a
		} else {
			return b
		}
	})
}

func (x Float64Slice) Max() (res float64) {
	return x.Inject(math.MinInt32, func(a float64, b float64) float64 {
		if a > b {
			return a
		} else {
			return b
		}
	})
}

func (x Float64Slice) Sum() (res float64) {
	return x.Inject(0, func(a float64, b float64) float64 {
		return a + b
	})
}

func (x Float64Slice) ToString() (res InterfaceSlice) {
	for _, i := range x {
		res = append(res, strconv.FormatFloat(i, 'e', 2, 64))
	}
	return
}

func (x Float64Slice) Desc() Float64Slice {
	sort.Slice(x, func(i int, j int) bool { return x[i] > x[j] })
	return x
}

func (x Float64Slice) Asc() Float64Slice {
	sort.Slice(x, func(i int, j int) bool { return x[i] < x[j] })
	return x
}
