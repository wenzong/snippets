package main

import "github.com/pkg/errors"

type Dequeue struct {
	array []int
	size  int
	phead int
	ptail int
}

func NewDequeue(size int) *Dequeue {
	return &Dequeue{
		array: make([]int, size+1),
		size:  size,
	}
}

func (q *Dequeue) IsFull() bool  { return (q.ptail+1)%(q.size+1) == q.phead }
func (q *Dequeue) IsEmpty() bool { return q.phead == q.ptail }

// head/pushhead/pophead
func (q *Dequeue) Head() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Get head failed: Dequeue is empty")
	}

	return q.array[q.phead], nil
}

func (q *Dequeue) PushHead(val int) error {
	if q.IsFull() {
		return errors.New("Push head failed: Dequeue is full")
	}

	q.phead--
	if q.phead < 0 {
		q.phead = q.size
	}
	q.array[q.phead] = val

	return nil
}

func (q *Dequeue) PopHead() (int, error) {
	item, err := q.Head()
	if err != nil {
		return 0, errors.Wrapf(err, "Pop head failed: Dequeue is empty")
	}

	q.phead = (q.phead + 1) % (q.size + 1)

	return item, nil
}

// tail/pushtail/poptail
func (q *Dequeue) Tail() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Get tail failed: Dequeue is empty")
	}

	if q.ptail == 0 {
		return q.array[q.size], nil
	}
	return q.array[q.ptail-1], nil
}

func (q *Dequeue) PushTail(val int) error {
	if q.IsFull() {
		return errors.New("Push tail failed: Dequeue is full")
	}

	q.array[q.ptail] = val
	q.ptail = (q.ptail + 1) % (q.size + 1)

	return nil
}

func (q *Dequeue) PopTail() (int, error) {
	item, err := q.Tail()
	if err != nil {
		return 0, errors.Wrapf(err, "Pop tail failed: Dequeue is empty")
	}

	q.ptail--
	if q.ptail < 0 {
		q.ptail = q.size
	}

	return item, nil
}
