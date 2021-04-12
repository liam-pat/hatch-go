package main

import "fmt"

type Queue []int

func (q *Queue) Push(value int) {
	*q = append(*q, value)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func main() {
	myQueue := Queue{1}

	fmt.Println("is empty: ", myQueue.IsEmpty())

	myQueue.Push(10)

	fmt.Println(myQueue)

	myQueue.Pop()
	myQueue.Pop()
	fmt.Println(myQueue.IsEmpty())
}
