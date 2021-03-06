package main

import (
	"errors"
)

//	1.将栈定义为一个借口方便插入任意类型数据
//	2.采用切片可以动态分配大小。还有Len() 和Cap()这个2个战术可以免费使用
type Stack []interface{}

//	返回栈现在元素个数
func (stack Stack) Len() int {
	return len(stack)
}

//	返回栈的容量
func (stack Stack) Cap() int {
	return cap(stack)
}

//	判断栈是否为空
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

//	元素入栈
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

//	返回栈顶元素，如果栈为空则元素为空error不为空
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

//	元素出栈，如果栈为空则元素为空error不为空
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}
