package goslicer

// Interface
type InterfaceSlice []interface{}

func (x InterfaceSlice) ToInt() (res []int) {
	res = []int{}
	for _, v := range x {
		res = append(res, v.(int))
	}
	return
}

func (x InterfaceSlice) ToString() (res []string) {
	res = []string{}
	for _, v := range x {
		res = append(res, v.(string))
	}
	return
}

