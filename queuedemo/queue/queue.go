package queue

import (
	"sync"
)

type Item interface{}

type Queue struct{
	items []Item
	mu sync.Mutex
}

func (q *Queue) Enqueue(i Item){
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue()Item{
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items)==0{
		return nil
	}
	element := q.items[0]
	q.items = q.items[1:]
	return element
}

func (q Queue) Peek() Item{
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0{
		return nil
	}
	return q.items[0]
}

func (q *Queue) IsEmpty()bool{
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)==0
}