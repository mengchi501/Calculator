package elements

type Stack struct {
	sp  *Item
	num int
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.num = 0
	return stack
}

func (stack *Stack) Push(item interface{}) {
	stack.sp = &Item{item: item, next: stack.sp}
	stack.num++
}

func (stack *Stack) Pop() interface{} {
	if stack.num > 0 {
		item := stack.sp.item
		stack.sp = stack.sp.next
		stack.num--
		return item
	}
	return nil
}

type Item struct {
	item interface{}
	next *Item
}
