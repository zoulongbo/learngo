package main

import "testing"

func TestNonRepeat(t *testing.T) {
	tests := []struct{s string; c int}{
		{"1ecf423rcf42ef1erc1", 8},
		{"我的山鸡舞镜我饿收到欧水情人我去玩儿欧水啊", 13},
		{"11121312341234132413413424213413", 4},
		{"1122231", 3},
		{" ", 1},
	}
	for _, tt := range tests {
		if actual := lengthOfNonRepeatSubstr(tt.s); actual != tt.c {
			t.Errorf("%s lengthOfNonRepeatSubstr: %d  expected:%d\n", tt.s, actual, tt.c)
		}
	}
}

func BenchmarkNonRepeat(t *testing.B){
	s := "1ecf423rcf42ef1erc1"
	c := 8
	for i:=0; i<100; i++ {
		if actual := lengthOfNonRepeatSubstr(s); actual != c {
			t.Errorf("%s lengthOfNonRepeatSubstr: %d  expected:%d\n", s, actual, c)
		}
	}

}