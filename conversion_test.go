package int256

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/holiman/uint256"
	"math/big"
	"math/rand"
	"testing"
)

const numSamples = 1024

var (
	int256Samples      [numSamples]Int
	int256SamplesLt    [numSamples]Int // int256SamplesLt[i] <= int256Samples[i]
	big256Samples      [numSamples]big.Int
	big256SamplesLt    [numSamples]big.Int // big256SamplesLt[i] <= big256Samples[i]
	int256SamplesNeg   [numSamples]Int
	int256SamplesLtNeg [numSamples]Int // int256SamplesLtNrg[i] <= int256SamplesNeg[i]
	big256SamplesNeg   [numSamples]big.Int
	big256SamplesLtNeg [numSamples]big.Int // big256SamplesLtNeg[i] <= big256SamplesNeg[i]
	_                  = initSamples()

	encodeBigTests = []marshalTest{
		{referenceBig("0"), "0x0"},
		{referenceBig("1"), "0x1"},
		{referenceBig("ff"), "0xff"},
		{referenceBig("112233445566778899aabbccddeeff"), "0x112233445566778899aabbccddeeff"},
		{referenceBig("80a7f2c1bcc396c00"), "0x80a7f2c1bcc396c00"},
		{referenceBig("-0"), "0x0"},
		{referenceBig("-1"), "-0x1"},
		{referenceBig("-ff"), "-0xff"},
		{referenceBig("-112233445566778899aabbccddeeff"), "-0x112233445566778899aabbccddeeff"},
		{referenceBig("-80a7f2c1bcc396c00"), "-0x80a7f2c1bcc396c00"},
	}

	decodeBigTests = []unmarshalTest{
		// invalid
		{input: ``, wantErr: uint256.ErrEmptyString},
		{input: `0`, wantErr: uint256.ErrMissingPrefix},
		{input: `0x`, wantErr: uint256.ErrEmptyNumber},
		{input: `0x01`, wantErr: uint256.ErrLeadingZero},
		{input: `0xx`, wantErr: uint256.ErrSyntax},
		{input: `0x1zz01`, wantErr: uint256.ErrSyntax},
		{
			input:   `0x10000000000000000000000000000000000000000000000000000000000000000`,
			wantErr: uint256.ErrBig256Range,
		},
		{input: `-`, wantErr: uint256.ErrEmptyString},
		{input: `-0`, wantErr: uint256.ErrMissingPrefix},
		{input: `-0x`, wantErr: uint256.ErrEmptyNumber},
		{input: `-0x01`, wantErr: uint256.ErrLeadingZero},
		{input: `-0xx`, wantErr: uint256.ErrSyntax},
		{input: `-0x1zz01`, wantErr: uint256.ErrSyntax},
		{
			input:   `-0x10000000000000000000000000000000000000000000000000000000000000000`,
			wantErr: uint256.ErrBig256Range,
		},
		// valid
		{input: `0x0`, want: big.NewInt(0)},
		{input: `0x2`, want: big.NewInt(0x2)},
		{input: `0x2F2`, want: big.NewInt(0x2f2)},
		{input: `0X2F2`, want: big.NewInt(0x2f2)},
		{input: `0x1122aaff`, want: big.NewInt(0x1122aaff)},
		{input: `0xbBb`, want: big.NewInt(0xbbb)},
		{input: `0xfffffffff`, want: big.NewInt(0xfffffffff)},
		{
			input: `0x112233445566778899aabbccddeeff`,
			want:  referenceBig("112233445566778899aabbccddeeff"),
		},
		{
			input: `0xffffffffffffffffffffffffffffffffffff`,
			want:  referenceBig("ffffffffffffffffffffffffffffffffffff"),
		},
		{
			input: `0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff`,
			want:  referenceBig("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"),
		},
		{input: `-0x0`, want: big.NewInt(0)},
		{input: `-0x2`, want: big.NewInt(-0x2)},
		{input: `-0x2F2`, want: big.NewInt(-0x2f2)},
		{input: `-0X2F2`, want: big.NewInt(-0x2f2)},
		{input: `-0x1122aaff`, want: big.NewInt(-0x1122aaff)},
		{input: `-0xbBb`, want: big.NewInt(-0xbbb)},
		{input: `-0xfffffffff`, want: big.NewInt(-0xfffffffff)},
		{
			input: `-0x112233445566778899aabbccddeeff`,
			want:  referenceBig("-112233445566778899aabbccddeeff"),
		},
		{
			input: `-0xffffffffffffffffffffffffffffffffffff`,
			want:  referenceBig("-ffffffffffffffffffffffffffffffffffff"),
		},
		{
			input: `-0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff`,
			want:  referenceBig("-ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"),
		},
	}
)

