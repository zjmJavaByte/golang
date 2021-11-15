package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"sync"
)

/*

优化了二个点：
1、将保存到磁盘的操作异步到chan中
2、减少了Encoder的创建
*/

type urlStore struct {
	mu sync.RWMutex

	urls map[string]string

	save chan record
}

type record struct {
	key, url string
}

func (r *urlStore) saveLoop(filename string) {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println("文件不存在")
	}

	e := gob.NewEncoder(file)

	s := <-r.save // takes a record from the channel
	if err := e.Encode(s); err != nil {
		log.Println("Error saving to URLStore: ", err)
	}

}
