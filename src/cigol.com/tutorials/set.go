package main

import (
	"fmt"
	//	"sort"
	"sync"
)

type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: map[interface{}]bool{},
	}
}

func (s *Set) Add(items ...interface{}) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		s.m[item] = true
	}
}

func (s *Set) Remove(items ...interface{}) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		delete(s.m, item)
	}
}

func (s *Set) Contains(item interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) Len() int {
	return len(s.List())
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[interface{}]bool{}
}

func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *Set) List() []interface{} {
	s.RLock()
	defer s.RUnlock()
	list := []interface{}{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

//func (s *Set) SortList() []interface{} {
//	s.RLock()
//	defer s.RUnlock()
//	list := []interface{}{}
//	for item := range s.m {
//		list = append(list, item)
//	}
//	//	sort.Ints(list)
//	sort.
//	sort.Sort(list)
//	return list
//}

func main() {
	//初始化
	s := New()

	s.Add(1, 1, 0, 2, 4, 3)

	s.Clear()
	if s.IsEmpty() {
		fmt.Println("0 item")
	}

	s.Add(1, 2, 3)

	if s.Contains(2) {
		fmt.Println("2 does exist")
	}

	s.Remove(2)
	s.Remove(3)
	fmt.Println("无序的切片", s.List())
	//	fmt.Println("有序的切片", s.SortList())

}
