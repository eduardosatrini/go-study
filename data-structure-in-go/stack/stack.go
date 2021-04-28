package stack

type Stack struct {
	values []string
}

func (s *Stack) Push(name string) int {

	s.values = append(s.values, name)

	return len(s.values)
}

func (s *Stack) Pop() int {

	if len(s.values) != 0 {
		s.values = s.values[:len(s.values)-1]
	}

	return len(s.values)
}
