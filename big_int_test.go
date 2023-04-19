package int256

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/holiman/uint256"
)

func TestInt_ToBig(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *big.Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when parsing positive number",
			fields: fields{
				abs: uint256.NewInt(10),
				neg: false,
			},
			want: big.NewInt(10),
		},
		{
			name: "Should return correct value when parsing negative number",
			fields: fields{
				abs: uint256.NewInt(10),
				neg: true,
			},
			want: big.NewInt(-10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.ToBig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.ToBig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromBig(t *testing.T) {
	type args struct {
		x *big.Int
	}
	tests := []struct {
		name  string
		args  args
		want  *Int
		want1 bool
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when parsing positive number",
			args: args{
				x: big.NewInt(10),
			},
			want: &Int{
				abs: uint256.NewInt(10),
				neg: false,
			},
			want1: false,
		},
		{
			name: "Should return correct value when parsing negative number",
			args: args{
				x: big.NewInt(-10),
			},
			want: &Int{
				abs: uint256.NewInt(10),
				neg: true,
			},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FromBig(tt.args.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromBig() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FromBig() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
