package int256

import (
	"math/big"

	"github.com/holiman/uint256"
)

var one = uint256.NewInt(1)
var maxUint256 = uint256.MustFromHex("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")

type Int struct {
	Abs *uint256.Int
	Neg bool
}

// Sign returns:
//
//	-1 if x <  0
//	 0 if x == 0
//	+1 if x >  0
func (z *Int) Sign() int {
	if len(z.Abs) == 0 {
		return 0
	}
	if z.Neg {
		return -1
	}
	return 1
}

func New() *Int {
	return &Int{
		Abs: new(uint256.Int),
	}
}

// SetInt64 sets z to x and returns z.
func (z *Int) SetInt64(x int64) *Int {
	neg := false
	if x < 0 {
		neg = true
		x = -x
	}
	if z.Abs == nil {
		panic("abs is nil")
	}
	z.Abs = z.Abs.SetUint64(uint64(x))
	z.Neg = neg
	return z
}

// SetUint64 sets z to x and returns z.
func (z *Int) SetUint64(x uint64) *Int {
	if z.Abs == nil {
		panic("abs is nil")
	}
	z.Abs = z.Abs.SetUint64(x)
	z.Neg = false
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

	neg := x.Neg

	if x.Neg == y.Neg {
		// x + y == x + y
		// (-x) + (-y) == -(x + y)
		z.Abs = z.Abs.Add(x.Abs, y.Abs)
	} else {
		// x + (-y) == x - y == -(y - x)
		// (-x) + y == y - x == -(x - y)
		if x.Abs.Cmp(y.Abs) >= 0 {
			z.Abs = z.Abs.Sub(x.Abs, y.Abs)
		} else {
			neg = !neg
			z.Abs = z.Abs.Sub(y.Abs, x.Abs)
		}
	}
	z.Neg = neg // 0 has no sign
	return z
}

// Sub sets z to the difference x-y and returns z.
func (z *Int) Sub(x, y *Int) *Int {
	z.initiateAbs()

	neg := x.Neg
	if x.Neg != y.Neg {
		// x - (-y) == x + y
		// (-x) - y == -(x + y)
		z.Abs = z.Abs.Add(x.Abs, y.Abs)
	} else {
		// x - y == x - y == -(y - x)
		// (-x) - (-y) == y - x == -(x - y)
		if x.Abs.Cmp(y.Abs) >= 0 {
			z.Abs = z.Abs.Sub(x.Abs, y.Abs)
		} else {
			neg = !neg
			z.Abs = z.Abs.Sub(y.Abs, x.Abs)
		}
	}
	z.Neg = neg // 0 has no sign
	return z
}

// Mul sets z to the product x*y and returns z.
func (z *Int) Mul(x, y *Int) *Int {
	z.initiateAbs()

	z.Abs = z.Abs.Mul(x.Abs, y.Abs)
	z.Neg = x.Neg != y.Neg // 0 has no sign
	return z
}

// Sqrt sets z to ⌊√x⌋, the largest integer such that z² ≤ x, and returns z.
// It panics if x is negative.
func (z *Int) Sqrt(x *Int) *Int {
	z.initiateAbs()

	if x.Neg {
		panic("square root of negative number")
	}
	z.Neg = false
	z.Abs = z.Abs.Sqrt(x.Abs)
	return z
}

// Rsh sets z = x >> n and returns z.
func (z *Int) Rsh(x *Int, n uint) *Int {
	z.initiateAbs()

	if !x.Neg {
		z.Abs.Rsh(x.Abs, n)
		z.Neg = x.Neg
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

	z.Abs = z.Abs.Div(x.Abs, y.Abs)
	z.Neg = len(z.Abs) > 0 && x.Neg != y.Neg // 0 has no sign
	return z
}

// Rem sets z to the remainder x%y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Rem implements truncated modulus (like Go); see QuoRem for more details.
func (z *Int) Rem(x, y *Int) *Int {
	z.initiateAbs()

	z.Abs.Mod(x.Abs, y.Abs)
	z.Neg = len(z.Abs) > 0 && x.Neg // 0 has no sign
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
	case z.Neg == x.Neg:
		r = z.Abs.Cmp(x.Abs)
		if z.Neg {
			r = -r
		}
	case z.Abs.IsZero() && x.Abs.IsZero():
		r = 0
	case z.Neg:
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
	if !x.Neg && !y.Neg && m == nil {
		z.Neg = false
		z.Abs.Exp(x.Abs, y.Abs)
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

	z.Abs.Div(x.Abs, y.Abs)
	if x.Neg == y.Neg {
		z.Neg = false
	} else {
		z.Neg = true
	}
	return z
}

// Lsh sets z = x << n and returns z.
func (z *Int) Lsh(x *Int, n uint) *Int {
	z.initiateAbs()
	b := new(big.Int).Lsh(x.Abs.ToBig(), n)
	z.Abs = uint256.MustFromBig(b)
	z.Neg = x.Neg
	return z
}

// Or sets z = x | y and returns z.
func (z *Int) Or(x, y *Int) *Int {
	z.initiateAbs()

	if x.Neg == y.Neg {
		if x.Neg {
			// (-x) | (-y) == ^(x-1) | ^(y-1) == ^((x-1) & (y-1)) == -(((x-1) & (y-1)) + 1)
			x1 := new(uint256.Int).Sub(x.Abs, one)
			y1 := new(uint256.Int).Sub(y.Abs, one)
			z.Abs = z.Abs.Add(z.Abs.And(x1, y1), one)
			z.Neg = true // z cannot be zero if x and y are negative
			return z
		}

		// x | y == x | y
		z.Abs = z.Abs.Or(x.Abs, y.Abs)
		z.Neg = false
		return z
	}

	// x.neg != y.neg
	if x.Neg {
		x, y = y, x // | is symmetric
	}

	// x | (-y) == x | ^(y-1) == ^((y-1) &^ x) == -(^((y-1) &^ x) + 1)
	y1 := new(uint256.Int).Sub(y.Abs, one)
	z.Abs = z.Abs.Add(z.Abs.And(y1, new(uint256.Int).Xor(x.Abs, maxUint256)), one)
	z.Neg = true // z cannot be zero if one of x or y is negative

	return z
}

// And sets z = x & y and returns z.
func (z *Int) And(x, y *Int) *Int {
	z.initiateAbs()

	if x.Neg == y.Neg {
		if x.Neg {
			// (-x) & (-y) == ^(x-1) & ^(y-1) == ^((x-1) | (y-1)) == -(((x-1) | (y-1)) + 1)
			x1 := new(uint256.Int).Sub(x.Abs, one)
			y1 := new(uint256.Int).Sub(y.Abs, one)
			z.Abs = z.Abs.Add(z.Abs.Or(x1, y1), one)
			z.Neg = true // z cannot be zero if x and y are negative
			return z
		}

		// x & y == x & y
		z.Abs = z.Abs.And(x.Abs, y.Abs)
		z.Neg = false
		return z
	}

	// x.neg != y.neg
	if x.Neg {
		x, y = y, x // & is symmetric
	}

	// x & (-y) == x & ^(y-1) == x &^ (y-1)
	y1 := new(uint256.Int).Sub(y.Abs, one)
	z.Abs = z.Abs.And(x.Abs, new(uint256.Int).Xor(y1, maxUint256))
	z.Neg = false

	return z
}

// initiateAbs sets default value for `z.abs` value if is nil
func (z *Int) initiateAbs() {
	if z.Abs == nil {
		z.Abs = new(uint256.Int)
	}

}
