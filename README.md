# Circular buffer

This project implements the ring buffer data structure. This structure uses a single fixed-size buffer. Buffer has no real end and it can loop around the buffer.

## Main features

### Push method

Writes the value to the buffer. If the buffer is already full, the new value should overwrite the oldest value. 
The complexity of the method is O(1).

Realization:
```
func (mb *MyBuffer) Push(val int) {
	index := mb.tail % mb.Capacity()
	if mb.Len() == mb.Capacity() {
		mb.head++
	}
	mb.arr[index] = val
	mb.tail++
	if mb.Len() < mb.Capacity() {
		mb.size++
	}
}
```
### Pop method

Clears the oldest buffer element and returns true if there was a value, or zero-value and false if there is no value in the buffer to be cleaned. 
The complexity of the method is O(1).

Realization:
```
func (mb *MyBuffer) Pop() (int, bool) {
	if mb.Len() == 0 {
		return 0, false
	}
	index := mb.head % mb.Capacity()
	result := mb.arr[index]
	mb.head++
	mb.size--
	return result, true
}
```

### Iterate method

Passes through all the buffer elements from the oldest to the newest and applies the passed callback function to each one.
The complexity of the method is O(n).

Realization:
```
func (mb *MyBuffer) Iterate(callback func(int, int)) {
	if mb.Len() == 0 {
		return
	}
	for i := 0; i <= mb.Len()-1; i++ {
		index := (mb.head + i) % mb.Capacity()
		val := mb.arr[index]
		callback(i, val)
	}
}
```

The methods are also implemented:

Peek() (int, bool) - returns the oldest element without deleting it.

Clear() - clears the buffer.

