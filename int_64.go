package int256

func (z *Int) Int64() int64 {
	absUint64 := z.Abs.Uint64()
	if z.Neg {
		return -int64(absUint64)
	}
	return int64(absUint64)
}

func (z *Int) String() string {
	z.initiateAbs()

	s := z.Abs.ToBig().String()
	if !z.Neg {
		return s
	}
	return "-" + s
}
