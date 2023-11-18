package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnqueue(t *testing.T) {
	c := require.New(t)

	queue := NewQueue()
	queue.Enqueue("test")
	queue.Enqueue("test 2")

	c.Equal("test", queue.Peek())
	c.Equal(2, queue.Size())
}

func TestDeQueue(t *testing.T) {
	c := require.New(t)

	queue := NewQueue()
	queue.Enqueue("test")

	item := queue.DeQueue()
	c.Equal("test", item)

	item = queue.DeQueue()
	c.Empty(item)
}
