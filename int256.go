package int256

import (
	"math/big"

	"github.com/holiman/uint256"
)

var one = uint256.NewInt(1)
var maxUint256 = uint256.MustFromHex("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")

type Int struct {
	abs *uint256.Int
	neg bool
}

// Sign returns:
//
//	-1 if x <  0
//	 0 if x == 0
//	+1 if x >  0
func (z *Int) Sign() int {
	if len(z.abs) == 0 {
		return 0
	}
	if z.neg {
		return -1
	}
	return 1
}

func New() *Int {
	return &Int{
		abs: new(uint256.Int),
	}
}

// SetInt64 sets z to x and returns z.
func (z *Int) SetInt64(x int64) *Int {
	neg := false
	if x < 0 {
		neg = true
		x = -x
	}
	if z.abs == nil {
		panic("abs is nil")
	}
	z.abs = z.abs.SetUint64(uint64(x))
	z.neg = neg
	return z
}

// SetUint64 sets z to x and returns z.
func (z *Int) SetUint64(x uint64) *Int {
	if z.abs == nil {
		panic("abs is nil")
	}
	z.abs = z.abs.SetUint64(x)
	z.neg = false
	return z
}

// NewInt allocates and returns a new Int set to x.
func NewInt(x int64) *Int {
	return New().SetInt64(x)
}

// SetUint64 sets z to x and returns z.
func (z *Int) SetString(s string) (*Int, error) {
	origin := s
	neg := false
	// Remove max one leading +
	if len(s) > 0 && s[0] == '+' {
		neg = false
		s = s[1:]
	}

	if len(s) > 0 && s[0] == '-' {
		neg = true
		s = s[1:]
	}
	var (
		abs *uint256.Int
		err error
	)
	abs, err = uint256.FromDecimal(s)
	if err != nil {
		// TODO: parse base as input param
		b, ok := new(big.Int).SetString(origin, 16)
		if !ok {
			return nil, err
		}
		return MustFromBig(b), nil
	}

	return &Int{
		abs,
		neg,
	}, nil
}

// // setFromScanner implements SetString given an io.ByteScanner.
// // For documentation see comments of SetString.
// func (z *Int) setFromScanner(r io.ByteScanner, base int) (*Int, bool) {
// 	if _, _, err := z.scan(r, base); err != nil {
// 		return nil, false
// 	}
// 	// entire content must have been consumed
// 	if _, err := r.ReadByte(); err != io.EOF {
// 		return nil, false
// 	}
// 	return z, true // err == io.EOF => scan consumed all content of r
// }

func (z *Int) Add(x, y *Int) *Int {
	z.initiateAbs()

	neg := x.neg

	if x.neg == y.neg {
		// x + y == x + y
		// (-x) + (-y) == -(x + y)
		z.abs = z.abs.Add(x.abs, y.abs)
	} else {
		// x + (-y) == x - y == -(y - x)
		// (-x) + y == y - x == -(x - y)
		if x.abs.Cmp(y.abs) >= 0 {
			z.abs = z.abs.Sub(x.abs, y.abs)
		} else {
			neg = !neg
			z.abs = z.abs.Sub(y.abs, x.abs)
		}
	}
	z.neg = neg // 0 has no sign
	return z
}

// Sub sets z to the difference x-y and returns z.
func (z *Int) Sub(x, y *Int) *Int {
	z.initiateAbs()

	neg := x.neg
	if x.neg != y.neg {
		// x - (-y) == x + y
		// (-x) - y == -(x + y)
		z.abs = z.abs.Add(x.abs, y.abs)
	} else {
		// x - y == x - y == -(y - x)
		// (-x) - (-y) == y - x == -(x - y)
		if x.abs.Cmp(y.abs) >= 0 {
			z.abs = z.abs.Sub(x.abs, y.abs)
		} else {
			neg = !neg
			z.abs = z.abs.Sub(y.abs, x.abs)
		}
	}
	z.neg = neg // 0 has no sign
	return z
}

// Mul sets z to the product x*y and returns z.
func (z *Int) Mul(x, y *Int) *Int {
	z.initiateAbs()

	z.abs = z.abs.Mul(x.abs, y.abs)
	z.neg = x.neg != y.neg // 0 has no sign
	return z
}

// Sqrt sets z to ⌊√x⌋, the largest integer such that z² ≤ x, and returns z.
// It panics if x is negative.
func (z *Int) Sqrt(x *Int) *Int {
	z.initiateAbs()

	if x.neg {
		panic("square root of negative number")
	}
	z.neg = false
	z.abs = z.abs.Sqrt(x.abs)
	return z
}

// Rsh sets z = x >> n and returns z.
func (z *Int) Rsh(x *Int, n uint) *Int {
	z.initiateAbs()

	if !x.neg {
		z.abs.Rsh(x.abs, n)
		z.neg = x.neg
		return z
	}
	// TODO: implement
	b := x.ToBig()
	return MustFromBig(b.Rsh(b, n))
}

// Quo sets z to the quotient x/y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Quo implements truncated division (like Go); see QuoRem for more details.
func (z *Int) Quo(x, y *Int) *Int {
	z.initiateAbs()

	z.abs = z.abs.Div(x.abs, y.abs)
	z.neg = len(z.abs) > 0 && x.neg != y.neg // 0 has no sign
	return z
}

