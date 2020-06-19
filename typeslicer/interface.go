package typeslicer

// Interface
type InterfaceSlice []interface{}

func (x InterfaceSlice) ToInt() (res IntSlice) {
	res = IntSlice{}
	for _, v := range x {
		res = append(res, v.(int))
	}
	return
}

func (x InterfaceSlice) ToString() (res StringSlice) {
	res = StringSlice{}
	for _, v := range x {
		res = append(res, v.(string))
	}
	return
}

func (x InterfaceSlice) ToInt64() (res Int64Slice) {
	res = Int64Slice{}
	for _, v := range x {
		res = append(res, v.(int64))
	}
	return
}
