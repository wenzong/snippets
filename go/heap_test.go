package main

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	assert.EqualValues(t, (*h)[0], 2)

	heap.Init(h)
	assert.EqualValues(t, (*h)[0], 1)

	heap.Push(h, 0)
	assert.EqualValues(t, (*h)[0], 0)

	(*h)[0] = 100
	heap.Fix(h, 0)
	assert.EqualValues(t, (*h)[0], 1)

	heap.Pop(h)
	assert.EqualValues(t, (*h)[0], 2)
}
