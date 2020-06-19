package typeslicer

import (
	"math"
	"sort"
	"strconv"
)

// Int
type IntSlice []int

func (x IntSlice) Any(f func(int, int) bool, n int) (found bool) {
	for _, i := range x {
		if f(i, n) {
			return true
		}
	}
	return false
}

func (x IntSlice) Where(f func(int) bool) (res IntSlice) {
	for _, i := range x {
		if f(i) {
			res = append(res, i)
		}
	}
	return
}

func (x IntSlice) Inject(defaultValue int, f func(int, int) int) (res int) {
	res = defaultValue
	for _, i := range x {
		res = f(res, i)
	}
	return
}

func (x IntSlice) Collect(f func(int) int) (res IntSlice) {
	for _, i := range x {
		res = append(res, f(i))
	}
	return
}

func (x IntSlice) CollectInterface(f func(int) interface{}) (res InterfaceSlice) {
	for _, i := range x {
		res = append(res, f(i))
	}
	return
}

func (x IntSlice) Find(n int) (found bool) {
	return x.Any(func(a int, b int) bool { return a == b }, n)
}

func (x IntSlice) Min() (res int) {
	return x.Inject(math.MaxInt32, func(a int, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	})
}

func (x IntSlice) Max() (res int) {
	return x.Inject(math.MinInt32, func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	})
}

func (x IntSlice) Sum() (res int) {
	return x.Inject(0, func(a int, b int) int {
		return a + b
	})
}

func (x IntSlice) ToString() (res InterfaceSlice) {
	for _, i := range x {
		res = append(res, strconv.FormatInt(int64(i), 10))
	}
	return
}

func (x IntSlice) Desc() IntSlice {
	sort.Slice(x, func(i int, j int) bool { return x[i] > x[j] })
	return x
}

func (x IntSlice) Asc() IntSlice {
	sort.Slice(x, func(i int, j int) bool { return x[i] < x[j] })
	return x
}
