package person

import "sync"

type Person struct {
	sync.RWMutex
	firstName string
	lastName  string
	student   Student
}

type Student struct {
	name string
}

func (receiver *Person) GetFirstName() string {
	/*
		对象的字段（属性）不应该由 2 个或 2 个以上的不同线程在同一时间去改变。如果在程序发生这种情况，为了安全并发访问，可以使用包 sync 中的方法
	*/
	receiver.RLock()
	defer receiver.RUnlock()
	return receiver.firstName
}

func (receiver *Person) SetFirstName(firstName string) {
	receiver.Lock()
	defer receiver.Unlock()
	receiver.firstName = firstName
}
