package treemap

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/iterator"
)

// MapIterator is an iterator for Map
type MapIterator struct {
	node *rbtree.Node
}

// IsValid returns whether iter is valid
func (iter *MapIterator) IsValid() bool {
	if iter.node != nil {
		return true
	}
	return false
}

// Next returns the next iterator
func (iter *MapIterator) Next() iterator.ConstIterator {
	if iter.IsValid() {
		iter.node = iter.node.Next()
	}
	return iter
}

// Prev returns the previous iterator
func (iter *MapIterator) Prev() iterator.ConstBidIterator {
	if iter.IsValid() {
		iter.node = iter.node.Prev()
	}
	return iter
}

// Key returns the key of iter
func (iter *MapIterator) Key() interface{} {
	return iter.node.Key()
}

// Value returns the value of iter
func (iter *MapIterator) Value() interface{} {
	return iter.node.Value()
}

// SetValue sets the value of iter
func (iter *MapIterator) SetValue(val interface{}) error {
	iter.node.SetValue(val)
	return nil
}

// Clone clones iter to a new MapIterator
func (iter *MapIterator) Clone() iterator.ConstIterator {
	return &MapIterator{iter.node}
}

// Equal returns whether iter is equal to other
func (iter *MapIterator) Equal(other iterator.ConstIterator) bool {
	otherIter, ok := other.(*MapIterator)
	if !ok {
		return false
	}
	if otherIter.node == iter.node {
		return true
	}
	return false
}
