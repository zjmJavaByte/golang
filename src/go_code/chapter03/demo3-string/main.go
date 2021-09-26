package main

import "fmt"

/*
字符串类型
字符串是一串固定长度的字符连接起来的字符序列
*/
func main() {

	var a string = "朱健民"
	fmt.Println(a)

	b := "zjm"
	//b[0] = 'y' 这里就不能修改b的内容，即go中的字符串是不可变的
	fmt.Println(b)

	//反引号,以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等
	d := `
		package main
		
		import "fmt"
		
		/*
		字符类型
		*/
		func main() {
		
		
			//参考ASCII
		
			var a byte = '\t'
			b := '1'
			c := 'a'
			d := '{'
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
		
			//如果希望输出对应的字符，需要使用格式化输出
			fmt.Printf("a=%c b=%c c=%c d=%c  \n",a,b,c,d)
		
		
		
			//var e byte = '朱'   编译报错
			var e int = '朱'
			fmt.Printf("e=%c e对应的码值：%d \n",e,e)  //e=朱 e对应的码值：26417
		
			//直接给某个变量赋值，然后格式化输出对应的unicode字符
			var f int = 122
			fmt.Printf("f=%c  \n",f) //z
		
		}
		`
	fmt.Println(d)

	//字符串拼接
	var str = "hello" + "word"
	str += "123"
	fmt.Println(str)

	/*
		这一种形式是错误的  + 应该放在每一行的最后，因为go会默认在每一行最后加上一个 ；
		var str2 = "1" + "2" + "3"
			+ "4" + "5" + "" + "6"
	*/
	var str2 = "1" + "2" + "3" +
		"4" + "5" + "" + "6"
	fmt.Println(str2)
}
