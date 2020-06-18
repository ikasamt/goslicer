package typeslicer

import (
	"strconv"
	"strings"
)

// String
type StringSlice []string

func (ss StringSlice) Where(f func(string, string)bool, substr string) (res StringSlice){
	for _, s := range ss{
		if  f(s, substr) {
			res = append(res, s)
		}
	}
	return
}

func (ss StringSlice) Contains(substr string) (res StringSlice){
	return ss.Where(strings.Contains, substr)
}

func (ss StringSlice) Find(s string) (found bool){
	return ss.Any(func(a string, b string)bool{return a == b }, s)
}

func (ss StringSlice) Any(f func(string,string)bool, s string) (found bool){
	for _, x := range ss {
		if f(x, s) {
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

func (ss StringSlice) HasSuffix(substr string) (res StringSlice){
	return ss.Where(strings.HasSuffix, substr)
}

func (ss StringSlice) HasPrefix(substr string) (res StringSlice){
	return ss.Where(strings.HasPrefix, substr)
}

func (ss StringSlice) StartsWith(substr string) (res StringSlice){
	return ss.HasPrefix(substr)
}

func (ss StringSlice) EndsWith(substr string) (res StringSlice){
	return ss.HasSuffix(substr)
}
