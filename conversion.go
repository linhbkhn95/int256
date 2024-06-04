package int256

import (
	"database/sql/driver"
	"errors"
	"strings"
)

// SetFromDecimal sets z from the given string, interpreted as a decimal number.
// OBS! This method is _not_ strictly identical to the (*big.Int).SetString(..., 10) method.
// Notable differences:
// - This method does not accept underscore input, e.g. "100_000"
func (z *Int) SetFromDecimal(decimal string) (err error) {
	z.initiateAbs()
	if strings.HasPrefix(decimal, "-") {
		z.neg = true
		decimal = decimal[1:]
		// take into account the double sign at the beginning which create an error in big.Int
		if strings.HasPrefix(decimal, "+") || strings.HasPrefix(decimal, "-") {
			return ErrMultipleSignAtStart
		}
	}
	return z.abs.SetFromDecimal(decimal)
}

// FromDecimal is a convenience-constructor to create an Int from a
// decimal (base 10) string. Numbers larger than 256 bits are not accepted.
func FromDecimal(decimal string) (*Int, error) {
	var z Int
	if err := z.SetFromDecimal(decimal); err != nil {
		return nil, err
	}
	return &z, nil
}

// MustFromDecimal is a convenience-constructor to create an Int from a
// decimal (base 10) string.
// Returns a new Int and panics if any error occurred.
func MustFromDecimal(decimal string) *Int {
	var z Int
	if err := z.SetFromDecimal(decimal); err != nil {
		panic(err)
	}
	return &z
}

// SetFromHex sets z from the given string, interpreted as a hexadecimal number.
// OBS! This method is _not_ strictly identical to the (*big.Int).SetString(..., 16) method.
// Notable differences:
// - This method _require_ "0x", "0X", "-0x" or "-0X" prefix.
// - This method does not accept zero-prefixed hex, e.g. "0x0001"
// - This method does not accept underscore input, e.g. "100_000",
// - negative value should be prefixed with "-" like "-0x0"
func (z *Int) SetFromHex(hex string) error {
	z.initiateAbs()
	if strings.HasPrefix(hex, "-") {
		z.neg = true
		hex = hex[1:]
		// take into account the double sign at the beginning which create an error in big.Int
		if strings.HasPrefix(hex, "+") || strings.HasPrefix(hex, "-") {
			return ErrMultipleSignAtStart
		}
	}
	return z.abs.SetFromHex(hex)
}

// FromHex is a convenience-constructor to create an Int from
// a hexadecimal string. The string is required to be '0x' or "-0x"-prefixed
// Numbers larger than 256 bits are not accepted.
func FromHex(hex string) (*Int, error) {
	var z Int
	if err := z.SetFromHex(hex); err != nil {
		return nil, err
	}
	return &z, nil
}

// MustFromHex is a convenience-constructor to create an Int from
// a hexadecimal string.
// Returns a new Int and panics if any error occurred.
func MustFromHex(hex string) *Int {
	var z Int
	if err := z.SetFromHex(hex); err != nil {
		panic(err)
	}
	return &z
}

// Hex encodes z in 0x-prefixed or -0x-prefixed  hexadecimal form.
func (z *Int) Hex() string {
	if z.abs == nil {
		return "<nil>"
	}
	if z.abs.IsZero() {
		return "0x0"
	}
	if z.neg {
		return "-" + z.abs.Hex()
	}
	return z.abs.Hex()
}

// Dec returns the decimal representation of z.
func (z *Int) Dec() string {
	z.initiateAbs()
	if z.abs.IsZero() {
		return "0"
	}
	if z.neg {
		return "-" + z.abs.Dec()
	}
	return z.abs.Dec()
}

// MarshalText implements encoding.TextMarshaler
// MarshalText marshals using the decimal representation
func (z *Int) MarshalText() ([]byte, error) {
	return []byte(z.Dec()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler. This method
// can unmarshal either hexadecimal or decimal.
// - For hexadecimal, the input _must_ be prefixed with -0x, -0X, 0x or 0X
func (z *Int) UnmarshalText(input []byte) error {
	if len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X') {
		return z.SetFromHex(string(input))
	}
	if len(input) >= 3 && input[0] == '-' && input[1] == '0' && (input[2] == 'x' || input[2] == 'X') {
		return z.SetFromHex(string(input))
	}
	return z.SetFromDecimal(string(input))
}

// MarshalJSON implements json.Marshaler.
// MarshalJSON marshals using the 'decimal string' representation. This is _not_ compatible
// with big.Int: big.Int marshals into JSON 'native' numeric format.
//
// The JSON  native format is, on some platforms, (e.g. javascript), limited to 53-bit large
// integer space. Thus, U256 uses string-format, which is not compatible with
// big.int (big.Int refuses to unmarshal a string representation).
func (z *Int) MarshalJSON() ([]byte, error) {
	return []byte(`"` + z.Dec() + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler. UnmarshalJSON accepts either
// - Quoted string: either hexadecimal OR decimal
// - For hexadecimal, the input _must_ be prefixed with -0x, -0X, 0x or 0X
// - Not quoted string: only decimal
func (z *Int) UnmarshalJSON(input []byte) error {
	if len(input) < 2 || input[0] != '"' || input[len(input)-1] != '"' {
		// if not quoted, it must be decimal
		return z.SetFromDecimal(string(input))
	}

	return z.UnmarshalText(input[1 : len(input)-1])
}

// String returns the decimal encoding of z
func (z *Int) String() string {
	return z.Dec()
}

// Value implements the database/sql/driver Valuer interface.
// It encodes a base 10 string.
// In Postgres, this will work with both integer and the Numeric/Decimal types
// In MariaDB/MySQL, this will work with the Numeric/Decimal types up to 65 digits, however any more and you should use either VarChar or Char(79)
// In SqLite, use TEXT
func (z *Int) Value() (driver.Value, error) {
	return z.Dec(), nil
}

var (
	ErrMultipleSignAtStart = errors.New("multiple sign at the beginning")
)
