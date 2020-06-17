package goslicer

import "strconv"

// String
type StringSlice []string

func (ss StringSlice) Contains(str string) bool{
	for _, s := range ss{
		if s == str{
			return true
		}
	}
	return false
}

func (ss StringSlice) Mapper(f func(string) string) (res StringSlice) {
	for _, s := range ss {
		res = append(res, f(s))
	}
	return
}

func (ss StringSlice) ToInt() (res IntSlice) {
	res = IntSlice{}
	for _, s := range ss {
		i,_ := strconv.Atoi(s)
		res = append(res, i)
	}
	return res
}
