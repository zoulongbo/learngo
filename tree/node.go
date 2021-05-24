package tree

import (
	"fmt"
)

type Node struct {
	Value int
	Left,Right	*Node
}

func (node Node) Print()  {
	fmt.Println(node.Value, " ")
}

func (node Node) SetValue(value int) {
	node.Value= value
}

func (node *Node) SetValues(value int) {
	if node == nil {
		fmt.Println("Setting value to nil" + " node . Ignored")
		return
	}
	node.Value = value
}

func (node *Node) Traverse()  {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func CreateNode(value int) *Node{
	return &Node{Value:value}
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)

	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
