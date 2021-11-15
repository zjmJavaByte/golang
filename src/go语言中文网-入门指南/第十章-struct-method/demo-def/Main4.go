package main

/*
递归结构体
*/
func main() {

}

/*
链表：
https://cdn.learnku.com/uploads/images/201808/27/23/T9sRQUcN1b.jpg?imageView2/2/w/1240/h/0
*/
//单向链表
type Node1 struct {
	data float64
	su   *Node1
}

//双向链表
type Node2 struct {
	pr   *Node2
	data float64
	su   *Node2
}

/*
二叉树：
https://cdn.learnku.com/uploads/images/201808/27/23/K6NRjDiDKD.jpg?imageView2/2/w/1240/h/0
*/
type Tree struct {
	le   *Tree
	data float64
	ri   *Tree
}
