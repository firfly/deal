package thinking

import "testing"

func Test_cal8queens(t *testing.T) {
	type args struct {
		row int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"1",
			args{row: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cal8queens(tt.args.row)
		})
	}
}
