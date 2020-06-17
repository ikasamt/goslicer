package typeslicer

// Interface
type InterfaceSlice []interface{}

func (x InterfaceSlice) AsInt() (res IntSlice) {
	res = IntSlice{}
	for _, v := range x {
		res = append(res, v.(int))
	}
	return
}

func (x InterfaceSlice) AsString() (res StringSlice) {
	res = StringSlice{}
	for _, v := range x {
		res = append(res, v.(string))
	}
	return
}

func (x InterfaceSlice) AsInt64() (res Int64Slice) {
	res = Int64Slice{}
	for _, v := range x {
		res = append(res, v.(int64))
	}
	return
}


