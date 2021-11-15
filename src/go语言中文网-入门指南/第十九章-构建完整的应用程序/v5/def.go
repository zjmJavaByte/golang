package main

import (
	"encoding/gob"
	"errors"
	"log"
	"net/rpc"
	"os"
	"sync"
)

var saveQueueLength = 500

type Store interface {
	Put(url, key *string) error
	Get(key, url *string) error
}

type record struct {
	key, url string
}

type UrlStore struct {
	mu sync.RWMutex

	urls map[string]string

	save chan record
}

type ProxyStore struct {
	client *rpc.Client
	urls   *UrlStore
}

func NewProxyStore(addr string) *ProxyStore {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Println("Error constructing ProxyStore:", err)
	}
	return &ProxyStore{client: client, urls: NewUrlStore("")}
}

func (s *ProxyStore) Get(key, url *string) error {
	if err := s.urls.Get(key, url); err == nil { //如果本地存在就在本地获取
		return nil
	}
	//本地map中没有找到，就远程调用
	if err := s.client.Call("Store.Get", key, url); err != nil {
		return err
	}
	s.urls.Set(url, key)
	return nil
}
func (s *ProxyStore) Put(url, key *string) error {
	return s.client.Call("Store.Put", url, key)
}

func NewUrlStore(filename string) *UrlStore {
	s := &UrlStore{urls: make(map[string]string)}
	if filename != "" {
		s.save = make(chan record, saveQueueLength)
		_, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal("Error opening URLStore:", err)
		}
	}
	return s
}

func (s *UrlStore) Get(key, url *string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if u, present := s.urls[*key]; present {
		*url = u
		return nil
	}
	return errors.New("key not found")
}

func (s *UrlStore) Put(url, key *string) error {
	for true {
		*key = genKey(s.Count())
		if err := s.Set(url, key); err == nil {
			break
		}
	}
	if s.save != nil {
		s.save <- record{*key, *url}
	}
	return errors.New("")
}

func (s *UrlStore) Set(url, key *string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.urls[*key]; present {
		return errors.New("key already exist")
	}
	s.urls[*key] = *url
	return nil
}

func (s *UrlStore) saveLoop(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error opening URLStore: ", err)
	}
	for true {
		e := gob.NewEncoder(f)
		save := <-s.save
		err := e.Encode(save)
		if err != nil {
			log.Println("Error saving to URLStore: ", err)
		}
	}
}

func (s *UrlStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}
