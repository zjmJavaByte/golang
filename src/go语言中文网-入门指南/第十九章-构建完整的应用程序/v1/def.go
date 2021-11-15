package main

import (
	"sync"
)

type urlMap map[string]string

type UrlStore struct {
	urls urlMap

	sy sync.RWMutex
}

func (s *UrlStore) Get(key string) string {
	s.sy.RLock()
	defer s.sy.RUnlock()
	return s.urls[key]
}

func (s *UrlStore) Set(key, value string) bool {
	s.sy.Lock()
	defer s.sy.Unlock()
	_, present := s.urls[key]
	if present {
		return false
	}
	s.urls[key] = value
	return true
}

// NewUrlStore 函数工厂
func NewUrlStore() *UrlStore {
	return &UrlStore{urls: make(map[string]string)}
}

func (s *UrlStore) Count() int {
	s.sy.RLock()
	defer s.sy.RUnlock()
	return len(s.urls)
}

func (s *UrlStore) Put(url string) string {
	for true {
		key := genKey(s.Count())
		if s.Set(key, url) {
			return key
		}
	}
	return ""
}
