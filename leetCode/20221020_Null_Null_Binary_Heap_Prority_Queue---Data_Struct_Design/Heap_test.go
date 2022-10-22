package _0221020_Null_Null_Binary_Heap_Prority_Queue___Data_Struct_Design

import (
	"testing"
)

func TestBigHeapQueue(t *testing.T) {
	queue := InitBigHeapQueue(7)

	queue.put(11)
	queue.introduce()

	queue.put(10)
	queue.introduce()

	queue.put(9)
	queue.introduce()

	queue.del(1)
	queue.introduce()
}

func TestSmallHeapQueue(t *testing.T) {
	queue := InitSmallHeapQueue(7)

	queue.put(9)
	queue.introduce()

	queue.put(10)
	queue.introduce()

	queue.put(11)
	queue.introduce()

	queue.del(1)
	queue.introduce()

}
