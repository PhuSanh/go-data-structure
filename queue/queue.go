package queue

import (
	"sync"
)

// Item the type of the queue
type Item interface {}

// ItemQueue the queue of Items
type ItemQueue struct {
	items []Item
	lock sync.RWMutex
}

func (s *ItemQueue) New() *ItemQueue {
	s.items = []Item{}
	return s
}

// New creates a new ItemQueue
func (s *ItemQueue) Enqueue(t Item) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = append(s.items, t)
}

// Dequeue removes an Item from the start of the queue
func (s *ItemQueue) Dequeue() *Item {
	s.lock.Lock()
	defer s.lock.Unlock()

	t := s.items[0]
	s.items = s.items[1:]
	return &t
}

// Front returns the item next in the queue, without removing it
func (s *ItemQueue) Front() *Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	t := s.items[len(s.items) - 1]
	return &t
}

// IsEmpty returns true if the queue is empty
func (s *ItemQueue) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of Items in the queue
func (s *ItemQueue) Size() int {
	return len(s.items)
}