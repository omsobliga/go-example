// 工厂方法创建设计模式允许创建对象，而无需指定将要创建对象的确切类型。
package main

import "fmt"

type Store interface {
	Open(string) error
}

type DiskStore struct{}

func (DiskStore) Open(name string) error {
	fmt.Println("open disk", name)
	return nil
}

func NewDiskStore() Store {
	return DiskStore{}
}

type MemoryStore struct{}

func (MemoryStore) Open(name string) error {
	fmt.Println("open memory", name)
	return nil
}

func NewMemoryStore() Store {
	return MemoryStore{}
}

func NewStore(name string) Store {
	switch name {
	case "disk":
		return NewDiskStore()
	case "memory":
		return NewMemoryStore()
	}
	return nil
}

func main() {
	store := NewStore("disk")
	store.Open("aaa")

	store = NewStore("memory")
	store.Open("bbb")
}
