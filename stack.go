package main

// test implemnt general stack
type Stack struct {
	items []any
	len   int // len(items)
}

func (stack *Stack) push(data any) {
	stack.items = append(stack.items, data)
	stack.len++
}

func (stack *Stack) pop() any {
	if stack.len > 0 {
		data := stack.items[stack.len-1]
		stack.items = stack.items[:stack.len-1]
		stack.len--
		return data
	} else {
		return nil
	}
}

func (stack *Stack) lookTop() any {
	if stack.len > 0 {
		return stack.items[stack.len-1]
	} else {
		return nil
	}
}