// Rem sets z to the remainder x%y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Rem implements truncated modulus (like Go); see QuoRem for more details.
func (z *Int) Rem(x, y *Int) *Int {
	z.initiateAbs()

	z.abs.Mod(x.abs, y.abs)
	z.neg = len(z.abs) > 0 && x.neg // 0 has no sign
	return z
}

// Cmp compares x and y and returns:
//
//	-1 if x <  y
//	 0 if x == y
//	+1 if x >  y
func (z *Int) Cmp(x *Int) (r int) {
	z.initiateAbs()

	// x cmp y == x cmp y
	// x cmp (-y) == x
	// (-x) cmp y == y
	// (-x) cmp (-y) == -(x cmp y)
	switch {
	case z == x:
		// nothing to do
	case z.neg == x.neg:
		r = z.abs.Cmp(x.abs)
		if z.neg {
			r = -r
		}
	case z.abs.IsZero() && x.abs.IsZero():
		r = 0
	case z.neg:
		r = -1
	default:
		r = 1
	}
	return
}

// Exp sets z = x**y mod |m| (i.e. the sign of m is ignored), and returns z.
// If m == nil or m == 0, z = x**y unless y <= 0 then z = 1. If m != 0, y < 0,
// and x and m are not relatively prime, z is unchanged and nil is returned.
//
// Modular exponentiation of inputs of a particular size is not a
// cryptographically constant-time operation.
func (z *Int) Exp(x, y, m *Int) *Int {
	z.initiateAbs()

	if x == nil {
		panic("x is nil")
	}
	if !x.neg && !y.neg && m == nil {
		z.neg = false
		z.abs.Exp(x.abs, y.abs)
		return z
	}
	// TODO: implement
	var mBigInt *big.Int
	if m != nil {
		mBigInt = m.ToBig()
	}
	big := new(big.Int).Exp(x.ToBig(), y.ToBig(), mBigInt)
	z, _ = FromBig(big)
	return z
}

func (z *Int) Div(x, y *Int) *Int {
	z.initiateAbs()

	z.abs.Div(x.abs, y.abs)
	if x.neg == y.neg {
		z.neg = false
	} else {
		z.neg = true
	}
	return z
}

// Lsh sets z = x << n and returns z.
func (z *Int) Lsh(x *Int, n uint) *Int {
	z.initiateAbs()
	b := new(big.Int).Lsh(x.abs.ToBig(), n)
	z.abs = uint256.MustFromBig(b)
	z.neg = x.neg
	return z
}

// Or sets z = x | y and returns z.
func (z *Int) Or(x, y *Int) *Int {
	z.initiateAbs()

	if x.neg == y.neg {
		if x.neg {
			// (-x) | (-y) == ^(x-1) | ^(y-1) == ^((x-1) & (y-1)) == -(((x-1) & (y-1)) + 1)
			x1 := new(uint256.Int).Sub(x.abs, one)
			y1 := new(uint256.Int).Sub(y.abs, one)
			z.abs = z.abs.Add(z.abs.And(x1, y1), one)
			z.neg = true // z cannot be zero if x and y are negative
			return z
		}

		// x | y == x | y
		z.abs = z.abs.Or(x.abs, y.abs)
		z.neg = false
		return z
	}

	// x.neg != y.neg
	if x.neg {
		x, y = y, x // | is symmetric
	}

	// x | (-y) == x | ^(y-1) == ^((y-1) &^ x) == -(^((y-1) &^ x) + 1)
	y1 := new(uint256.Int).Sub(y.abs, one)
	z.abs = z.abs.Add(z.abs.And(y1, new(uint256.Int).Xor(x.abs, maxUint256)), one)
	z.neg = true // z cannot be zero if one of x or y is negative

	return z
}

// And sets z = x & y and returns z.
func (z *Int) And(x, y *Int) *Int {
	z.initiateAbs()

	if x.neg == y.neg {
		if x.neg {
			// (-x) & (-y) == ^(x-1) & ^(y-1) == ^((x-1) | (y-1)) == -(((x-1) | (y-1)) + 1)
			x1 := new(uint256.Int).Sub(x.abs, one)
			y1 := new(uint256.Int).Sub(y.abs, one)
			z.abs = z.abs.Add(z.abs.Or(x1, y1), one)
			z.neg = true // z cannot be zero if x and y are negative
			return z
		}

		// x & y == x & y
		z.abs = z.abs.And(x.abs, y.abs)
		z.neg = false
		return z
	}

	// x.neg != y.neg
	if x.neg {
		x, y = y, x // & is symmetric
	}

	// x & (-y) == x & ^(y-1) == x &^ (y-1)
	y1 := new(uint256.Int).Sub(y.abs, one)
	z.abs = z.abs.And(x.abs, new(uint256.Int).Xor(y1, maxUint256))
	z.neg = false

	return z
}

// initiateAbs sets default value for `z.abs` value if is nil
func (z *Int) initiateAbs() {
	if z.abs == nil {
		z.abs = new(uint256.Int)
	}

}
