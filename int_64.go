package int256

func (z *Int) Int64() int64 {
	absUint64 := z.abs.Uint64()
	if z.neg {
		return -int64(absUint64)
	}
	return int64(absUint64)
}
