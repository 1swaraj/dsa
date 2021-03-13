package main

import "fmt"

type Stack []int

var top int

func main() {
	stack := make(Stack, 5)
	stack.push(3)
	stack.push(5)
	stack.print()
	stack.push(7)
	stack.push(8)
	stack.push(9)
	err := stack.push(8)
	if err != nil {
		fmt.Println(err.Error())
	}
	stack.print()
	val,_ := stack.pop()
	fmt.Println(val)
	val,_ = stack.pop()
	fmt.Println(val)
	val,_ = stack.pop()
	fmt.Println(val)
	val,_ = stack.pop()
	fmt.Println(val)
	val,_ = stack.pop()
	fmt.Println(val)
	val,err = stack.pop()
	fmt.Println(err.Error())

	stack.push(9)
	stack.push(8)
	stack.print()
}

func (stack *Stack) push(val int) (err error) {
	if top >= 5 {
		err = fmt.Errorf("Overflows")
		return
	}
	st := *stack
	st[top] = val
	top++
	return
}

func (stack *Stack) pop() (popped int, err error) {
	if top <= 0 {
		err = fmt.Errorf("Underflows")
		return
	}
	st := *stack
	popped = st[top-1]
	top--
	return
}

func (stack *Stack) print() {
	st := *stack
	for i := 0; i < top; i++ {
		fmt.Print(st[i])
		fmt.Print(",")
	}
	fmt.Println()
}
