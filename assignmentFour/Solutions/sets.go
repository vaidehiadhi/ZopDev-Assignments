package Solutions

type Set struct {
	elements map[string]struct{}
}

// creating empty set
func NewSet() *Set {
	return &Set{elements: make(map[string]struct{})}
}

// adding element to set
func (s *Set) Add(element string) {
	s.elements[element] = struct{}{}
}

// deleting element from set
func (s *Set) Delete(element string) {
	delete(s.elements, element)
}
