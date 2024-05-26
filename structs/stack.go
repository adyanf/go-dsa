package structs

type Stack struct {
	stackList []int
}

func NewStack() *Stack {
	return &Stack{
		stackList: []int{},
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.stackList) == 0
}

func (s *Stack) Size() int {
	return len(s.stackList)
}

func (s *Stack) Top() int {
	if len(s.stackList) == 0 {
		return -1
	}

	return s.stackList[len(s.stackList)-1]
}

func (s *Stack) Push(element int) {
	s.stackList = append(s.stackList, element)
}

func (s *Stack) Pop() int {
	if len(s.stackList) == 0 {
		return -1
	}
	pop := s.stackList[len(s.stackList)-1]
	s.stackList = s.stackList[:len(s.stackList)-1]

	return pop
}
