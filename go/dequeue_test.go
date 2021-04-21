package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDequeueEmptyAndFull(t *testing.T) {
	q := NewDequeue(0)
	assert.True(t, q.IsEmpty())
	assert.True(t, q.IsFull())

	assert.Error(t, q.PushTail(1))
	assert.Error(t, q.PushHead(1))
}

func TestDequeue(t *testing.T) {
	q := NewDequeue(10)

	// []int{100}
	q.PushTail(100)
	head, err := q.Head()
	assert.EqualValues(t, head, 100)
	assert.NoError(t, err)
	tail, err := q.Tail()
	assert.EqualValues(t, tail, 100)
	assert.NoError(t, err)

	// []int{10, 100}
	q.PushHead(10)
	head, _ = q.Head()
	tail, _ = q.Tail()
	assert.EqualValues(t, head, 10)
	assert.EqualValues(t, tail, 100)

	// []int{10}
	q.PopTail()
	head, _ = q.Head()
	tail, err = q.Tail()
	assert.EqualValues(t, head, 10)
	assert.EqualValues(t, tail, 10)
	assert.NoError(t, err)

	// []int{}
	q.PopTail()
	_, err = q.Head()
	assert.Error(t, err)
}
