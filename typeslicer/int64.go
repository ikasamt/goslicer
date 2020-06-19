package typeslicer

import (
	"math"
	"sort"
	"strconv"
)

// Int
type Int64Slice []int64

func (x Int64Slice) Any(f func(int64, int64) bool, n int64) (found bool) {
	for _, i := range x {
		if f(i, n) {
			return true
		}
	}
	return false
}

func (x Int64Slice) Where(f func(int64) bool) (res Int64Slice) {
	for _, i := range x {
		if f(i) {
			res = append(res, i)
		}
	}
	return
}

func (x Int64Slice) Inject(defaultValue int64, f func(int64, int64) int64) (res int64) {
	res = defaultValue
	for _, i := range x {
		res = f(res, i)
	}
	return
}

func (x Int64Slice) Collect(f func(int64) int64) (res Int64Slice) {
	for _, i := range x {
		res = append(res, f(i))
	}
	return
}

func (x Int64Slice) CollectInterface(f func(int64) interface{}) (res InterfaceSlice) {
	for _, i := range x {
		res = append(res, f(i))
	}
	return
}

func (x Int64Slice) Find(n int64) (found bool) {
	return x.Any(func(a int64, b int64) bool { return a == b }, n)
}

func (x Int64Slice) Min() (res int64) {
	return x.Inject(math.MaxInt32, func(a int64, b int64) int64 {
		if a < b {
			return a
		} else {
			return b
		}
	})
}

func (x Int64Slice) Max() (res int64) {
	return x.Inject(math.MinInt32, func(a int64, b int64) int64 {
		if a > b {
			return a
		} else {
			return b
		}
	})
}

func (x Int64Slice) Sum() (res int64) {
	return x.Inject(0, func(a int64, b int64) int64 {
		return a + b
	})
}

func (x Int64Slice) ToString() (res InterfaceSlice) {
	for _, i := range x {
		res = append(res, strconv.FormatInt(int64(i), 10))
	}
	return
}

func (x Int64Slice) Desc() Int64Slice {
	sort.Slice(x, func(i int, j int) bool { return x[i] > x[j] })
	return x
}

func (x Int64Slice) Asc() Int64Slice {
	sort.Slice(x, func(i int, j int) bool { return x[i] < x[j] })
	return x
}
