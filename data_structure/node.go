package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

type myTreeNode struct {
	node *treeNode
}

func (node treeNode) print() {
	fmt.Print(node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {

	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func (node *myTreeNode) postOrder() {
	if node == nil || node.node == nil {
		return
	}

	left := myTreeNode{node.node.left}
	right := myTreeNode{node.node.right}
	left.postOrder()
	right.postOrder()
	node.node.print()
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, createNode(10)}

	root.traverse()
	fmt.Println()
	myNode := myTreeNode{&root}
	myNode.postOrder()
	fmt.Println()
}
