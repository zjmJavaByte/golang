package main

import (
	"runtime"
	"time"
)

const NCPU = 8

func DoAll() {

	sem := make(chan int, NCPU)

	for i := 0; i < NCPU; i++ {

		// Buffering optional but sensible. 合理的缓冲区选项（个人理解就是和 CPU 的核心数相同）

		go DoPart(sem)

	}

	// 等待 NCPU 任务完成，释放通道 sem 的缓冲区

	for i := 0; i < NCPU; i++ {

		<-sem // 等待一个任务完成

	}

	// 全部完成。

}

func DoPart(sem chan int) {

	// 进行计算的部分
	time.Sleep(1e9)

	sem <- 1 // 发送一个这部分已经完成的信号，用来释放 sem 的缓冲区

}

func main() {

	runtime.GOMAXPROCS(NCPU)

	DoAll()

}
