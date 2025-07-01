package main

import "fmt"

type LNode struct {
	Data interface{}
	Next *LNode
}

func CreateNode(node *LNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

func Reverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var pre *LNode
	var cur *LNode

	next := node.Next

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}

func main() {
	head := &LNode{}
	fmt.Println("sort")
	CreateNode(head, 10)
	PrintNode("before sort : ", head)

	Reverse(head)
	PrintNode("after sort: ", head)
}
