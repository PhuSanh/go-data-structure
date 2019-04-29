package queue

import (
	"sync"
)

// ItemQueue the queue of Items
type ItemQueueCustom struct {
	items []interface{}
	lock sync.RWMutex
}

func (s *ItemQueueCustom) New() interface{} {
	var empty []interface{}
	s.items = empty
	return s
}

// New creates a new ItemQueue
func (s *ItemQueueCustom) Enqueue(t Item) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = append(s.items, t)
}

// Dequeue removes an Item from the start of the queue
func (s *ItemQueueCustom) Dequeue() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	t := s.items[0]
	s.items = s.items[1:]
	return t
}

// Front returns the item next in the queue, without removing it
func (s *ItemQueueCustom) Front() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	t := s.items[len(s.items) - 1]
	return t
}

// IsEmpty returns true if the queue is empty
func (s *ItemQueueCustom) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of Items in the queue
func (s *ItemQueueCustom) Size() int {
	return len(s.items)
}