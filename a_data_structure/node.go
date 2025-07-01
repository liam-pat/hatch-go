package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) getValue() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("the node is nil")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.getValue()
	node.right.traverse()
}

type myTreeNode struct {
	node *treeNode
}

func (node *myTreeNode) postOrder() {
	if node == nil || node.node == nil {
		return
	}
	left := myTreeNode{node.node.left}
	right := myTreeNode{node.node.right}
	left.postOrder()
	right.postOrder()
	node.node.getValue()
}

func main() {
	var tree treeNode

	tree = treeNode{value: 3}
	tree.left = &treeNode{value: 4, left: nil, right: nil}
	tree.right = &treeNode{5, nil, &treeNode{value: 10, right: nil, left: nil}}
	tree.traverse()

	fmt.Println()
	myNode := myTreeNode{&tree}
	myNode.postOrder()
	fmt.Println()
}
