package ts

import (
	"sync"
)

// LinkedList is a thread safe first in first out (FIFO) list.
type LinkedList struct {
	head  *LinkedListItem
	tail  *LinkedListItem
	mutex sync.Mutex
}

// Push pushes a new item to the end of the LinkedList.
func (l *LinkedList) Push(value interface{}) {
	item := &LinkedListItem{
		Value: value,
		next:  nil,
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.tail != nil {
		l.tail.next = item
	}

	l.tail = item
	if l.head == nil {
		l.head = l.tail
	}
}

// Pop pops the first item from the start of the LinkedList.
func (l *LinkedList) Pop() interface{} {
	if l.head == nil {
		return nil
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	value := l.head.Value
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	}

	return value
}

// IsEmpty checks if the LinkedList is empty.
func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

// GetTail returns a copy of the value in the last item of the LinkedList.
func (l *LinkedList) GetTail() interface{} {
	if l.tail == nil {
		return nil
	}

	return l.tail.Value
}
