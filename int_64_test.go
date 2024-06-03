package int256

import (
	"testing"

	"github.com/holiman/uint256"
)

func TestInt_Int64(t *testing.T) {
	type fields struct {
		abs *uint256.Int
		neg bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when parsing positive number",
			fields: fields{
				abs: uint256.NewInt(10),
				neg: false,
			},
			want: 10,
		},
		{
			name: "Should return correct value when parsing negative number",
			fields: fields{
				abs: uint256.NewInt(10),
				neg: true,
			},
			want: -10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				Abs: tt.fields.abs,
				Neg: tt.fields.neg,
			}
			if got := z.Int64(); got != tt.want {
				t.Errorf("Int.Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}
