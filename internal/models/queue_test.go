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
	queue.Enqueue("test2")

	item := queue.DeQueue()
	c.Equal("test", item)
	c.Equal(1, queue.Size())

	item = queue.DeQueue()
	c.Equal("test2", item)
	c.Equal(0, queue.Size())

	c.Empty(queue.DeQueue())
}
