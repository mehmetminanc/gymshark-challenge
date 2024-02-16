package algo

import (
	"reflect"
	"testing"
)

func Test_computePacking(t *testing.T) {
	type args struct {
		packs []int
		order int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "1",
			args: args{
				packs: []int{250, 500, 1000, 2000, 5000},
				order: 1,
			},
			want: map[int]int{
				250: 1,
			},
		},
		{
			name: "250",
			args: args{
				packs: []int{250, 500, 1000, 2000, 5000},
				order: 250,
			},
			want: map[int]int{
				250: 1,
			},
		},
		{
			name: "251",
			args: args{
				packs: []int{250, 500, 1000, 2000, 5000},
				order: 251,
			},
			want: map[int]int{
				500: 1,
			},
		},
		{
			name: "501",
			args: args{
				packs: []int{250, 500, 1000, 2000, 5000},
				order: 501,
			},
			want: map[int]int{
				250: 1,
				500: 1,
			},
		},
		{
			name: "12001",
			args: args{
				packs: []int{250, 500, 1000, 2000, 5000},
				order: 12001,
			},
			want: map[int]int{
				250:  1,
				2000: 1,
				5000: 2,
			},
		},
		{
			name: "263 small packs",
			args: args{
				packs: []int{23, 31, 53},
				order: 263,
			},
			want: map[int]int{
				23: 2,
				31: 7,
			},
		},
		{
			name: "263",
			args: args{
				packs: []int{23, 31, 53},
				order: 250,
			},
			want: map[int]int{
				23: 4,
				53: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputePacking(tt.args.packs, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computePacking() = %v, want %v", got, tt.want)
			}
		})
	}
}
