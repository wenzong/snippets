package main

import (
	"testing"
	"time"

	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
)

func TestApplyFunc(t *testing.T) {
	patches := gomonkey.ApplyFunc(time.Now, func() time.Time {
		return time.Date(2001, time.January, 1, 12, 30, 0, 0, time.UTC)
	})

	defer patches.Reset()

	assert.EqualValues(t, time.Now().Year(), 2001)
}
