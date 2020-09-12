package chapter1

import (
	"fmt"
	"testing"
)

func TestFixedCapacityStackOfStrings(t *testing.T) {
	stack := New(10)
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Size())
	fmt.Println()
	stack.Push("1")
	stack.Push("2")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Size())
	fmt.Println()
	stack.Push("3")
	stack.Push("4")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Size())
	fmt.Println()
	stack.Push("5")
	stack.Push("6")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Size())
	fmt.Println()
	stack.Push("7")
	fmt.Println(stack.Pop())
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Size())
}
