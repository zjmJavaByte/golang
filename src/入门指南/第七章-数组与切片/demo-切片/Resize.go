package main

import "fmt"

/**
GoLang中的切片扩容机制
*/
func main() {

	/**
	由于 a 的长度超出了容量，所以切片 a 指向了一个增长后的新数组，而 b 仍然指向原来的老数组。所以之后对 a 进行的操作，对 b 不会产生影响。
	*/
	//a := make([]int, 5)
	//b := a[0:5]
	//a = append(a, 1)
	//a[1] = 5
	//fmt.Println(b)
	//// [0 0 0 0]
	//fmt.Println(a)
	// [0 5 0 0 0 1]</pre>

	println()

	/**
	a 的容量为6，因此在 append 后并未超出容量，所以 array 指针没有改变。因此，对 a 进行的操作，对 b 同样产生了影响。
	*/
	//c := make([]int, 5, 6)
	//d := a[0:4]
	//c = append(c, 1)
	//c[1] = 5
	//fmt.Println(c, d)
	// [0 5 0 0 0 1] [0 5 0 0]

	/**
	下面看看用 a := []int{} 这种方式来创建切片会是什么情况
	*/
	a := []int{}
	for i := 0; i < 16; i++ {
		a = append(a, i)
		fmt.Print(cap(a), " ")
	}
	// 1 2 4 4 8 8 8 8 16 16 16 16 16 16 16 16

	/**
	扩容源码部分

	// src/runtime/slice.go
	func growslice(et *_type, old slice, cap int) slice {
	// ...省略部分
	    newcap := old.cap
	    doublecap := newcap + newcap
	    if cap > doublecap {
	        newcap = cap
	    } else {
	        if old.len < 1024 {
	            newcap = doublecap
	        } else {
	            // Check 0 < newcap to detect overflow
	            // and prevent an infinite loop.
	            for 0 < newcap && newcap < cap {
	                newcap += newcap / 4
	            }
	            // Set newcap to the requested cap when
	            // the newcap calculation overflowed.
	            if newcap <= 0 {
	                newcap = cap
	            }
	        }
	    }
	// ...省略部分
	}
	*/

	/**
	当需要的容量超过原切片容量的两倍时，会使用需要的容量作为新容量。
	当原切片长度小于1024时，新切片的容量会直接翻倍。而当原切片的容量大于等于1024时，会反复地增加25%，直到新容量超过所需要的容量。

	GoLang中的切片扩容机制，与切片的数据类型、原本切片的容量、所需要的容量都有关系，比较复杂。
	对于常见数据类型，在元素数量较少时，大致可以认为扩容是按照翻倍进行的。但具体情况需要具体分析。
	*/
	/**
	  参考链接
	    https://github.com/zjmJavaByte/JavaQaaQ/blob/master/docs/golang/%E5%88%87%E7%89%87/GoLang%E4%B8%AD%E7%9A%84%E5%88%87%E7%89%87%E6%89%A9%E5%AE%B9%E6%9C%BA%E5%88%B6.md
	*/

}
