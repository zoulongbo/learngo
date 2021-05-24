package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {

	req, err := http.NewRequest(http.MethodGet, "https://www.imooc.com", nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	if err != nil {
		panic(err)
	}
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req)
			return nil
		},
	}
	resp, err := client.Do(req)
	//resp, err := http.DefaultClient.Do(req)
	//resp, err := http.Get("https://www.imooc.com")
	if err != nil {
		panic(err)
	}

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", s)
}
