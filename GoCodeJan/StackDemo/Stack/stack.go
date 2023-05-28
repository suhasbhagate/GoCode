package stack

import(
	"sync"
)

type Item interface{}

type Stack struct{
	mu sync.Mutex
	items []Item
}

func (st *Stack) Push(i Item){
	st.mu.Lock()
	defer st.mu.Unlock()
	st.items = append(st.items, i)
}

func (st *Stack) Pop()Item{
	st.mu.Lock()
	defer st.mu.Unlock()
	if len(st.items)==0{
		return nil
	}
	element:= st.items[len(st.items)-1]
	st.items = st.items[:len(st.items)-1]
	return element
}

func (st *Stack) IsEmpty()bool{
	return len(st.items)==0
}

func (st *Stack) Top()Item{
	if len(st.items) ==0{
		return nil
	}
	return st.items[len(st.items)-1]
}