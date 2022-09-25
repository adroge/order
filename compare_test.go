package order_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/adroge/order"
)

func TestCompareUint64Equal(t *testing.T) {
	var l, r uint64 = 1, 1
	assert.Zero(t, order.Compare(l, r))
}

func TestCompareUint64Less(t *testing.T) {
	var l, r uint64 = 0, 1
	assert.Equal(t, -1, order.Compare(l, r))
}

func TestCompareUint64Greater(t *testing.T) {
	var l, r uint64 = 1, 0
	assert.Equal(t, 1, order.Compare(l, r))
}

func TestCompareByteEqual(t *testing.T) {
	var l, r byte = 1, 1
	assert.Zero(t, order.Compare(l, r))
}

func TestCompareCustomStringTypeEqual(t *testing.T) {
	type aString string
	var a, b aString = "a", "a"
	assert.Zero(t, order.Compare(a, b))
}

func TestCompareCustomStringTypeLess(t *testing.T) {
	type aString string
	var a, b aString = "a", "b"
	assert.Equal(t, -1, order.Compare(a, b))
}

func TestCompareCustomStringTypeGreater(t *testing.T) {
	type aString string
	var a, b aString = "b", "a"
	assert.Equal(t, 1, order.Compare(a, b))
}

func TestComparePointerStringLess(t *testing.T) {
	a, b := "a", "b"
	left, right := &a, &b
	assert.Equal(t, -1, order.ComparePv(left, right))
}

func TestComparePointerStringEqual(t *testing.T) {
	a, b := "a", "a"
	left, right := &a, &b
	assert.Zero(t, order.ComparePv(left, right))
}

func TestComparePointerStringGreater(t *testing.T) {
	a, b := "b", "a"
	left, right := &a, &b
	assert.Equal(t, 1, order.ComparePv(left, right))
}
