package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem895(t *testing.T) {
	obj := Constructor_895()
	fmt.Printf("obj = %v\n", obj)
	obj.Push(5)
	fmt.Printf("obj = %v\n", obj)
	obj.Push(7)
	fmt.Printf("obj = %v\n", obj)
	obj.Push(5)
	fmt.Printf("obj = %v\n", obj)
	obj.Push(7)
	fmt.Printf("obj = %v\n", obj)
	obj.Push(4)
	fmt.Printf("obj = %v\n", obj)
	obj.Push(5)
	fmt.Printf("obj = %v\n", obj)
	param_1 := obj.Pop()
	fmt.Printf("param_1 = %v\n", param_1)
	param_1 = obj.Pop()
	fmt.Printf("param_1 = %v\n", param_1)
	param_1 = obj.Pop()
	fmt.Printf("param_1 = %v\n", param_1)
	param_1 = obj.Pop()
	fmt.Printf("param_1 = %v\n", param_1)
	param_1 = obj.Pop()
	fmt.Printf("param_1 = %v\n", param_1)
	param_1 = obj.Pop()
	fmt.Printf("param_1 = %v\n", param_1)
}
