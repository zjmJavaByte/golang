package main

func main() {

	/*
		在 第 4.9 小节，我们已经知道，切片实际是一个指向潜在数组的指针。
		我们常常需要把切片作为一个参数传递给函数是因为：实际就是传递一个指向变量的指针，在函数内可以改变这个变量，而不是传递数据的拷贝。

		因此应该这样做：

		   func findBiggest( listOfNumbers []int ) int {}
		而不是：

		   func findBiggest( listOfNumbers *[]int ) int {}
		当切片作为参数传递时，切记不要解引用切片。

	*/

	/*

		参考文章：https://www.cnblogs.com/pythonywy/p/11900551.html


		&变量 获取变量在内存空间的地址

		*变量地址 获取变量的值

		一.普通数据

		package main

		import "fmt"

		func main(){
			b :=1111
			c :=&b  //获取b的地址c的类型时*int
			test(c)
			fmt.Println(b)  //值为333发送了变化
		}
		func test(a *int){
			*a=333
		}
		//可以与下面进行对比
		func main(){
			b :=1111
			test(b)
			fmt.Println(b)
		}
		func test(a int){  //如果不是传入地址,他就会开辟一个新的内存空间对于原来值没有影响
			a=333
		}
		二.数组与切片(切片比较特殊)

		//写法一
		package main

		import "fmt"

		func main(){
			b :=&[]int{1,2,3}
			test(b)
			fmt.Println(b)
		}
		func test(a *[]int){
			(*a)[1]=3
		}

		//GO对于切片做了优化可以省略写内容
		package main

		import "fmt"

		func main(){
			b :=[]int{1,2,3}
			test(b)
			fmt.Println(b)
		}
		func test(a []int){
			a[1]=3
		}
		//如果传入对象是值类型,不是引用类型这个不生效,只正对引用类型切片才生效,数组值类型不生效,只能按照方式一写
	*/
}
