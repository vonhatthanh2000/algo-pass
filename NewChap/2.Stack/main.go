package main

// type MinStack struct {
// 	min   int
// 	stack []int
// }

// func Constructor() MinStack {
// 	return MinStack{
// 		min:   math.MaxInt64,
// 		stack: []int{},
// 	}
// }

// func (this *MinStack) Push(val int) {
// 	if len(this.stack) == 0 {
// 		this.stack = append(this.stack, 0)
// 		this.min = val
// 	} else {
// 		this.stack = append(this.stack, val-this.min)
// 		if val < this.min {
// 			this.min = val
// 		}
// 	}

// 	fmt.Println(this.stack)
// 	fmt.Println(this.min)
// 	fmt.Println("------")
// }

// func (this *MinStack) Pop() {
// 	if len(this.stack) == 0 {
// 		return
// 	}
// 	pop := this.stack[len(this.stack)-1]
// 	this.stack = this.stack[:len(this.stack)-1]
// 	if pop < 0 {
// 		this.min = this.min - pop
// 	}
// }

// func (this *MinStack) Top() int {
// 	top := this.stack[len(this.stack)-1]
// 	if top > 0 {
// 		return top + this.min
// 	}
// 	return this.min
// }

// func (this *MinStack) GetMin() int {
// 	return this.min
// }

// func main() {

// 	s := Constructor()

// 	s.Push(2)
// 	s.Push(3)
// 	s.Push(4)
// 	s.Push(1)
// 	s.Push(0)
// 	s.Push(5)

// }
