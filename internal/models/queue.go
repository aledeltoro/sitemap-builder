package models

// Queue data structure that performs operations perform in
// First in First Out (FIFO) order
type Queue struct {
	items []string
}

// NewQueue creates a new Queue instances
func NewQueue() *Queue {
	return &Queue{
		items: make([]string, 0),
	}
}

// Peek returns the first item in the Queue
func (q *Queue) Peek() string {
	if len(q.items) == 0 {
		return ""
	}

	return q.items[0]
}

// Enqueue adds an item to the back of the queue
func (q *Queue) Enqueue(item string) {
	q.items = append(q.items, item)
}

// DeQueue removes the item at the front of the Queue
func (q *Queue) DeQueue() string {
	if len(q.items) == 0 {
		return ""
	}

	removedItem := q.items[0]

	newItems := make([]string, 0, len(q.items))
	copy(newItems, q.items[1:])
	q.items = newItems

	return removedItem
}

// Size returns the amount of items in the Queue
func (q *Queue) Size() int {
	return len(q.items)
}
