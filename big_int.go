package int256

import (
	"math/big"

	"github.com/holiman/uint256"
)

var negativeOneBigInt = big.NewInt(-1)
var zero = big.NewInt(0)

func (z *Int) ToBig() *big.Int {
	b := z.abs.ToBig()
	if z.neg {
		return b.Mul(b, negativeOneBigInt)
	}
	return b
}

func MustFromBig(x *big.Int) *Int {
	big, ok := FromBig(x)
	if !ok {
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
	abs, ok := uint256.FromBig(num)
	if !ok {
		return nil, false
	}
	return &Int{
		abs,
		neg,
	}, true
}