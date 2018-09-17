package lru

type node struct {
	Prev  *node
	Next  *node
	Key   string
	Value int
}

type list struct {
	MaxSize int
	CurSize int
	Head    *node
	Tail    *node
}

type cache struct {
	c    map[string]*node
	size int
	l    *list
}

func NewCache(size int) cache {
	l := list{size, 0, nil, nil}
	return cache{map[string]*node{}, 0, &l}
}

func (C *cache) set(key string, value int) {
	var n *node
	if C.c[key] == nil {
		n = &node{
			Prev:  nil,
			Next:  nil,
			Key:   key,
			Value: value,
		}
		C.c[key] = n
	} else {
		n = C.c[key]
		n.Value = value
	}

	if dt := C.l.push(n); dt != nil {
		C.c[dt.Key] = nil
	}
	return
}

func (C *cache) get(key string) int {
	n := C.c[key]
	if n == nil {
		return -1
	}
	if n != C.l.Head {
		C.l.removeNode(n)
		C.l.push(n)
	}
	return n.Value
}

func (l *list) push(n *node) (dt *node) {
	if l.CurSize >= l.MaxSize {
		dt = l.dropTail()
	}
	l.appendHead(n)
	return
}

func (l *list) dropTail() (n *node) {
	if l.CurSize == 0 {
		return
	}
	l.CurSize -= 1
	n = l.Tail

	if l.Tail.Prev != nil {
		l.Tail.Prev.Next = nil
		l.Tail = l.Tail.Prev
		return
	}
	l.Tail = nil
	return
}

func (l *list) appendHead(n *node) {
	if l.CurSize == l.MaxSize {
		return
	}
	if l.CurSize == 0 {
		l.CurSize = 1
		l.Head = n
		l.Tail = n
		return
	}
	l.CurSize += 1
	n.Prev = nil
	n.Next = l.Head
	l.Head.Prev = n
	l.Head = n
	return
}

func (l *list) removeNode(n *node) {
	if l.CurSize == 0 {
		return
	}

	if n.Next == nil && n.Prev == nil {
		l.Head = nil
		l.Tail = nil
		l.CurSize = 0
		return
	}

	if n.Next != nil {
		n.Next.Prev = n.Prev
	} else {
		n.Prev.Next = nil
		l.Tail = n.Prev
	}

	if n.Prev != nil {
		n.Prev.Next = n.Next
	} else {
		n.Next.Prev = nil
		l.Head = n.Next
	}
	l.CurSize -= 1
	return
}
