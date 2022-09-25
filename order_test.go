package order_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/adroge/order"
)

type stInFl struct {
	st string
	in int
	fl float64
}

func TestStringSort(t *testing.T) {
	list := []string{"g", "f", "d", "e", "c", "b", "a"}
	ascending := func(l, r string) int {
		if l < r {
			return -1
		}
		if l > r {
			return 1
		}
		return 0
	}
	order.By(ascending).Sort(list)
	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f", "g"}, list)
}

func TestPointerStringSort(t *testing.T) {
	a, b, c := "a", "b", "c"
	list := []*string{&c, &a, &b}
	ascending := func(l, r *string) int {
		if *l < *r {
			return -1
		}
		if *l > *r {
			return 1
		}
		return 0
	}
	order.By(ascending).Sort(list)
	assert.Equal(t, []*string{&a, &b, &c}, list)
}

func TestIntegerSort(t *testing.T) {
	list := []int{9, 8, 7, 5, 6, 4, 3, 2, 1, 0}
	ascending := func(l, r int) int {
		if l < r {
			return -1
		}
		if l > r {
			return 1
		}
		return 0
	}
	order.By(ascending).Sort(list)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, list)
}

func TestFloatSort(t *testing.T) {
	list := []float64{9, 8, 7, 5, 6, 4, 3, 2, 1, 0}
	ascending := func(l, r float64) int {
		switch {
		case l < r:
			return -1
		case l > r:
			return 1
		}
		return 0
	}
	order.By(ascending).Sort(list)
	assert.Equal(t, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, list)
}

func TestUintIsSorted(t *testing.T) {
	list := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ascending := func(l, r uint) int {
		if l < r {
			return -1
		}
		if l > r {
			return 1
		}
		return 0
	}
	assert.True(t, order.By(ascending).IsSorted(list))
}

func TestUintIsNotSorted(t *testing.T) {
	list := []uint{9, 0, 1, 2, 3, 4, 5, 6, 7, 8}
	ascending := func(l, r uint) int {
		if l < r {
			return -1
		}
		if l > r {
			return 1
		}
		return 0
	}
	assert.False(t, order.By(ascending).IsSorted(list))
}

func TestSortStructureBySingle(t *testing.T) {
	list := []stInFl{
		{"d", 2, 1},
		{"b", 2, 2},
		{"a", 1, 1},
		{"c", 1, 2},
	}
	expected := []stInFl{
		{"a", 1, 1},
		{"b", 2, 2},
		{"c", 1, 2},
		{"d", 2, 1},
	}
	st := func(left, right stInFl) int {
		return order.Compare(left.st, right.st)
	}
	order.By(st).Sort(list)
	assert.Equal(t, expected, list)
}

func TestSortStructureByTwo(t *testing.T) {
	list := []stInFl{
		{"b", 2, 1},
		{"a", 2, 2},
		{"a", 1, 1},
		{"b", 1, 2},
	}
	expected := []stInFl{
		{"a", 1, 1},
		{"a", 2, 2},
		{"b", 1, 2},
		{"b", 2, 1},
	}
	st := func(left, right stInFl) int {
		return order.Compare(left.st, right.st)
	}
	in := func(left, right stInFl) int {
		return order.Compare(left.in, right.in)
	}
	order.By(st, in).Sort(list)
	assert.Equal(t, expected, list)
}

func TestSortStructureByThree(t *testing.T) {
	list := []stInFl{
		{"b", 2, 1},
		{"b", 1, 1},
		{"a", 2, 2},
		{"c", 1, 2},
		{"c", 2, 1},
		{"a", 1, 2},
		{"a", 1, 1},
		{"b", 1, 2},
		{"c", 1, 1},
		{"c", 2, 2},
		{"b", 1, 2},
	}
	expected := []stInFl{
		{"a", 1, 1},
		{"b", 1, 1},
		{"b", 2, 1},
		{"c", 1, 1},
		{"c", 2, 1},
		{"a", 1, 2},
		{"a", 2, 2},
		{"b", 1, 2},
		{"b", 1, 2},
		{"c", 1, 2},
		{"c", 2, 2},
	}
	st := func(left, right stInFl) int {
		return order.Compare(left.st, right.st)
	}
	in := func(left, right stInFl) int {
		return order.Compare(left.in, right.in)
	}
	fl := func(left, right stInFl) int {
		return order.Compare(left.fl, right.fl)
	}
	order.By(fl, st, in).Sort(list)
	assert.Equal(t, expected, list)
}

func TestStableSortStructureBySingle(t *testing.T) {
	stableList := []stInFl{
		{"a2", 2, 2},
		{"a1", 1, 1},
		{"b2", 2, 2},
		{"b1", 1, 1},
		{"c1", 1, 1},
		{"c2", 2, 2},
		{"d1", 1, 1},
		{"d2", 2, 2},
		{"e1", 1, 1},
		{"f1", 1, 1},
		{"g1", 1, 1},
		{"e2", 2, 2},
		{"f2", 2, 2},
		{"g2", 2, 2},
	}
	unstableList := make([]stInFl, len(stableList))
	copy(unstableList, stableList)
	expected := []stInFl{
		{"a1", 1, 1},
		{"b1", 1, 1},
		{"c1", 1, 1},
		{"d1", 1, 1},
		{"e1", 1, 1},
		{"f1", 1, 1},
		{"g1", 1, 1},
		{"a2", 2, 2},
		{"b2", 2, 2},
		{"c2", 2, 2},
		{"d2", 2, 2},
		{"e2", 2, 2},
		{"f2", 2, 2},
		{"g2", 2, 2},
	}
	in := func(left, right stInFl) int {
		return order.Compare(left.in, right.in)
	}
	order.By(in).Stable(stableList)
	assert.Equal(t, expected, stableList)
	order.By(in).Sort(unstableList)
	assert.NotEqual(t, expected, unstableList)
}

func TestIsSorted(t *testing.T) {
	list := []int{9, 8, 7}
	descending := func(left, right int) int {
		return -1 * order.Compare(left, right)
	}
	assert.True(t, order.By(descending).IsSorted(list))
}

func TestIsNotSorted(t *testing.T) {
	list := []int{9, 8, 9}
	descending := func(left, right int) int {
		return -1 * order.Compare(left, right)
	}
	assert.False(t, order.By(descending).IsSorted(list))
}
