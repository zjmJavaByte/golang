package main

import (
	"fmt"
	"strconv"
)

type T struct {
	a int
	b float32
	c string
}

type Celsius float64

type TZ int

type Day int

var dayName = [7]string{"1", "2", "3", "4", "5", "6", "7"}

const (
	HOUR TZ = 60 * 60
	UTC  TZ = 0 * HOUR
	EST  TZ = -5 * HOUR
	CST  TZ = -6 * HOUR
)

var m = map[TZ]string{
	UTC: "Universal Greenwich time",
	EST: "Eastern Standard time",
	CST: "Central Standard time"}

const (
	MO Day = iota //关于iota  https://studygolang.com/articles/2192
	TU
	WE
	TH
	FR
	SA
	SU
)

/*

如果类型定义了 String() 方法，它会被用在 fmt.Printf() 中生成默认的输出：等同于使用格式化描述符 %v 产生的输出。
还有 fmt.Print() 和 fmt.Println() 也会自动使用 String() 方法。
*/
func main() {

	/*
		给 T 定义 String()，使得 fmt.Printf("%v\n", t) 输出：7 / -2.350000 / "abc\tdef"
	*/
	t := new(T)
	t.a = 7
	t.b = -2.350000
	t.c = "abc\tdef"
	//  7 / -2.350000 / abc\tdef
	//  7 / -2.350000 / "abc\tdef"
	fmt.Printf("%v\n", t)

	/*
		为 float64 定义一个别名类型 Celsius，并给它定义 String()，它输出一个十进制数和 °C 表示的温度值。
	*/
	var c Celsius = 18.36
	fmt.Println(c)

	var tt T
	tt.a = 1
	fmt.Println(tt)
	fmt.Printf("%T %T \n", t, tt)

	/*
		为 int 定义别名类型 TZ，定义一些常量表示时区，比如 UTC，定义一个 map，它将时区的缩写映射为它的全称，比如：UTC -> "Universal Greenwich time"。
		为类型 TZ 定义 String() 方法，它输出时区的全称
	*/
	fmt.Println(EST) // Print* knows about method String() of type TZ
	fmt.Println(0 * HOUR)
	fmt.Println(-6 * HOUR)

	/*
		为 int 定义一个别名类型 Day，定义一个字符串数组它包含一周七天的名字，为类型 Day 定义 String() 方法，
		它输出星期几的名字。使用 iota 定义一个枚举常量用于表示一周的中每天（MO、TU...）
	*/

	var th Day
	th = 3
	fmt.Println(th)

	var day = SU
	fmt.Println(day) // prints Sunday
	fmt.Println(0, MO, 1, TU)
}

func (t T) String() string {
	return fmt.Sprintf("%d %f %q", t.a, t.b, t.c)
}

func (c Celsius) String() string {
	return "The temperature is: " + strconv.FormatFloat(float64(c), 'f', 1, 32) + " °C"
}

/*
if _, ok := map1[key1]; ok {
    // ...
}
*/
func (tz TZ) String() string {
	if zone, ok := m[tz]; ok {
		return zone
	}
	return ""
}

func (d Day) String() string {
	return dayName[d]
}
