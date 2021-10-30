package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
}

type Teacher struct {
	name string
	/*
		当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，这在效果上等同于外层类型 继承 了这些方法：将父类型放在子类型中来实现亚型
	*/
	Student
}

type Student struct {
	name string
}

func main() {
	t := new(Teacher)
	t.Student.name = "yap"
	t.work("zjm")
}

func (receiver Car) GoToWorkIn() {
	receiver.Start()
	receiver.Stop()
}

func (receiver Student) work(teacherName string) {
	fmt.Println(teacherName + "要求学生" + receiver.name + "写作业。。。")
}

/*
和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法
*/
func (receiver Teacher) work(studentName string) {
	fmt.Println(receiver.name + "要求学生" + studentName + "写作业。。。")
}
