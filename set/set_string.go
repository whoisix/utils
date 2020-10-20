package set

import (
	"sort"
	"sync"
)

// 实现 set 集合，变相实现 切片去重
type stringSet struct {
	m map[string]struct{}
	sync.RWMutex
}

func NewStringSet() *stringSet {
	return &stringSet{
		m: make(map[string]struct{}),
	}
}

func (s *stringSet) Add(items ...string) {
	s.Lock()
	defer s.Unlock()
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		s.m[item] = struct{}{}
	}
}

func (s *stringSet) Remove(items ...string) {
	s.Lock()
	defer s.Unlock()
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		delete(s.m, item)
	}
}

func (s *stringSet) Has(item string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *stringSet) Len() int {
	return len(s.List())
}

func (s *stringSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = make(map[string]struct{})
}

func (s *stringSet) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *stringSet) List() []string {
	s.RLock()
	defer s.RUnlock()
	var list []string
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func (s *stringSet) SortList() []string {
	s.RLock()
	defer s.RUnlock()
	var list []string
	for item := range s.m {
		list = append(list, item)
	}
	sort.Strings(list)
	return list
}