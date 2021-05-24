package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseProfileMock(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfileMock(contents, "昨夜星辰","http://album.zhenai.com/u/1496876501")

	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element ; but was %v", result.Items)
	}
}