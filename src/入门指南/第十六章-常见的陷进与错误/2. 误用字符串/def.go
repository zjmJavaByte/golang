package main

import (
	"bytes"
	"fmt"
)

/*
当需要对一个字符串进行频繁的操作时，谨记在 go 语言中字符串是不可变的（类似 java 和 c#）。
使用诸如 a += b 形式连接字符串效率低下，尤其在一个循环内部使用这种形式。这会导致大量的内存开销和拷贝。
应该使用一个字符数组代替字符串，将字符串内容写入一个缓存中。 例如以下的代码示例

*/
func main() {

	var b bytes.Buffer
	for i := 0; i < 100; i++ {
		b.WriteString("1")
	}
	fmt.Println(b.String())

}
