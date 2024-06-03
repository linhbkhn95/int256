package int256

import (
	"math/big"

	"github.com/holiman/uint256"
)

var negativeOneBigInt = big.NewInt(-1)
var zero = big.NewInt(0)

func (z *Int) ToBig() *big.Int {
	b := z.Abs.ToBig()
	if z.Neg {
		return b.Mul(b, negativeOneBigInt)
	}
	return b
}

func MustFromBig(x *big.Int) *Int {
	big, overflow := FromBig(x)
	if overflow {
		panic("cannot parsing from big.Int")
	}
	return big
}

func FromBig(x *big.Int) (*Int, bool) {
	num := x
	neg := false
	if x.Cmp(zero) == -1 {
		num = new(big.Int).Mul(x, negativeOneBigInt)
		neg = true
	}
	abs, overflow := uint256.FromBig(num)
	/*
		type Int [4]uint64
		Currently, uint26 has maxWord is 4


		bigInt has len(words) that can great than 4. So we can receive overflow error.

		words := b.Bits()
		overflow := len(words) > maxWords
		ref from uint256 code
		https://github.com/holiman/uint256/blob/master/conversion.go#L202
	*/
	if overflow {
		abs, err := uint256.FromDecimal(x.String())
		if err != nil {
			return nil, overflow
		}
		neg = x.Sign() < 0
		return &Int{
			abs,
			neg,
		}, true
	}
	return &Int{
		abs,
		neg,
	}, false
}
