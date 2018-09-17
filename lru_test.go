package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApendHead(t *testing.T) {
	n1, n2, n3, n4 := &node{
		Prev:  nil,
		Next:  nil,
		Key:   "A",
		Value: 1,
	}, &node{
		Prev:  nil,
		Next:  nil,
		Key:   "B",
		Value: 2,
	}, &node{
		Prev:  nil,
		Next:  nil,
		Key:   "C",
		Value: 3,
	}, &node{
		Prev:  nil,
		Next:  nil,
		Key:   "D",
		Value: 4,
	}

	var cases = []struct {
		n       *node
		expHead *node
		expTail *node
	}{
		{
			n:       n1,
			expHead: n1,
			expTail: n1,
		},
		{
			n:       n2,
			expHead: n2,
			expTail: n1,
		},
		{
			n:       n3,
			expHead: n3,
			expTail: n1,
		},
		{
			n:       n4,
			expHead: n3,
			expTail: n1,
		},
	}

	l := list{
		MaxSize: 3,
		CurSize: 0,
		Head:    nil,
		Tail:    nil,
	}

	for _, tt := range cases {
		l.appendHead(tt.n)
		assert.Equal(t, l.Head, tt.expHead, "The node was not propely appended")
		assert.Equal(t, l.Tail, tt.expTail, "The node was not propely appended")
	}
}

func TestDropTail(t *testing.T) {
	n1, n2, n3 := &node{
		Prev:  nil,
		Next:  nil,
		Key:   "A",
		Value: 1,
	}, &node{
		Prev:  nil,
		Next:  nil,
		Key:   "B",
		Value: 2,
	}, &node{
		Prev:  nil,
		Next:  nil,
		Key:   "C",
		Value: 3,
	}

	var emptyNode *node

	l := list{
		MaxSize: 3,
		CurSize: 0,
		Head:    nil,
		Tail:    nil,
	}

	var cases = []struct {
		expDrop *node
		expTail *node
	}{
		{
			expDrop: n1,
			expTail: n2,
		},
		{
			expDrop: n2,
			expTail: n3,
		},
		{
			expDrop: n3,
			expTail: emptyNode,
		},
	}
	l.appendHead(n1)
	l.appendHead(n2)
	l.appendHead(n3)

	for _, tt := range cases {
		dt := l.dropTail()
		assert.Equal(t, dt, tt.expDrop, "Droped the wrong node")
		assert.Equal(t, l.Tail, tt.expTail, "The new tail is wrongly assigned")

	}
}

func TestRemove(t *testing.T) {
	n1, n2, n3 := &node{
		Prev:  nil,
		Next:  nil,
		Key:   "A",
		Value: 1,
	}, &node{
		Prev:  nil,
		Next:  nil,
		Key:   "B",
		Value: 2,
	}, &node{
		Prev:  nil,
		Next:  nil,
		Key:   "C",
		Value: 3,
	}

	var emptyNode *node

	l := list{
		MaxSize: 3,
		CurSize: 0,
		Head:    nil,
		Tail:    nil,
	}

	l.appendHead(n1)
	l.appendHead(n2)
	l.appendHead(n3)

	var cases = []struct {
		remove  *node
		expHead *node
		expTail *node
		expSize int
	}{
		{
			remove:  n2,
			expHead: n3,
			expTail: n1,
			expSize: 2,
		},
		{
			remove:  n1,
			expHead: n3,
			expTail: n3,
			expSize: 1,
		},
		{
			remove:  n3,
			expHead: emptyNode,
			expTail: emptyNode,
			expSize: 0,
		},
		{
			remove:  n3,
			expHead: emptyNode,
			expTail: emptyNode,
			expSize: 0,
		},
	}

	for _, tt := range cases {
		l.removeNode(tt.remove)
		assert.Equal(t, l.Head, tt.expHead, "The node was not propely removed, wrong Head")
		assert.Equal(t, l.Tail, tt.expTail, "The node was not propely removed, wrong Tail")
		assert.Equal(t, l.CurSize, tt.expSize, "The node was not propely removed, wrong Size")
	}
}

func TestCache(t *testing.T) {
	c := NewCache(3)

	c.set("A", 1)
	c.set("B", 2)
	c.set("C", 3)
	c.set("D", 4)
	c.set("F", 5)

	assert.Equal(t, c.l.Head.Value, 5, "wrong Head")
	assert.Equal(t, c.l.Tail.Value, 3, "wrong Tail")

	assert.Equal(t, c.get("B"), -1, "Removed node found")
	assert.Equal(t, c.l.Head.Value, 5, " wrong Head")
	assert.Equal(t, c.l.Tail.Value, 3, "wrong Tail")

	assert.Equal(t, c.get("C"), 3, "wrong value for key")
	assert.Equal(t, c.l.Head.Value, 3, " wrong Head")
	assert.Equal(t, c.l.Tail.Value, 4, "wrong Tail")

	assert.Equal(t, c.get("F"), 5, "wrong value for key")
	assert.Equal(t, c.l.Head.Value, 5, " wrong Head")
	assert.Equal(t, c.l.Tail.Value, 4, "wrong Tail")
}
