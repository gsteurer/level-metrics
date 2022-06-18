package containers

// we can't do generics on type aliases in go 1.18: https://github.com/golang/go/issues/46477
type Set[T comparable] struct {
	set map[T]struct{}
}

func CreateSet[T comparable]() Set[T] {
	return Set[T]{set: map[T]struct{}{}}
}

func (s Set[T]) Contains(item T) bool {
	_, ok := s.set[item]
	return ok
}

func (s Set[T]) Insert(item T) bool {
	var ok bool
	if _, ok = s.set[item]; !ok {
		s.set[item] = struct{}{}

	}
	return !ok
}

func (s Set[T]) Remove(item T) bool {
	var ok bool
	if _, ok = s.set[item]; ok {
		delete(s.set, item)
	}

	return ok
}