func initSamples() bool {
	rnd := rand.New(rand.NewSource(0))

	// newRandInt creates new Int with so many highly likely non-zero random words.
	newRandInt := func(numWords int) Int {
		var z Int
		z.initiateAbs()
		for i := 0; i < numWords; i++ {
			z.abs[i] = rnd.Uint64()
		}
		return z
	}

	for i := 0; i < numSamples; i++ {
		x32g := rnd.Uint32()
		x32l := rnd.Uint32()
		if x32g < x32l {
			x32g, x32l = x32l, x32g
		}
		lp := newRandInt(4)
		gp := newRandInt(4)
		if gp.Cmp(&lp) < 0 {
			gp, lp = lp, gp
		}
		if gp.abs[3] == 0 {
			gp.abs[3]++
		}
		int256Samples[i] = gp
		big256Samples[i] = *int256Samples[i].ToBig()
		int256SamplesLt[i] = lp
		big256SamplesLt[i] = *int256SamplesLt[i].ToBig()
		ln := newRandInt(4)
		gn := newRandInt(4)
		ln.neg = true
		gn.neg = true
		if gn.Cmp(&ln) < 0 {
			gn, ln = ln, gn
		}
		if gn.abs[3] == 0 {
			gn.abs[3]++
		}
		int256SamplesNeg[i] = gn
		big256SamplesNeg[i] = *int256SamplesNeg[i].ToBig()
		int256SamplesLtNeg[i] = ln
		big256SamplesLtNeg[i] = *int256SamplesLtNeg[i].ToBig()
	}

	return true
}

type marshalTest struct {
	input interface{}
	want  string
}

type unmarshalTest struct {
	input   string
	want    interface{}
	wantErr error // if set, decoding must fail on any platform
}

func referenceBig(s string) *big.Int {
	b, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("invalid")
	}
	return b
}

func checkError(t *testing.T, input string, got, want error) bool {
	if got == nil {
		if want != nil {
			t.Errorf("input %s: got no error, want %q", input, want)
			return false
		}
		return true
	}
	if want == nil {
		t.Errorf("input %s: unexpected error %q", input, got)
	} else if got.Error() != want.Error() {
		t.Errorf("input %s: got error %q, want %q", input, got, want)
	}
	return false
}

// causesPanic returns true if panic occurred when executing fn.
func causesPanic(fn func()) bool {
	done := make(chan struct{})
	var ok bool
	go func() {
		defer func() {
			ok = recover() != nil
			done <- struct{}{}
		}()
		fn()
	}()
	<-done
	return ok
}

