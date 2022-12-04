package day4

import (
	"reflect"
	"testing"
)

func Test_encapsulates(t *testing.T) {
	type args struct {
		ranges string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Does not contain",
			args: args{
				ranges: "2-4,6-8",
			},
			want: false,
		},
		{
			name: "First contains second",
			args: args{
				ranges: "2-8,3-7",
			},
			want: true,
		},
		{
			name: "Second contains first",
			args: args{
				ranges: "3-7,2-8",
			},
			want: true,
		},
		{
			name: "First contains second length one",
			args: args{
				ranges: "3-7,7-7",
			},
			want: true,
		},
		{
			name: "Second contains first length one",
			args: args{
				ranges: "7-7,3-7",
			},
			want: true,
		},
		{
			name: "Both one",
			args: args{
				ranges: "7-7,7-7",
			},
			want: true,
		},
		{
			name: "One in middle of range",
			args: args{
				ranges: "5-8,7-7",
			},
			want: true,
		},
		{
			name: "Two distinct",
			args: args{
				ranges: "5-5,7-7",
			},
			want: false,
		},
		{
			name: "Same long range",
			args: args{
				ranges: "0-1000,0-1000",
			},
			want: true,
		},
		{
			name: "Lexicographically true, numerically false",
			args: args{
				"10-99,9-50",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encapsulates(tt.args.ranges); got != tt.want {
				t.Errorf("encapsulates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_overlaps(t *testing.T) {
	type args struct {
		ranges string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Does not contain",
			args: args{
				ranges: "2-4,6-8",
			},
			want: false,
		},
		{
			name: "First contains second",
			args: args{
				ranges: "2-8,3-7",
			},
			want: true,
		},
		{
			name: "Second contains first",
			args: args{
				ranges: "3-7,2-8",
			},
			want: true,
		},
		{
			name: "First contains second length one",
			args: args{
				ranges: "3-7,7-7",
			},
			want: true,
		},
		{
			name: "Second contains first length one",
			args: args{
				ranges: "7-7,3-7",
			},
			want: true,
		},
		{
			name: "Both one",
			args: args{
				ranges: "7-7,7-7",
			},
			want: true,
		},
		{
			name: "One in middle of range",
			args: args{
				ranges: "5-8,7-7",
			},
			want: true,
		},
		{
			name: "Two distinct",
			args: args{
				ranges: "5-5,7-7",
			},
			want: false,
		},
		{
			name: "Same long range",
			args: args{
				ranges: "0-1000,0-1000",
			},
			want: true,
		},
		{
			name: "Lexicographically true, numerically false",
			args: args{
				"10-99,9-50",
			},
			want: true,
		},
		{
			name: "First bigger than second",
			args: args{
				"98-99,96-97",
			},
			want: false,
		},
		{
			name: "Second bigger than first",
			args: args{
				"96-97,98-99",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := overlaps(tt.args.ranges); got != tt.want {
				t.Errorf("encapsulates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encapsulatesWithInvalidInput(t *testing.T) {
	defer func() { recover() }()
	encapsulates("1-2,3")
	t.Errorf("Input should have paniced")
}

func Test_overlapsWithInvalidInput(t *testing.T) {
	defer func() { recover() }()
	overlaps("1-2,3")
	t.Errorf("Input should have paniced")
}

func Test_rangesToIntSlice(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "All one digit",
			args: args{
				input: "1-2,3-4",
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Multiple digits",
			args: args{
				input: "12-345,6789-101112",
			},
			want: []int{12, 345, 6789, 101112},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangesToIntSlice(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rangesToIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
