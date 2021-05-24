package main

import (
	"fmt"
	"github.com/zoulongbo/learngo/rpcdemo"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)

	var result float64

	err = client.Call("DemoService.Div", rpcdemo.Args{
		A: 19,
		B: 2,
	}, &result)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}



	err = client.Call("DemoService.Div", rpcdemo.Args{
		A: 28,
		B: 0,
	}, &result)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}
