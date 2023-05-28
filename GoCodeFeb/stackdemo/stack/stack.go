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
	element:= s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return element
}

func (s *Stack) IsEmpty()bool{
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items)==0
}

func (s *Stack) Peek()Item{
	if len(s.items)==0{
		return nil
	}
	return s.items[len(s.items)-1]
}