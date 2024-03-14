package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverse(tree *Tree) {
	if tree == nil {
		return
	}

	traverse(tree.Left)
	fmt.Println(tree.Value, " ")
	traverse(tree.Right)
}

func create(number int) *Tree {
	var tree *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*number; i++ {
		temp := rand.Intn(number * 2)

		tree = insert(tree, temp)
	}

	return tree
}

func insert(tree *Tree, value int) *Tree {
	if tree == nil {
		return &Tree{nil, value, nil}
	}

	if value == tree.Value {
		return tree
	}

	if value < tree.Value {
		tree.Left = insert(tree.Left, value)
		return tree
	}

	tree.Right = insert(tree.Right, value)

	return tree
}

func main() {
	tree := create(10)
	fmt.Println("Print root of the tree : ", tree.Value)

	traverse(tree)
	fmt.Println()

	insert(tree, -10)
	insert(tree, -20)

	traverse(tree)
	fmt.Println()
}
