package main

import "fmt"

// Буфер фиксированного размера
type MyBuffer struct {
	arr  []int
	head int
	tail int
	size int
	cap  int
}

func NewMyBuffer(capacity int) *MyBuffer {
	return &MyBuffer{
		arr:  make([]int, capacity),
		head: 0,
		tail: 0,
		size: 0,
		cap:  capacity,
	}
}

// Push добавляет элемент в конец.
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

// Pop удаляет самый старый элемент (с начала).
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

// Iterate проходит по элементам от старых к новым
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

// Peek ищет самый старый элемент без удаления
func (mb *MyBuffer) Peek() (int, bool) {
	if mb.Len() == 0 {
		return 0, false
	}
	index := mb.head % mb.Capacity()
	result := mb.arr[index]
	return result, true
}

// Clear очищает буфер
func (mb *MyBuffer) Clear() {
	mb.head = 0
	mb.tail = 0
	mb.size = 0
}

// Capacity возвращает максимальную емкость буфера
func (mb *MyBuffer) Capacity() int {
	return mb.cap
}

// Len возвращает количество элементов в буфере
func (mb *MyBuffer) Len() int {
	return mb.size
}

func main() {
	mb := NewMyBuffer(4)

	fmt.Println("Записываем 3 значения: 5, 23 и 41")
	mb.Push(5)
	mb.Push(23)
	mb.Push(41)

	fmt.Print("Текущие значения: ")
	mb.Iterate(func(i, v int) {
		fmt.Printf("%d ", v)
	})
	fmt.Printf("\n\n")

	fmt.Println("Удаляем самый старый элемент")
	val, _ := mb.Pop()
	fmt.Printf("Самый старый элемент: %d", val)
	fmt.Printf("\n\n")

	fmt.Println("Записываем 32, 54, 67")
	mb.Push(32)
	mb.Push(54)
	mb.Push(67)

	fmt.Print("Текущие значения: ")
	mb.Iterate(func(i, v int) {
		fmt.Printf("%d ", v)
	})
	fmt.Println()

	oldestElement, _ := mb.Peek()
	fmt.Printf("Самый старый элемент: %d\n", oldestElement)
}
