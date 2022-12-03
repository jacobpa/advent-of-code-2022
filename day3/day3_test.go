package day3

import (
	"testing"
)

func Test_calculatePriority(t *testing.T) {
	type args struct {
		item byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				item: 'a',
			},
			want: 1,
		},
		{
			name: "z",
			args: args{
				item: 'z',
			},
			want: 26,
		},
		{
			name: "A",
			args: args{
				item: 'A',
			},
			want: 27,
		},
		{
			name: "Z",
			args: args{
				item: 'Z',
			},
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculatePriority(tt.args.item); got != tt.want {
				t.Errorf("calculatePriority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findBadge(t *testing.T) {
	type args struct {
		first  string
		second string
		third  string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "Example 1",
			args: args{
				first:  "vJrwpWtwJgWrhcsFMMfFFhFp",
				second: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				third:  "PmmdzqPrVvPwwTWBwg",
			},
			want: 'r',
		},
		{
			name: "Example 2",
			args: args{
				first:  "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				second: "ttgJtRGJQctTZtZT",
				third:  "CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			want: 'Z',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBadge(tt.args.first, tt.args.second, tt.args.third); got != tt.want {
				t.Errorf("findBadge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCommonItem(t *testing.T) {
	type args struct {
		sack string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "Example 1",
			args: args{
				sack: "vJrwpWtwJgWrhcsFMMfFFhFp",
			},
			want: 'p',
		},
		{
			name: "Example 2",
			args: args{
				sack: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			},
			want: 'L',
		},
		{
			name: "Example 3",
			args: args{
				sack: "PmmdzqPrVvPwwTWBwg",
			},
			want: 'P',
		},
		{
			name: "Example 4",
			args: args{
				sack: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			},
			want: 'v',
		},
		{
			name: "Example 5",
			args: args{
				sack: "ttgJtRGJQctTZtZT",
			},
			want: 't',
		},
		{
			name: "Example 6",
			args: args{
				sack: "CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			want: 's',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCommonItem(tt.args.sack); got != tt.want {
				t.Errorf("findCommonItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
