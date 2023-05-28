package queue

import(
	"sync"
)

type Item interface{}

type Queue struct{
	mu sync.Mutex
	items []Item
}

func (q *Queue) Enque(i Item){
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, i)
}

func (q *Queue) Deque()Item{
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items)==0{
		return nil
	}
	element := q.items[0]
	q.items = q.items[1:]
	return element
}

func (q *Queue) IsEmpty()bool{
	return len(q.items)==0
}

func (q *Queue) Peek()Item{
	if len(q.items)==0{
		return nil
	}
	return q.items[0]
}