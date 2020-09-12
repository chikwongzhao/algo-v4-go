package chapter1

type FixedCapacityStackOfStrings struct {
	arr []string
	n   int
}

func New(n int) *FixedCapacityStackOfStrings {
	stack := &FixedCapacityStackOfStrings{}
	stack.arr = make([]string, n)
	stack.n = 0
	return stack
}

func (stack *FixedCapacityStackOfStrings) IsEmpty() bool {
	if stack == nil {
		return true
	}
	return stack.n == 0
}

func (stack *FixedCapacityStackOfStrings) Size() int {
	if stack == nil {
		return 0
	}
	return stack.n
}

func (stack *FixedCapacityStackOfStrings) Push(s string) {
	stack.arr[stack.n] = s
	stack.n++
}

func (stack *FixedCapacityStackOfStrings) Pop() string {
	if stack == nil || stack.IsEmpty() {
		return ""
	}
	s := stack.arr[stack.n-1]
	stack.n--
	return s
}
