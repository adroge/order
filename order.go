/*

Package order provides a means to order (sort) simple and complex slices
based upon multiple ordering criteria, generically, and in-place.

*/

package order

import (
	"sort"
)

type compareFunc[T any] func(left, right T) int

// order tracks the criteria for ordering a list.
type order[T any] struct {
	list    []T
	compare []compareFunc[T]
}

// By returns a Sorter that can be used to perform sort related operations on
// the list.
//
//	list := []string{"b", "a"}
//	stringValue := func(l, r string) int {return order.Compare(l,r)}
//	order.By(stringValue).Sort(list)
func By[T any](compare ...compareFunc[T]) *order[T] {
	return &order[T]{
		compare: compare,
	}
}

// Sort sorts the argument slice according to the compare functions passed.
//
//	list := []string{"b", "a"}
//	stringValue := func(l, r string) int {return order.Compare(l,r)}
//	order.By(stringValue).Sort(list)
func (o *order[T]) Sort(listItems []T) {
	o.list = listItems
	sort.Sort(o)
}

// Stable does a stable sort.
//
//	list := []string{"b", "a"}
//	stringValue := func(l, r string) int {return order.Compare(l,r)}
//	order.By(stringValue).Stable(list)
func (o *order[T]) Stable(listItems []T) {
	o.list = listItems
	sort.Stable(o)
}

// IsSorted checks if a slice is sorted.
//
//	list := []string{"b", "a"}
//	stringValue := func(l, r string) int {return order.Compare(l,r)}
//	sorted := order.By(stringValue).IsSorted(list)
func (o *order[T]) IsSorted(listItems []T) bool {
	o.list = listItems
	return sort.IsSorted(o)
}

// Len is part of sort.Interface.
func (o *order[T]) Len() int {
	return len(o.list)
}

// Swap is part of sort.Interface.
func (o *order[T]) Swap(left, right int) {
	o.list[left], o.list[right] = o.list[right], o.list[left]
}

// Less is part of sort.Interface.
func (o *order[T]) Less(left, right int) bool {
	leftItem, rightItem := o.list[left], o.list[right]
	var compareIndex int
	for compareIndex = 0; compareIndex < len(o.compare)-1; compareIndex++ {
		switch comparison := o.compare[compareIndex](leftItem, rightItem); {
		case comparison < 0:
			return true
		case comparison > 0:
			return false
		}
	}
	return o.compare[compareIndex](leftItem, rightItem) < 0
}
