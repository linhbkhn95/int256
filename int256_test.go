package int256

import (
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
					abs: uint256.NewInt(10),
				},
				y: &Int{
					abs: uint256.NewInt(7),
				},
			},
			want: &Int{
				abs: uint256.NewInt(17),
				neg: false,
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
					abs: uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: uint256.NewInt(7),
					neg: true,
				},
			},
			want: &Int{
				abs: uint256.NewInt(17),
				neg: true,
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
					abs: uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: uint256.NewInt(7),
					neg: false,
				},
			},
			want: &Int{
				abs: uint256.NewInt(3),
				neg: true,
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
					abs: uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: uint256.NewInt(7),
					neg: true,
				},
			},
			want: &Int{
				abs: uint256.NewInt(3),
				neg: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.Add(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
