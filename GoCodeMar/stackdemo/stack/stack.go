package stack

import(
	"sync"
)

type Item interface{}

type Stack struct{
	items []Item
	mu sync.Mutex
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
	st.mu.Lock()
	defer st.mu.Unlock()
	return len(st.items)==0
}

func (st *Stack) Peek()Item{
	st.mu.Lock()
	defer st.mu.Unlock()
	if len(st.items)==0{
		return nil
	}
	element:= st.items[len(st.items)-1]
	return element
}