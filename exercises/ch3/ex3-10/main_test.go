package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"1", "1"},
		{"12", "12"},
		{"120", "120"},
		{"1776", "1,776"},
		{"12345", "12,345"},
		{"123456789", "123,456,789"},
	}

	for _, tt := range tests {
		got := comma(tt.input)
		if got != tt.want {
			t.Errorf("TestComma - %v: got %v, want %v \n", tt, got, tt.want)
		}
	}
}
