package heap

type heap interface {
	insert(value int)
	delete() int
}

type minHeap struct {
	Length int
	data   []int
}

func NewMinHeap() minHeap {
	return minHeap{Length: 0, data: make([]int, 0)}
}

func (m *minHeap) insert(value int) {
	m.data = append(m.data, value)
	m.heapifyUp(m.Length)
	m.Length++
}

func (m *minHeap) delete() int {
	if m.Length == 0 {
		return -1
	}
	out := m.data[0]

	m.Length--
	if m.Length == 0 {
		m.data = m.data[:0]
		return out
	}

	m.data[0] = m.data[m.Length]
	m.heapifyDown(0)
	return out
}

func (m *minHeap) heapifyDown(idx int) {
	lIdx := leftChild(idx)
	rIdx := rightChild(idx)

	if lIdx >= m.Length || idx >= m.Length {
		return
	}

	lValue := m.data[lIdx]
	rValue := m.data[rIdx]
	value := m.data[idx]

	if lValue > rValue && value > rValue {
		m.data[idx] = rValue
		m.data[rIdx] = value
		m.heapifyDown(rIdx)
	} else if rValue > lValue && value > lValue {
		m.data[idx] = lValue
		m.data[lIdx] = value
		m.heapifyDown(lIdx)
	}

}

func (m *minHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	p := parent(idx)
	parentV := m.data[p]
	v := m.data[idx]

	if parentV > v {
		m.data[idx] = parentV
		m.data[p] = v
		m.heapifyUp(p)
	}
}
func parent(idx int) int {
	return (idx - 1) / 2
}

func leftChild(idx int) int {
	return (idx * 2) + 1
}
func rightChild(idx int) int {
	return (idx * 2) + 2
}
