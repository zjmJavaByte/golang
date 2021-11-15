package main

import (
	"fmt"
	"reflect"
)

func main() {

	/*
		reflect.Type 和 reflect.Value 都有许多方法用于检查和操作它们。一个重要的例子是 Value 有一个 Type 方法返回 reflect.Value 的 Type。
		另一个是 Type 和 Value 都有 Kind 方法返回一个常量来表示类型：Uint、Float64、Slice 等等。同样 Value 有叫做 Int 和 Float 的方法可以获取存储在内部的值

	*/
	/*var x = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)


	v2 := reflect.TypeOf(x)
	fmt.Println("value:", v2)
	fmt.Println("kind:", v2.Kind())*/

	/*
		通过反射修改 (设置) 值

		假设我们要把 x 的值改为 3.1415。Value 有一些方法可以完成这个任务，但是必须小心使用：v.SetFloat(3.1415)。

		这将产生一个错误：reflect.Value.SetFloat using unaddressable value。

		为什么会这样呢？问题的原因是 v 不是可设置的（这里并不是说值不可寻址）。是否可设置是 Value 的一个属性，并且不是所有的反射值都有这个属性：可以使用 CanSet() 方法测试是否可设置。

		在例子中我们看到 v.CanSet() 返回 false： settability of v: false

		当 v := reflect.ValueOf(x) 函数通过传递一个 x 拷贝创建了 v，那么 v 的改变并不能更改原始的 x。要想 v 的更改能作用到 x，那就必须传递 x 的地址 v = reflect.ValueOf(&x)。

		通过 Type () 我们看到 v 现在的类型是 *float64 并且仍然是不可设置的。

		要想让其可设置我们需要使用 Elem() 函数，这间接的使用指针：v = v.Elem()

		现在 v.CanSet() 返回 true 并且 v.SetFloat(3.1415) 设置成功了！
	*/
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// setting a value:
	// v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())
	fmt.Println(v)

	/*
		反射结构体
	*/
}
