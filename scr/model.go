package scr

type Pair struct {
	Wight int
	Price int
}

type NodeDfs struct {
	Edges  int
	Vertex []int
}
type Graph map[int][]int

type Stack struct { //stack
	values []int
	top    int
}

func NewStack() *Stack {
	return &Stack{
		values: make([]int, 0),
		top:    -1,
	}
}

// Push adds a value to the top of the stack.
func (st *Stack) Push(value int) {
	st.top++
	st.values = append(st.values, value)
}

// Pop removes and returns the top value from the stack.
func (st *Stack) Pop() int {
	if st.IsEmpty() {
		return 0
	}
	value := st.values[st.top]
	st.top--
	return value
}

func (st *Stack) IsEmpty() bool {
	return st.top == -1
}
