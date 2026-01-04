package main

import "fmt"

type SingleLinkedList[T comparable] struct {
	Head *Element[T]
	Tail *Element[T]
}

type Element[T comparable] struct {
	Value     T
	NextValue *Element[T]
}

func (s *SingleLinkedList[T]) Add(t T) {
	// Creates the element which it is going
	// to be added to the end of the linked list
	n := Element[T]{
		Value: t,
	}

	// There is no other values, so the head and
	// tail must point to the new element
	if s.Head == nil {
		s.Head = &n
		s.Tail = &n
		return
	}
	// If there is more values
	// Before to move the tail, you point to the next value
	// Then you can move the tail pointer to the new value
	s.Tail.NextValue = &n
	s.Tail = s.Tail.NextValue
}

func (s *SingleLinkedList[T]) Insert(t T, pos int) {
	n := Element[T]{
		Value: t,
	}

	// There is no other values, so the head and
	// tail must point to the new element
	if s.Head == nil {
		s.Head = &n
		s.Tail = &n
		return
	}

	// If position is 0 or below, new value
	// is added at the beginning, being now the Head
	if pos <= 0 {
		n.NextValue = s.Head
		s.Head = &n
		return
	}

	curNode := s.Head
	// Iteration over the values before the position
	for i := 1; i < pos; i++ {
		// If current node has no next value
		// Added new value as next value
		// And tail now points to that new value
		if curNode.NextValue == nil {
			curNode.NextValue = &n
			s.Tail = curNode.NextValue
			return
		}
		// If there is new value, move on
		curNode = curNode.NextValue
	}
	// We reached the position where we want to
	// introduce a new value
	// We move the value there making new value point to it
	// And then, the currentValue next value pointer
	// must point to the new value
	n.NextValue = curNode.NextValue
	curNode.NextValue = &n
	// If we are on the last value, tail must be updated
	if s.Tail == curNode {
		s.Tail = &n
	}
	return
}

func (s *SingleLinkedList[T]) Index(t T) int {
	i := 0
	// Iteration over values while node is not nil
	for curNode := s.Head; curNode != nil; curNode = curNode.NextValue {
		if curNode.Value == t {
			return i
		}
		i++
	}
	return -1
}

func main() {
	l := &SingleLinkedList[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))

	l.Insert(100, 0)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	l.Insert(200, 1)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(200))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	for curNode := l.Head; curNode != nil; curNode = curNode.NextValue {
		fmt.Print(curNode.Value)
		fmt.Print(" ")
	}
	fmt.Print("\n")

	l.Insert(300, 10)
	for curNode := l.Head; curNode != nil; curNode = curNode.NextValue {
		fmt.Print(curNode.Value)
		fmt.Print(" ")
	}
	fmt.Print("\n")

	l.Add(400)
	for curNode := l.Head; curNode != nil; curNode = curNode.NextValue {
		fmt.Print(curNode.Value)
		fmt.Print(" ")
	}
	fmt.Print("\n")

	l.Insert(500, 6)
	for curNode := l.Head; curNode != nil; curNode = curNode.NextValue {
		fmt.Print(curNode.Value)
		fmt.Print(" ")
	}
}
