package lru

type Node[t any] struct {
	value t
	next  *Node[t]
	prev  *Node[t]
}
type LruCache[k comparable, v comparable] struct {
	length        int
	capacity      int
	head          *Node[v]
	tail          *Node[v]
	lookup        map[k]*Node[v]
	reverseLookup map[*Node[v]]k
}

func NewLruCache[k comparable, v comparable](capacity int) LruCache[k, v] {
	return LruCache[k, v]{
		length:        0,
		capacity:      capacity,
		head:          nil,
		tail:          nil,
		lookup:        make(map[k]*Node[v]),
		reverseLookup: make(map[*Node[v]]k),
	}
}

func (l *LruCache[k, v]) Update(key k, value v) {
	// does exist?
	node := l.lookup[key]
	if node == nil {
		node = &Node[v]{value: value}
		l.length++
		l.trimCache()
		l.lookup[key] = node
		l.reverseLookup[node] = key
	} else {
		l.detach(node)
		l.prepend(node)
		node.value = value
	}
	// if not insert
	// if does, update to front of list and update value
}

func (l *LruCache[K, V]) Get(key K) *V {
	// check cache for existance
	node := l.lookup[key]
	if node == nil {
		return nil
	}
	// update value we found and move to front
	l.detach(node)
	l.prepend(node)

	// return value found or nil if not exist
	return &node.value
}

func (l *LruCache[k, v]) detach(node *Node[v]) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if l.head == node {
		l.head = l.head.next
	}

	if l.tail == node {
		l.tail = l.tail.prev
	}

	node.next = nil
	node.prev = nil
}

func (l *LruCache[k, v]) prepend(node *Node[v]) {
	if l.head == nil {
		l.head, l.tail = node, node
		return
	}

	node.next = l.head
	l.head.prev = node

	l.head = node
}

func (l *LruCache[k, v]) trimCache() {
	if l.length <= l.capacity {
		return
	}

	tail := l.tail
	l.detach(l.tail)

	key := l.reverseLookup[tail]
	delete(l.lookup, key)
	delete(l.reverseLookup, tail)
	l.length--
}
