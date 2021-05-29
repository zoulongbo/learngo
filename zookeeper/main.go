package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

var (
	host = []string{"127.0.0.1:2181"}
)

func main()  {
	conn, _, err := zk.Connect(host, 5 * time.Second)
	if err != nil {
		panic("conn err: " + err.Error())
	}
	defer conn.Close()
	//新增
	if _, err = conn.Create("/test_tree", []byte("tree content"), 0, zk.WorldACL(zk.PermAll)); err != nil {
		fmt.Println("create err: ", err)
		return
	}
	//查询
	nodeV, dStat, err := conn.Get("/test_tree")
	if err != nil {
		fmt.Println("get err: ", err)
		return
	}
	fmt.Println("nodeValue:", string(nodeV))

	//改
	if _, err = conn.Set("/test_tree", []byte("tree content1"), dStat.Version); err != nil {
		fmt.Println("update err: ", err)
		return
	}
	//删
	_, dStat, _ = conn.Get("/test_tree")
	if err = conn.Delete("/test_tree", dStat.Version); err != nil {
		fmt.Println("delete err: ", err)
		return
	}
	//验证
	exists, _, err := conn.Exists("/test_tree")
	if err != nil {
		fmt.Println("exists err: ", err)
		return
	}
	fmt.Println("exists:", exists)

	//增加
	if _, err = conn.Create("/test_tree1", []byte("tree content11"), 0, zk.WorldACL(zk.PermAll)); err != nil {
		fmt.Println("once create err: ", err)

	}
	//设置子节点
	if _, err = conn.Create("/test_tree1/sub", []byte("tree content11 - sub"), 0, zk.WorldACL(zk.PermAll)); err != nil {
		fmt.Println("create sub err: ", err)

	}
	//获取子节点
	children, _, err := conn.Children("/test_tree1")
	if err != nil {
		fmt.Println("children err: ", err)
		return
	}
	fmt.Println("children: ", children)
}
