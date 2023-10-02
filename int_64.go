package int256

func (z *Int) Int64() int64 {
	absUint64 := z.abs.Uint64()
	if z.neg {
		return -int64(absUint64)
	}
	return int64(absUint64)
}

func (z *Int) String() string {
	z = z.setZero()

	s := z.abs.ToBig().String()
	if !z.neg {
		return s
	}
	return "-" + s
}
