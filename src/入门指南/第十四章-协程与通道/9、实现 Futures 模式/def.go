package main

func main() {

	/*
		参考条目文献给出了一个很精彩的例子：假设我们有一个矩阵类型，我们需要计算两个矩阵 A 和 B 乘积的逆，
		首先我们通过函数 Inverse(M) 分别对其进行求逆运算，在将结果相乘。

		func InverseProduct(a Matrix, b Matrix) {
		    a_inv := Inverse(a)
		    b_inv := Inverse(b)
		    return Product(a_inv, b_inv)
		}
		上面这个函数我们必须等待A计算完之后再去计算B
	*/

	/*

		如下的方式，只要两个协程都计算万后就可以计算了

		func InverseProduct(a Matrix, b Matrix) {
		    a_inv_future := InverseFuture(a)   // start as a goroutine
		    b_inv_future := InverseFuture(b)   // start as a goroutine
		    a_inv := <-a_inv_future
		    b_inv := <-b_inv_future
		    return Product(a_inv, b_inv)
		}

		func InverseFuture(a Matrix) {
		    future := make(chan Matrix)
		    go func() {
		        future <- Inverse(a)
		    }()
		    return future
		}
	*/

}
