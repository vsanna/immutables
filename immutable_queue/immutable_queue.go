package immutable_queue

// ImmutableQueue provides immutable queue.
type ImmutableQueue struct {
	items []interface{}
}

// New returns a pointer of ImmutableQueue
func New(items []interface{}) *ImmutableQueue {
	newItems := []interface{}{}

	for _, item := range items {
		newItems = append(newItems, item)
	}

	return &ImmutableQueue{items: newItems}
}

// IsEmpty checks items is blank or not
func (iq *ImmutableQueue) IsEmpty() bool {
	return len(iq.items) == 0
}

// Head returns head of items
func (iq *ImmutableQueue) Head() interface{} {
	if iq.IsEmpty() {
		return nil
	}

	return iq.items[0]
}

// Tail returns tail of items
func (iq *ImmutableQueue) Tail() interface{} {
	if iq.IsEmpty() {
		return nil
	}
	return iq.items[len(iq.items)-1]
}

// EnQueue create new ImmutableQueue with current items and new item
func (iq *ImmutableQueue) EnQueue(item interface{}) *ImmutableQueue {
	newItems := append(iq.items, item)
	return New(newItems)
}

// DeQueue create new ImmutableQueue without current items' head
func (iq *ImmutableQueue) DeQueue() *ImmutableQueue {
	if iq.IsEmpty() {
		return New([]interface{}{})
	}
	newIq := New(iq.items[1:])
	return newIq
}
