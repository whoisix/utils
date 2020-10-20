package set

import (
	"sort"
	"sync"
)

// 实现 set 集合，变相实现 切片去重
type intSet struct {
	m map[int]struct{}
	sync.RWMutex
}

func NewIntSet() *intSet {
	return &intSet{
		m: make(map[int]struct{}),
	}
}

func (s *intSet) Add(items ...int) {
	s.Lock()
	defer s.Unlock()
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		s.m[item] = struct{}{}
	}
}

func (s *intSet) Remove(items ...int) {
	s.Lock()
	defer s.Unlock()
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		delete(s.m, item)
	}
}

func (s *intSet) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *intSet) Len() int {
	return len(s.List())
}

func (s *intSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = make(map[int]struct{})
}

func (s *intSet) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *intSet) List() []int {
	s.RLock()
	defer s.RUnlock()
	var list []int
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func (s *intSet) SortList() []int {
	s.RLock()
	defer s.RUnlock()
	var list []int
	for item := range s.m {
		list = append(list, item)
	}
	sort.Ints(list)
	return list
}