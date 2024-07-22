package int256

import (
	"math/big"

	"github.com/holiman/uint256"
)

// UnmarshalJSON implements json.Unmarshaler.
func (z *Int) UnmarshalJSON(input []byte) error {
	b := new(big.Int)
	if err := b.UnmarshalJSON(input); err != nil {
		return err
	}
	if b.Cmp(zero) < 0 {
		b.Neg(b)
		z.neg = true
	} else {
		z.neg = false
	}
	z.abs = new(uint256.Int)
	z.abs.SetFromBig(b)
	return nil
}

// MarshalJSON implements json.Marshaler.
func (z *Int) MarshalJSON() ([]byte, error) {
	return z.ToBig().MarshalJSON()
}
