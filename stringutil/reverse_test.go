// Package stringutil contains utility functions for working with strings.

package stringutil

import "testing"

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Hello, world", args{"Hello, world"}, "dlrow ,olleH"},
		{"Hello, 世界", args{"Hello, 世界"}, "界世 ,olleH"},
		{"", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
