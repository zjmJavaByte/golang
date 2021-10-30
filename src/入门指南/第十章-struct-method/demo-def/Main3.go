package main

/**
结构体的内存布局：
Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。
不像 Java 中的引用类型，一个对象和它里面包含的对象可能会在不同的内存空间中，这点和 Go 语言中的指针很像
*/
func main() {

	//type Rect1 struct {Min, Max Point }
	//type Rect2 struct {Min, Max *Point }

	//https://cdn.learnku.com/uploads/images/201808/27/23/o4SlYwuXqy.jpg?imageView2/2/w/1240/h/0

}
