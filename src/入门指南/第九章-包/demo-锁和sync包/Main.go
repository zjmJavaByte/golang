package main

import (
	"fmt"
	"sync"
)

/*type info struct {
	mu sync.Mutex
}*/

type T struct {
	sync.RWMutex
	myMap map[int]int
}

func (t *T) getValue(key int) int {
	t.RLock()
	defer t.RUnlock()
	return t.myMap[key]
}

func (t *T) setValue(key int, value int) {
	t.Lock()
	defer t.Unlock()
	t.myMap[key] = value
}

//有问题  还是包并发错误
func main() {

	ty := T{myMap: map[int]int{}}

	wg := sync.WaitGroup{}
	wg.Add(50000)
	for i := 0; i < 50000; i++ {

		go func(i int) {
			ty.setValue(i, i)
			fmt.Printf("get key == %d, value == %d \n", i, ty.getValue(i))
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("program over !!")
}

/*func update(info *info) {
	info.mu.Lock()
	defer info.mu.Unlock()
	fmt.Println("锁住了")
}*/
