package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCity(contents, "http://m.zhenai.com/zhenghun/suzhou/nv")
	if len(result.Requests) == 0 {
		t.Errorf("Result should better than 1 Requests ; but was %v", result.Requests)
	}
	 for _, url := range result.Requests {
		 t.Logf("Result Requests %v", url)
	}

}
