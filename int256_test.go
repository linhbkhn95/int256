package int256

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/holiman/uint256"
)

func TestInt_Add(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct when add two positive numbers",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
				},
				y: &Int{
					Abs: uint256.NewInt(7),
				},
			},
			want: &Int{
				Abs: uint256.NewInt(17),
				Neg: false,
			},
		},
		{
			name: "Should return correct when add two negative numbers",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(7),
					Neg: true,
				},
			},
			want: &Int{
				Abs: uint256.NewInt(17),
				Neg: true,
			},
		},
		{
			name: "Should return correct when add two numbers with a is negative, b is positive and |a|>|b|",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(7),
					Neg: false,
				},
			},
			want: &Int{
				Abs: uint256.NewInt(3),
				Neg: true,
			},
		},
		{
			name: "Should return correct when add two numbers with a is negative, b is positive and |a|<|b|",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				y: &Int{
					Abs: uint256.NewInt(7),
					Neg: true,
				},
			},
			want: &Int{
				Abs: uint256.NewInt(3),
				Neg: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			if got := z.Add(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Sub(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct when sub two positive numbers |a|>|b|",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
				},
				y: &Int{
					Abs: uint256.NewInt(7),
				},
			},
			want: &Int{
				Abs: uint256.NewInt(3),
				Neg: false,
			},
		},
		{
			name: "Should return correct when sub two positive numbers |a|<|b|",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(7),
				},
				y: &Int{
					Abs: uint256.NewInt(10),
				},
			},
			want: &Int{
				Abs: uint256.NewInt(3),
				Neg: true,
			},
		},
		{
			name: "Should return correct when add two negative numbers |a|>|b|",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(7),
					Neg: true,
				},
			},
			want: &Int{
				Abs: uint256.NewInt(3),
				Neg: true,
			},
		},
		{
			name: "Should return correct when add two negative numbers |a|<|b|",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(7),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
			},
			want: &Int{
				Abs: uint256.NewInt(3),
				Neg: false,
			},
		},
		{
			name: "Should return correct when add two numbers with a is negative, b is positive and |a|>|b|",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(7),
					Neg: false,
				},
			},
			want: &Int{
				Abs: uint256.NewInt(17),
				Neg: true,
			},
		},
		{
			name: "Should return correct when add two numbers with a is negative, b is positive and |a|<|b|",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				y: &Int{
					Abs: uint256.NewInt(7),
					Neg: true,
				},
			},
			want: &Int{
				Abs: uint256.NewInt(17),
				Neg: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			if got := z.Sub(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Mul(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when multiple for two positive numbers",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(3),
				},
				y: &Int{
					Abs: uint256.NewInt(5),
				},
			},
			want: &Int{
				Abs: uint256.NewInt(15),
			},
		},
		{
			name: "Should return correct value when multiple for two negative numbers",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(5),
					Neg: true,
				},
			},
			want: &Int{
				Abs: uint256.NewInt(15),
				Neg: false,
			},
		},
		{
			name: "Should return correct value when multiple for a negative number and a positive number",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(5),
					Neg: false,
				},
			},
			want: &Int{
				Abs: uint256.NewInt(15),
				Neg: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			if got := z.Mul(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_MulPanic(t *testing.T) {
	i256 := new(Int)

	defer func() {
		if r := recover(); r != nil {
			t.Error("should not have paniced", r)
		}
	}()

	res := i256.Mul(NewInt(1), NewInt(2))
	if res.Cmp(NewInt(2)) != 0 {
		t.Errorf("want: 2, got: %v", res)
	}
}

func TestInt_Sqrt(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		x *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when performing for positive number",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(9),
				},
			},
			want: &Int{
				Abs: uint256.NewInt(3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			if got := z.Sqrt(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_SetString(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		s string
	}
	big1, _ := new(big.Int).SetString("-10a", 16)

	var MaxUint256, _ = new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)

	big, _ := new(big.Int).SetString("1461446703485210103287273052203988822378723970342", 10)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when parsing correct string value",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				s: "10",
			},
			want: &Int{
				Abs: uint256.NewInt(10),
				Neg: false,
			},
			wantErr: false,
		},
		{
			name: "Should return correct value when parsing correct string value",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				s: "-10",
			},
			want: &Int{
				Abs: uint256.NewInt(10),
				Neg: true,
			},
			wantErr: false,
		},
		{
			name: "Should return error value when parsing incorrect string value",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				s: "-10a",
			},
			want:    MustFromBig(big1),
			wantErr: false,
		},
		{
			name: "Should return error value when parsing correct string value",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				s: "1461446703485210103287273052203988822378723970342",
			},
			want:    MustFromBig(big),
			wantErr: false,
		},
		{
			name: "Should return error value when parsing correct string value",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: false,
			},
			args: args{
				s: "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			},
			want:    MustFromBig(MaxUint256),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			got, err := z.SetString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int.SetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.SetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_SetInt64(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		x int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when ",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: 24,
			},
			want: &Int{
				Abs: uint256.NewInt(24),
				Neg: false,
			},
		},
		{
			name: "Should return correct value when ",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: -24,
			},
			want: &Int{
				Abs: uint256.NewInt(24),
				Neg: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			if got := z.SetInt64(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.SetInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Rsh(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		n uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when perform positive number",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: NewInt(10),
				n: 4,
			},
			want: MustFromBig(new(big.Int).Rsh(big.NewInt(10), 4)),
		},
		{
			name: "Should return correct value when perform negative number",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: NewInt(-10),
				n: 4,
			},
			want: MustFromBig(new(big.Int).Rsh(big.NewInt(-10), 4)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			if got := z.Rsh(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Rsh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Rem(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when performing for two positive numbers",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
				},
				y: &Int{
					Abs: uint256.NewInt(3),
				},
			},
			want: &Int{
				Abs: uint256.NewInt(1),
			},
		},
		{
			name: "Should return correct value when performing for two negative numbers",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Rem(big.NewInt(-10), big.NewInt(-3))),
		},
		{
			name: "Should return correct value when performing for a negative and a numbers",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Rem(big.NewInt(10), big.NewInt(-3))),
		},
		{
			name: "Should return correct value when performing for a negative and a numbers",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Rem(big.NewInt(10), big.NewInt(-3))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			if got := z.Rem(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Rem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Exp(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
		m *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when perform x,y,m is positive number",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: false,
				},
				m: &Int{
					Abs: uint256.NewInt(3),
					Neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Exp(big.NewInt(10), big.NewInt(3), big.NewInt(3))),
		},
		{
			name: "Should return correct value when perform x,y is positive number and m is negative",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: false,
				},
				m: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Exp(big.NewInt(10), big.NewInt(3), big.NewInt(-3))),
		},
		{
			name: "Should return correct value when perform x,y is positive number and m is negative",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
				m: &Int{
					Abs: uint256.NewInt(3),
					Neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Exp(big.NewInt(10), big.NewInt(-3), big.NewInt(3))),
		},
		{
			name: "Should return correct value when perform x,y is positive number and m is negative",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: false,
				},
				m: &Int{
					Abs: uint256.NewInt(3),
					Neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Exp(big.NewInt(-10), big.NewInt(3), big.NewInt(3))),
		},
		{
			name: "Should return correct value when perform x,y is positive number and m is negative",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
				m: &Int{
					Abs: uint256.NewInt(3),
					Neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Exp(big.NewInt(-10), big.NewInt(-3), big.NewInt(3))),
		},
		{
			name: "Should return correct value when perform x,y is positive number and m is negative",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: false,
				},
				m: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
			},

			want: MustFromBig(new(big.Int).Exp(big.NewInt(-10), big.NewInt(3), big.NewInt(-3))),
		},
		{
			name: "Should return correct value when perform x,y is positive number and m is negative",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
				m: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Exp(big.NewInt(10), big.NewInt(-3), big.NewInt(-3))),
		},
		{
			name: "Should return correct value when perform x,y is positive number and m is negative",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
				m: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Exp(big.NewInt(-10), big.NewInt(-3), big.NewInt(-3))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			if got := z.Exp(tt.args.x, tt.args.y, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Exp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Lsh(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		n uint
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		{
			name: "Should return correct value",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				n: 3,
			},
			want: MustFromBig(new(big.Int).Lsh(big.NewInt(10), 3)),
		},
		{
			name: "Should return correct value when process negative number and n is odd",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				n: 3,
			},
			want: MustFromBig(new(big.Int).Lsh(big.NewInt(-10), 3)),
		},
		{
			name: "Should return correct value when process negative number and n is even",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				n: 4,
			},
			want: MustFromBig(new(big.Int).Lsh(big.NewInt(-10), 4)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			if got := z.Lsh(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Lsh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Or(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		{
			name: "Should return correct value",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Or(big.NewInt(10), big.NewInt(3))),
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Or(big.NewInt(-10), big.NewInt(3))),
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Or(big.NewInt(10), big.NewInt(-3))),
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Or(big.NewInt(-10), big.NewInt(-3))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			if got := z.Or(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_And(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: false,
				},
			},
			want: MustFromBig(new(big.Int).And(big.NewInt(10), big.NewInt(3))),
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: false,
				},
			},
			want: MustFromBig(new(big.Int).And(big.NewInt(-10), big.NewInt(3))),
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
			},
			want: MustFromBig(new(big.Int).And(big.NewInt(10), big.NewInt(-3))),
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
			},
			want: MustFromBig(new(big.Int).And(big.NewInt(-10), big.NewInt(-3))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			if got := z.And(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.And() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Quo(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when perform two positive numbers",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Quo(big.NewInt(10), big.NewInt(3))),
		},
		{
			name: "Should return correct value perform a < 0 and b > 0",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Quo(big.NewInt(-10), big.NewInt(3))),
		},
		{
			name: "Should return correct value perform a >0 and b < 0",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: false,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Quo(big.NewInt(10), big.NewInt(-3))),
		},
		{
			name: "Should return correct value perform two negative numbers",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
					Neg: true,
				},
				y: &Int{
					Abs: uint256.NewInt(3),
					Neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Quo(big.NewInt(-10), big.NewInt(-3))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			if got := z.Quo(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Quo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Cmp(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		x *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantR  int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value",
			fields: fields{
				abs: uint256.NewInt(10),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
				},
			},
			wantR: 0,
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: uint256.NewInt(4),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(10),
				},
			},
			wantR: -1,
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: uint256.NewInt(10),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(4),
				},
			},
			wantR: 1,
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: uint256.NewInt(10),
				neg: false,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(0),
				},
			},
			wantR: 1,
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: uint256.NewInt(0),
				neg: true,
			},
			args: args{
				x: &Int{
					Abs: uint256.NewInt(0),
				},
			},
			wantR: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			if gotR := z.Cmp(tt.args.x); gotR != tt.wantR {
				t.Errorf("Int.Cmp() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
