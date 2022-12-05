package util

import "testing"

func TestRoundToDec(t *testing.T) {
	type args struct {
		val float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "RoundToDec",
			args: args{
				val: 1.234,
			},
			want: 1.23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundToDec(tt.args.val); got != tt.want {
				t.Errorf("RoundToDec() = %v, want %v", got, tt.want)
			}
		})
	}
}
