package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{a, b, c int} {
		{3,4,5},
		{5,12,13},
		{8,15,17},
		{12,35,37},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d %d);  got %d;  expected %d", tt.a, tt.b, actual, tt.c )
		}
	}
}

func BenchmarkTriangle(t *testing.B) {
	a := 30000
	b := 40000
	c := 50000

	for i:=0; i < 100; i++ {
		if actual := calcTriangle(a, b); actual != c {
			t.Errorf("calcTriangle(%d %d);  got %d;  expected %d", a, b, actual, c )
		}
	}
}
