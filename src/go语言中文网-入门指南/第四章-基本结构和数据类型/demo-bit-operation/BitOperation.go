package main

/**
位运算

位运算只能用于整数类型的变量，且需当它们拥有等长位模式时。

%b 是用于表示位的格式化标识符

*/
func main() {

	/*
		按位与 &：

		对应位置上的值经过和运算结果，具体参见和运算符，第 4.5.1 节，并将 T（true）替换为 1，将 F（false）替换为 0
		1 & 1 -> 1
		1 & 0 -> 0
		0 & 1 -> 0
		0 & 0 -> 0
	*/
	println(1 & 1)
	println(1 & 0)
	println(0 & 1)
	println(0 & 0)
	println()

	/*
		按位或 |：

		对应位置上的值经过或运算结果，具体参见或运算符，第 4.5.1 节，并将 T（true）替换为 1，将 F（false）替换为 0
		1 | 1 -> 1
		1 | 0 -> 1
		0 | 1 -> 1
		0 | 0 -> 0
	*/
	println(1 | 1)
	println(1 | 0)
	println(0 | 1)
	println(0 | 0)
	println()

	/*
		按位异或 ^：

		对应位置上的值根据以下规则组合：

		1 ^ 1 -> 0
		1 ^ 0 -> 1
		0 ^ 1 -> 1
		0 ^ 0 -> 0
	*/
	println(1 ^ 1)
	println(1 ^ 0)
	println(0 ^ 1)
	println(0 ^ 0)
	println()

	/*
		位清除 &^：a&^b 就是把a中对应b为1的位清为0
	*/
	println(0110 &^ 1011)
	println(1011 &^ 1101)
	println()

	a := ^2
	b := ^10
	c := -01 ^ 10
	println(a)
	println(b)
	println(c)
	//= ^10 = -01 ^ 10 = -11

}
