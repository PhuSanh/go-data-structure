package queue

import (
	"github.com/cheekybits/genny/generic"
	"sync"
)

// Item the type of the queue
type Item generic.Type

// ItemQueue the queue of Items
type ItemQueue struct {
	items []Item
	lock sync.RWMutex
}

