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
