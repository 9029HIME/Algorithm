package _0221020_Null_Null_Binary_Heap_Prority_Queue___Data_Struct_Design

import (
	"testing"
)

func TestBigHeapQueue(t *testing.T) {
	queue := InitBigHeapQueue(7)

	queue.put(1)
	queue.introduce()

	queue.del(1)
	queue.introduce()

	queue.put(2)
	queue.introduce()

	queue.put(3)
	queue.introduce()

	queue.put(4)
	queue.introduce()

	queue.put(5)
	queue.introduce()

	queue.put(6)
	queue.introduce()

	queue.put(7)
	queue.introduce()

	queue.del(2)
	queue.introduce()
}

func TestSmallHeapQueue(t *testing.T) {
	queue := InitSmallHeapQueue(7)

	queue.put(1)
	queue.introduce()

	queue.del(1)
	queue.introduce()

	queue.put(2)
	queue.introduce()

	queue.put(3)
	queue.introduce()

	queue.put(4)
	queue.introduce()

	queue.put(5)
	queue.introduce()

	queue.put(6)
	queue.introduce()

	queue.put(7)
	queue.introduce()

	queue.del(2)
	queue.introduce()

	queue.del(4)
	queue.introduce()
}
