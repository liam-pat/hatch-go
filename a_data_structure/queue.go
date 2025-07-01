package main

import "fmt"

type queue []int

func (q *queue) Push(value int) {
	*q = append(*q, value)
}

func (q *queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func main() {
	queue := &queue{1}
	fmt.Printf("empty: %t\n", queue.IsEmpty())

	queue.Push(10)
	queue.Push(20)
	fmt.Println("queue:", queue)

	queue.Pop()
	queue.Pop()
	fmt.Println("queue:", queue)
}
