package main

// type Stack struct {
// 	items []interface{}
// }

// func (s *Stack) Push(item interface{}) {
// 	s.items = append(s.items, item)
// }

// func (s *Stack) Pop() (interface{}, bool) {
// 	if s.IsEmpty() {
// 		return nil, false
// 	}
// 	index := len(s.items) - 1
// 	item := s.items[index]
// 	s.items = s.items[:index]
// 	return item, true
// }

// func (s *Stack) Peek() (interface{}, bool) {
// 	if s.IsEmpty() {
// 		return nil, false
// 	}
// 	return s.items[len(s.items)-1], true
// }

// func (s *Stack) IsEmpty() bool {
// 	return len(s.items) == 0
// }

// func (s *Stack) Size() int {
// 	return len(s.items)
// }

// func NewStack() *Stack {
// 	return &Stack{}
// }

// func isValid(s string) bool {
// 	stack := Stack{}

// 	closeToOpen := map[rune]rune{'}': '{', ')': '(', ']': '['}

// 	for _, c := range s {
// 		if open, exist := closeToOpen[c]; exist {
// 			if !stack.IsEmpty() {
// 				a, ok := stack.Pop()
// 				if ok && a != open {
// 					return false
// 				}
// 			} else {
// 				return false
// 			}
// 		} else {
// 			stack.Push(c)
// 		}
// 	}

// 	return stack.IsEmpty()
// }

// ////

// type MinStack struct {
// 	stack    *Stack
// 	minStack *Stack
// }

// func Constructor() MinStack {
// 	return MinStack{
// 		stack:    NewStack(),
// 		minStack: NewStack(),
// 	}
// }

// func (this *MinStack) Push(val int) {
// 	this.stack.items = append(this.stack.items, val)

// 	min := val
// 	if !this.minStack.IsEmpty() {
// 		top := this.GetMin()
// 		if top < min {
// 			min = top
// 		}
// 	}
// 	this.minStack.Push(min)
// }

// func (this *MinStack) Pop() {
// 	this.stack.Pop()
// 	this.minStack.Pop()
// }

// func (this *MinStack) Top() int {
// 	top, _ := this.stack.Peek()
// 	return top.(int)
// }

// func (this *MinStack) GetMin() int {
// 	top, _ := this.minStack.Peek()
// 	return top.(int)
// }

// func main() {
// 	minStack := Constructor()
// 	minStack.Push(-2)

// 	minStack.Push(0)
// 	minStack.Push(-3)

// 	minStack.Pop()
// }
