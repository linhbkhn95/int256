package int256

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
)

func TestInt_UnmarshalJSON(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	type args struct {
		input []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    *Int
	}{
		{
			name: "Should perform correctly when unmarshal from marshal data from data",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				input: marshalJSON(big.NewInt(10)),
			},
			wantErr: false,
			want: &Int{
				abs: uint256.NewInt(10),
				neg: false,
			},
		},
		{
			name: "Should perform correctly when data is string bytes(positive number)",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				input: []byte("10"),
			},
			wantErr: false,
			want: &Int{
				abs: uint256.NewInt(10),
				neg: false,
			},
		},
		{
			name: "Should perform correctly when data is string bytes(negative number)",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				input: []byte("-10"),
			},
			wantErr: false,
			want: &Int{
				abs: uint256.NewInt(10),
				neg: true,
			},
		},
		{
			name: "Should err when data is not number",
			fields: fields{
				abs: new(uint256.Int),
				neg: false,
			},
			args: args{
				input: []byte("-a"),
			},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			err := z.UnmarshalJSON(tt.args.input)
			if tt.wantErr {
				assert.NotEqual(t, nil, err)

			} else {
				assert.Equal(t, nil, err)
				assert.Equal(t, tt.want, z)
			}

		})
	}
}

func marshalJSON(bigInt *big.Int) []byte {
	bytes, _ := bigInt.MarshalJSON()
	fmt.Println(string(bytes))
	return bytes
}

func TestInt_MarshalJSON(t *testing.T) {

	tests := []struct {
		name    string
		input   *Int
		want    []byte
		wantErr bool
	}{
		{
			name:    "Should return correct when int256 is positive number",
			input:   NewInt(10),
			want:    []byte("10"),
			wantErr: false,
		},
		{
			name:    "Should return correct when int256 is negative number",
			input:   NewInt(-10),
			want:    []byte("-10"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.input.MarshalJSON()
			if tt.wantErr {
				assert.NotEqual(t, nil, err)

			} else {
				assert.Equal(t, nil, err)
				assert.Equal(t, tt.want, got)

			}
		})
	}
}
