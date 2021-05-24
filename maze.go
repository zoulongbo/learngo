package main

import (
	"fmt"
	"github.com/zoulongbo/learngo/tree"
)

//方法重写	内嵌 可不起名，默认取引入量的名字
type myTreeNode struct {
	 *tree.Node
}

//方法重写
func (myNode *myTreeNode) myTraverse()  {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	left.myTraverse()
	right := myTreeNode{myNode.Right}
	right.myTraverse()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse()  {
	fmt.Println("This is shadowed method .")
}

func main() {
	var master = myTreeNode{&tree.Node{Value:0}}

	master.Left = tree.CreateNode(1)
	master.Left.Left = tree.CreateNode(3)
	master.Right = tree.CreateNode(2)
	master.Right.Right = tree.CreateNode(4)

	/*nodes := []treeNode{
		{value: 3},
		{},
		{111, &treeNode{}, &treeNode{}},
	}*/
	//fmt.Println(nodes)
	master.Left.SetValue(11)
	//master.left.print()
	master.Left.SetValues(11)
	//master.left.print()
	master.Traverse()
	fmt.Println()

	//master.Traverse()
	master.myTraverse()

	nodeCount := 0
	master.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("nodeCount:", nodeCount)


	//TraverseFunc
	maxValue := 0
	c := master.TraverseWithChannel()
	for node := range c {
		if node.Value > maxValue {
			maxValue = node.Value
		}
	}

	fmt.Println("The max Value is :", maxValue)
}
