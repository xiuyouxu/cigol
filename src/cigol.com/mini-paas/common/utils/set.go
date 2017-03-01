package utils

import (
	"sync"
)

type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}

func NewSet(items ...interface{}) *Set {
	s := &Set{
		m: map[interface{}]bool{},
	}
	s.Add(items...)
	return s
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
