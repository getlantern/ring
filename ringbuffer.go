// Package ringbuffer provides bounded circular buffers.
package ringbuffer

type RingBuffer interface {
	// Push() pushes a value to the head of the buffer
	Push(interface{})

	// Pop() pops the tail value off the buffer

	// IterateForward() iterates forward through the values starting at the tail.
	// Iteration stops if the callback function returns false.
	IterateForward(func(interface{}) bool)

	// IterateBackward() iterates backwards through the values starting at the
	// head. Iteration stops if the callback function returns false.
	IterateBackward(func(interface{}) bool)

	// Len() returns the number of items in the buffer
	Len() int
}

type ringBuffer struct {
	data []interface{}
	cap  int
	len  int
	head int
}

// New constructs a new RingBuffer bounded to cap. If cap <= 0, the RingBuffer
// will have a capacity of 1.
func New(cap int) RingBuffer {
	if cap <= 0 {
		// need at least 1
		cap = 1
	}
	return &ringBuffer{
		data: make([]interface{}, cap),
		cap:  cap,
		len:  0,
		head: -1,
	}
}

func (rb *ringBuffer) Push(val interface{}) {
	rb.head++
	if rb.head >= rb.cap {
		rb.head -= rb.cap
	}
	rb.data[rb.head] = val
	if rb.len < rb.cap {
		rb.len++
	}
}

func (rb *ringBuffer) IterateForward(cb func(interface{}) bool) {
	for i := rb.len - 1; i >= 0; i-- {
		idx := rb.head + i
		if idx >= rb.cap {
			// wrap around
			idx -= rb.cap
		}
		if !cb(rb.data[idx]) {
			return
		}
	}
}

func (rb *ringBuffer) IterateBackward(cb func(interface{}) bool) {
	if rb.empty() {
		return
	}

	for i := 0; i < rb.len; i++ {
		idx := rb.head + i
		if idx >= rb.cap {
			// wrap around
			idx -= rb.cap
		}
		if !cb(rb.data[idx]) {
			return
		}
	}
}

func (rb *ringBuffer) Len() int {
	return rb.len
}

func (rb *ringBuffer) empty() bool {
	return rb.len == 0
}
