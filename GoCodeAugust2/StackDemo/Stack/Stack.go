package stack

import(
	"sync"
)

type Item interface{}

type Stack struct{
	items []Item
	mu sync.Mutex
}

func (s *Stack) Push(i Item){
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, i)
}

func (s *Stack) Pop()Item{
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.items)==0{
		return nil
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *Stack) IsEmpty()bool{
	return len(s.items)==0
}