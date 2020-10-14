package godoc

// an int queue
type Queue []int

// push into queue
func (q *Queue) Push(value int) {
	*q = append(*q, value)
}

// pop into queue
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//check queue IsEmpty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}