package main

import "log"

type Work interface {
}

func main() {

}

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work) // start the goroutine for that work
	}
}

func safelyDo(work *Work) {
	defer func() {
		//一个用到 recover 的程序停掉了服务器内部一个失败的协程而不影响其他协程的工作
		if err := recover(); err != nil {
			log.Printf("Work failed with %s in %v", err, work)
		}
	}()
	//上边的代码，如果 do(work) 发生 panic，错误会被记录且协程会退出并释放，而其他协程不受影响
	do(work)
}

func do(work *Work) {

}