func TestEncode(t *testing.T) {
	for _, test := range encodeBigTests {
		z, _ := FromBig(test.input.(*big.Int))
		enc := z.Hex()
		if enc != test.want {
			t.Errorf("input %x: wrong encoding %s (exp %s)", test.input, enc, test.want)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, test := range decodeBigTests {
		dec, err := FromHex(test.input)
		if err != nil {
			if !causesPanic(func() { MustFromHex(test.input) }) {
				t.Fatalf("expected panic")
			}
		}
		if !checkError(t, test.input, err, test.wantErr) {
			continue
		}
		b := dec.ToBig()
		if b.Cmp(test.want.(*big.Int)) != 0 {
			t.Errorf("input %s: value mismatch: got %x, want %x", test.input, dec, test.want)
			continue
		}
		d2 := MustFromHex(test.input)
		if d2.Cmp(dec) != 0 {
			t.Errorf("input %s: value mismatch: got %x, want %x", test.input, d2, dec)
		}
	}
	// Some remaining json-tests
	type jsonStruct struct {
		Foo *Int
	}
	var jsonDecoded jsonStruct
	// This test was previously an "expected error", The U256 behaviour has now
	// changed, to be compatible with big.Int
	if err := json.Unmarshal([]byte(`{"Foo":1}`), &jsonDecoded); err != nil {
		t.Fatalf("Expected no error, have %v", err)
	}
	if err := json.Unmarshal([]byte(`{"Foo":0x1}`), &jsonDecoded); err == nil {
		t.Fatal("Expected error")
	}
	if err := json.Unmarshal([]byte(`{"Foo":""}`), &jsonDecoded); err == nil {
		t.Fatal("Expected error")
	}
	if err := json.Unmarshal([]byte(`{"Foo":"0x1"}`), &jsonDecoded); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	} else if jsonDecoded.Foo.Int64() != 1 {
		t.Fatal("Expected 1")
	}
}

func TestEnDecode(t *testing.T) {
	type jsonStruct struct {
		Foo *Int
	}
	type jsonBigStruct struct {
		Foo *big.Int
	}
	var testSample = func(i int, bigSample big.Int, intSample Int) {
		// Encoding
		var wantHex string
		if intSample.neg {
			wantHex = fmt.Sprintf("-0x%s", bigSample.Text(16)[1:])
		} else {
			wantHex = fmt.Sprintf("0x%s", bigSample.Text(16))
		}
		wantDec := bigSample.Text(10)

		if intSample.neg {
			if have, want := intSample.Hex(), fmt.Sprintf("-0x%s", bigSample.Text(16)[1:]); have != want {
				t.Fatalf("test %d #1, have %v, want %v", i, have, want)
			}
		} else {
			if have, want := intSample.Hex(), fmt.Sprintf("0x%s", bigSample.Text(16)); have != want {
				t.Fatalf("test %d #1, have %v, want %v", i, have, want)
			}
		}
		if have, want := intSample.String(), bigSample.String(); have != want {
			t.Fatalf("test %d String(), have %v, want %v", i, have, want)
		}
		{
			have, _ := intSample.MarshalText()
			want, _ := bigSample.MarshalText()
			if !bytes.Equal(have, want) {
				t.Fatalf("test %d MarshalText, have %q, want %q", i, have, want)
			}
		}
		{
			have, _ := intSample.MarshalJSON()
			want := []byte(fmt.Sprintf(`"%s"`, bigSample.Text(10)))
			if !bytes.Equal(have, want) {
				t.Fatalf("test %d MarshalJSON, have %q, want %q", i, have, want)
			}
		}
		if have, _ := intSample.Value(); wantDec != have.(string) {
			t.Fatalf("test %d #4, got %v, exp %v", i, have, wantHex)
		}
		if have, want := intSample.Dec(), wantDec; have != want {
			t.Fatalf("test %d Dec(), have %v, want %v", i, have, want)
		}
		{ // Json
			jsonEncoded, err := json.Marshal(&jsonStruct{&intSample})
			if err != nil {
				t.Fatalf("test %d: json encoding err: %v", i, err)
			}
			jsonEncodedBig, _ := json.Marshal(&jsonBigStruct{&bigSample})
			var jsonDecoded jsonStruct
			err = json.Unmarshal(jsonEncoded, &jsonDecoded)
			if err != nil {
				t.Fatalf("test %d error unmarshaling: %v", i, err)
			}
			if jsonDecoded.Foo.Cmp(&intSample) != 0 {
				t.Fatalf("test %d #8, have %v, want %v", i, jsonDecoded.Foo, intSample)
			}
			// See if we can also unmarshal from big.Int's non-string format
			err = json.Unmarshal(jsonEncodedBig, &jsonDecoded)
			if err != nil {
				t.Fatalf("test %d unmarshalling from big.Int err: %v", i, err)
			}
			if jsonDecoded.Foo.Cmp(&intSample) != 0 {
				t.Fatalf("test %d have %v, want %v", i, jsonDecoded.Foo, intSample)
			}
		}
		// Decoding
		//
		// FromHex
		decoded, err := FromHex(wantHex)
		{
			if err != nil {
				t.Fatalf("test %d #9, err: %v", i, err)
			}
			if decoded.Cmp(&intSample) != 0 {
				t.Fatalf("test %d #10, got %v, exp %v", i, decoded, intSample)
			}
		}
		// z.SetFromHex
		err = decoded.SetFromHex(wantHex)
		{
			if err != nil {
				t.Fatalf("test %d #11, err: %v", i, err)
			}
			if decoded.Cmp(&intSample) != 0 {
				t.Fatalf("test %d #12, got %v, exp %v", i, decoded, intSample)
			}
		}
		// UnmarshalText
		decoded = new(Int)
		{
			if err := decoded.UnmarshalText([]byte(wantHex)); err != nil {
				fmt.Println(wantHex)
				t.Fatalf("test %d #13, err: %v", i, err)
			}
			if decoded.Cmp(&intSample) != 0 {
				t.Fatalf("test %d #14, got %v, exp %v", i, decoded, intSample)
			}
		}
		// FromDecimal
		decoded, err = FromDecimal(wantDec)
		{
			if err != nil {
				t.Fatalf("test %d #15, err: %v", i, err)
			}
			if decoded.Cmp(&intSample) != 0 {
				t.Fatalf("test %d #16, got %v, exp %v", i, decoded, intSample)
			}
		}
	}
	for i, bigSample := range big256Samples {
		intSample := int256Samples[i]
		testSample(i, bigSample, intSample)
	}

	for i, bigSample := range big256SamplesLt {
		intSample := int256SamplesLt[i]
		testSample(i, bigSample, intSample)
	}
	for i, bigSample := range big256SamplesNeg {
		intSample := int256SamplesNeg[i]
		testSample(i, bigSample, intSample)
	}

	for i, bigSample := range big256SamplesLtNeg {
		intSample := int256SamplesLtNeg[i]
		testSample(i, bigSample, intSample)
	}
}

func TestDecimal(t *testing.T) {
	for i := uint(0); i < 255; i++ {
		a := NewInt(1)
		a.Lsh(a, i)
		want := a.ToBig().Text(10)
		if have := a.Dec(); have != want {
			t.Errorf("want '%v' have '%v', \n", want, have)
		}
		// Op must not modify the original
		if have := a.Dec(); have != want {
			t.Errorf("want '%v' have '%v', \n", want, have)
		}
		// same for negative values
		a.Mul(a, NewInt(-1))
		want = a.ToBig().Text(10)
		if have := a.Dec(); have != want {
			t.Errorf("want '%v' have '%v', \n", want, have)
		}
		// Op must not modify the original
		if have := a.Dec(); have != want {
			t.Errorf("want '%v' have '%v', \n", want, have)
		}
	}
	// test zero-case
	if have, want := new(Int).Dec(), new(big.Int).Text(10); have != want {
		t.Errorf("have '%v', want '%v'", have, want)
	}
	{ // max
		maxi, _ := FromHex("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
		maxb, _ := new(big.Int).SetString("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0)
		if have, want := maxi.Dec(), maxb.Text(10); have != want {
			t.Errorf("have '%v', want '%v'", have, want)
		}
	}
	{
		maxi, _ := FromDecimal("29999999999999999999")
		maxb, _ := new(big.Int).SetString("29999999999999999999", 0)
		if have, want := maxi.Dec(), maxb.Text(10); have != want {
			t.Errorf("have '%v', want '%v'", have, want)
		}
	}
}
