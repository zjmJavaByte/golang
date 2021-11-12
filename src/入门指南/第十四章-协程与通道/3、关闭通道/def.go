package main

func main() {

	/*
		第一个可以通过函数 close(ch) 来完成：这个将通道标记为无法通过发送操作 <- 接受更多的值；
		给已经关闭的通道发送或者再次关闭都会导致运行时的 panic。在创建一个通道后使用 defer 语句是个不错的办法（类似这种情况）
	*/
	ch := make(chan float64)
	defer close(ch)

	/*
		第二个问题可以使用逗号，ok 操作符：用来检测通道是否被关闭。

		v, ok := <-ch
			if !ok {
				break
			}
			process(v)
	*/

}
