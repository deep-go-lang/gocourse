package advanced

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 5

type Buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func NewBuffer(size int) *Buffer {
	b := &Buffer{
		items: make([]int, 0, size),
	}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *Buffer) Produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for len(b.items) == bufferSize {
		b.cond.Wait()
	}
	b.items = append(b.items, item)

	fmt.Printf("Produced: %d\n", item)

	b.cond.Signal()
}

func (b *Buffer) Consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	for len(b.items) == 0 {
		b.cond.Wait()
	}

	item := b.items[0]
	b.items = b.items[1:]

	fmt.Printf("Consumed: %d\n", item)

	b.cond.Signal()
	return item
}

func producer(b *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		b.Produce(i + 100)
		time.Sleep(time.Millisecond * 500)
	}
}

func consumer(b *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		b.Consume()
		time.Sleep(time.Millisecond * 500)
	}
}
func main() {

	buffer := NewBuffer(bufferSize)

	var wg sync.WaitGroup

	wg.Add(2)
	go producer(buffer, &wg)
	go consumer(buffer, &wg)

	wg.Wait()
}
