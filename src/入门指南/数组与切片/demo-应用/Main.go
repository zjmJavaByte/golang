package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {

	s := "\u00ff\u754c"
	//获取一个字节数组的切片 c
	c := []byte(s)
	for i, b := range c {
		fmt.Printf("%d:%c \n", i, b)
	}
	fmt.Println()
	for i, c := range s {
		fmt.Printf("%d:%c \n", i, c)
	}

	s2 := "12345"
	var bytes = make([]byte, 0)
	copy(bytes, s2)
	fmt.Println(bytes)

	f := []int32(s)
	fmt.Printf("[]int32(s) %v \n", f)
	r := []rune(s)
	fmt.Printf("[]rune(s) %v \n", r)
	i := len([]int32(s))
	fmt.Printf("[]int32(s)长度是： %v \n", i)
	inString := utf8.RuneCountInString(s)
	fmt.Printf("utf8.RuneCountInString长度是： %v \n", inString)

	fmt.Println("--------------7.6.2 获取字符串的某一部分----------------------")
	str := "12345"
	substr := str[2:4]
	fmt.Println(substr)

	fmt.Println("--------------7.6.4 修改字符串中的某个字----------------------")
	str2 := "hello"
	c2 := []byte(str2)
	c2[0] = 'c'
	s3 := string(c2) // s3 == "cello"
	fmt.Println(s3)

	str3 := "你好 世界"
	replace := strings.Replace(str3, "你", "哈", 3)
	fmt.Println(str3)
	fmt.Println(replace)

	fmt.Println("--------------7.6.5 字节数组对比函数----------------------")
	str4 := "cello"
	compare := Compare(c2, []byte(str4))
	fmt.Println(compare)

	fmt.Println("--------------7.6.6 搜索及排序切片和数组----------------------")
	//int 类型的切片排序
	arri := []int{4, 6, 1, 3, 8, 2, 9}
	sort.Ints(arri)
	fmt.Println(arri)
	sorted := sort.IntsAreSorted(arri) //判断切片是否是有序的
	fmt.Println(sorted)
	ints := sort.SearchInts(arri, 3) //返回对应结果的索引值
	fmt.Printf("索引值：%d", ints)

	//float64 类型的切片排序
	arri2 := []float64{1.3, 6.3, 2.5}
	sort.Float64s(arri2)
	fmt.Println(arri2)
	sorted2 := sort.Float64sAreSorted(arri2) //判断切片是否是有序的
	fmt.Println(sorted2)
	ints2 := sort.SearchFloat64s(arri2, 6.3) //返回对应结果的索引值
	fmt.Printf("索引值：%d", ints2)

	//string 类型的切片排序
	arri3 := []string{"你", "好", "世", "界"}
	sort.Strings(arri3)
	fmt.Println(arri3)
	sorted3 := sort.StringsAreSorted(arri3) //判断切片是否是有序的
	fmt.Println(sorted3)
	ints3 := sort.SearchStrings(arri3, "世") //返回对应结果的索引值
	fmt.Printf("索引值：%d", ints3)

	fmt.Println("--------------7.6.7 append 函数常见操作----------------------")

	/*
		将切片 b 的元素追加到切片 a 之后：a = append(a, b...)

		复制切片 a 的元素到新的切片 b 上：
		 b = make([]T, len(a))
		 copy(b, a)

		删除位于索引 i 的元素：a = append(a[:i], a[i+1:]...)

		切除切片 a 中从索引 i 至 j 位置的元素：a = append(a[:i], a[j:]...)

		为切片 a 扩展 j 个元素长度：a = append(a, make([]T, j)...)

		在索引 i 的位置插入元素 x：a = append(a[:i], append([]T{x}, a[i:]...)...)

		在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]T, j), a[i:]...)...)

		在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]...)...)

		取出位于切片 a 最末尾的元素 x：x, a = a[len(a)-1:], a[:len(a)-1]

		将元素 x 追加到切片 a：a = append(a, x)
	*/

	fmt.Println("--------------7.6.8 切片和垃圾回收----------------------")
	/*
		切片的底层指向一个数组，该数组的实际容量可能要大于切片所定义的容量。只有在没有任何切片指向的时候，底层的数组内存才会被释放，这种特性有时会导致程序占用多余的内存。

		示例 函数 FindDigits 将一个文件加载到内存，然后搜索其中所有的数字并返回一个切片。

		var digitRegexp = regexp.MustCompile("[0-9]+")

		func FindDigits(filename string) []byte {
		    b, _ := ioutil.ReadFile(filename)
		    return digitRegexp.Find(b)
		}
		这段代码可以顺利运行，但返回的 []byte 指向的底层是整个文件的数据。只要该返回的切片不被释放，垃圾回收器就不能释放整个文件所占用的内存。换句话说，一点点有用的数据却占用了整个文件的内存。

		想要避免这个问题，可以通过拷贝我们需要的部分到一个新的切片中：

		func FindDigits(filename string) []byte {
		   b, _ := ioutil.ReadFile(filename)
		   b = digitRegexp.Find(b)
		   c := make([]byte, len(b))
		   copy(c, b)
		   return c
		}
		事实上，上面这段代码只能找到第一个匹配正则表达式的数字串。要想找到所有的数字，可以尝试下面这段代码：

		func FindFileDigits(filename string) []byte {
		   fileBytes, _ := ioutil.ReadFile(filename)
		   b := digitRegexp.FindAll(fileBytes, len(fileBytes))
		   c := make([]byte, 0)
		   for _, bytes := range b {
		      c = append(c, bytes...)
		   }
		   return c
		}
	*/

	fmt.Println("---------------练习----------------------")
	/*
		练习 7.12
		编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。
	*/
	a, b := split("qwertyuiop", 5)
	fmt.Printf("分隔后的字符串1：%v 分隔后的字符串2：%v \n", a, b)

	/*
		练习 7.13
		假设有字符串 str，那么 str[len(str)/2:] + str[:len(str)/2] 的结果是什么？
	*/
	strs := str[len(str)/2:] + str[:len(str)/2]
	fmt.Println(strs)

	/*
		练习 7.14
		编写一个程序，要求能够反转字符串，即将 “Google” 转换成 “elgooG”（提示：使用 []byte 类型的切片）。
		如果您使用两个切片来实现反转，请再尝试使用一个切片（提示：使用交换法）。
		如果您想要反转 Unicode 编码的字符串，请使用 []int32 类型的切片。
	*/
	i2 := []byte("Google")
	var temps byte
	i3 := len(i2)
	for i := 0; i < i3/2; i++ {
		temps = i2[i]
		i2[i] = i2[i3-1-i]
		i2[i3-1-i] = temps
	}
	fmt.Printf("%c \n", i2)

	/*
		练习 7.15
		编写一个程序，要求能够遍历一个数组的字符，并将当前字符和前一个字符不相同的字符拷贝至另一个数组。
	*/

	sss := []byte{'d', 'f', 'w', 'w', 'g'}
	var temp byte
	var i4 []byte
	for i := 1; i < len(sss); i++ {
		if i != len(sss)-1 {
			temp = sss[i-1]
		}
		if temp != sss[i] {
			i4 = append(i4, sss[i])
		}

	}
	fmt.Printf("%c \n", i4)

	/*
		练习 7.16

		编写一个程序，使用冒泡排序的方法排序一个包含整数的切片（算法的定义可参考 维基百科）。
	*/
	buble := []int{3, 5, 1, 7, 3}
	var tempss int
	for i := 0; i < len(buble); i++ {
		for j := i; j < len(buble); j++ {
			i5 := buble[i]
			i6 := buble[j]
			if i5 > i6 {
				tempss = i5
				buble[i] = i6
				buble[j] = tempss
			}
		}
	}
	fmt.Println(buble)

	/*
		练习 7.17

		在函数式编程语言中，一个 map-function 是指能够接受一个函数原型和一个列表，并使用列表中的值依次执行函数原型，
		公式为：map ( F(), (e1,e2, . . . ,en) ) = ( F(e1), F(e2), ... F(en) )。

		编写一个函数 mapFunc 要求接受以下 2 个参数：

		一个将整数乘以 10 的函数
		一个整数列表
		最后返回保存运行结果的整数列表。
	*/

	i5 := mapFunc([]int{1, 2, 3, 4, 5}, mf)
	fmt.Println(i5)
}

func mapFunc(arr []int, mf func(a int) int) []int {
	for i := range arr {
		arr[i] = mf(arr[i])
	}
	return arr
}

func mf(int2 int) int {
	return int2 * 10
}

// Compare 比较两个字符串字符是否相等
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	// 数组的长度可能不同
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0 // 数组相等
}

func split(str string, i int) (a, b string) {
	a = str[:i]
	b = str[i:]
	return
}
