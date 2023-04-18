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
					abs: uint256.NewInt(10),
				},
				y: &Int{
					abs: uint256.NewInt(7),
				},
			},
			want: &Int{
				abs: uint256.NewInt(3),
				neg: false,
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
					abs: uint256.NewInt(7),
				},
				y: &Int{
					abs: uint256.NewInt(10),
				},
			},
			want: &Int{
				abs: uint256.NewInt(3),
				neg: true,
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
					abs: uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: uint256.NewInt(7),
					neg: true,
				},
			},
			want: &Int{
				abs: uint256.NewInt(3),
				neg: true,
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
					abs: uint256.NewInt(7),
					neg: true,
				},
				y: &Int{
					abs: uint256.NewInt(10),
					neg: true,
				},
			},
			want: &Int{
				abs: uint256.NewInt(3),
				neg: false,
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
				abs: uint256.NewInt(17),
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
				abs: uint256.NewInt(17),
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
					abs: uint256.NewInt(3),
				},
				y: &Int{
					abs: uint256.NewInt(5),
				},
			},
			want: &Int{
				abs: uint256.NewInt(15),
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
					abs: uint256.NewInt(3),
					neg: true,
				},
				y: &Int{
					abs: uint256.NewInt(5),
					neg: true,
				},
			},
			want: &Int{
				abs: uint256.NewInt(15),
				neg: false,
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
					abs: uint256.NewInt(3),
					neg: true,
				},
				y: &Int{
					abs: uint256.NewInt(5),
					neg: false,
				},
			},
			want: &Int{
				abs: uint256.NewInt(15),
				neg: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.Mul(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Mul() = %v, want %v", got, tt.want)
			}
		})
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
					abs: uint256.NewInt(9),
				},
			},
			want: &Int{
				abs: uint256.NewInt(3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
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
				abs: uint256.NewInt(10),
				neg: false,
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
				abs: uint256.NewInt(10),
				neg: true,
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
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
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
