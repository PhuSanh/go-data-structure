package stack

import (
	"github.com/cheekybits/genny/generic"
	"sync"
)

// Item the type of the stack
type Item generic.Type

// ItemStack the stack of Items
type ItemStack struct {
	items []Item
	lock sync.RWMutex
}

// New creates a new ItemStack
func (s *ItemStack) New() *ItemStack {
	s.items = []Item{}
	return s
}

// Push adds an Item to the top of the stack
func (s *ItemStack) Push(t Item) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, t)
}

// Pop removes an Item from the top of the stack and return Item
func (s *ItemStack) Pop() *Item {
	s.lock.Lock()
	defer s.lock.Unlock()

	lastIndex := len(s.items) - 1
	t := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return &t
}

func (s *ItemStack) Empty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.items) == 0
}