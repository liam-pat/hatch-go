package godoc

// Queue an int queue
type Queue []int

// Push push into queue
func (q *Queue) Push(value int) {
	*q = append(*q, value)
}

// Pop pop into queue
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty check queue IsEmpty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
