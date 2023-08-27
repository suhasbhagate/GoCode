package queue

import (
	"sync"
)

type Item interface{}

type Queue struct {
	items []Item
	mu    sync.Mutex
}

func (q *Queue) Enqueue(i Item) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() Item {
	q.mu.Lock()
	defer q.mu.Unlock()
	val := q.items[0]
	q.items = q.items[1:]
	return val
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
