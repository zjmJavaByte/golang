package main

import (
	"fmt"
	"strings"
)

func main() {

	//4.7.1 前缀和后缀

	s := "hello world~"
	prefix := strings.HasPrefix(s, "hello")
	suffix := strings.HasSuffix(s, "~")
	fmt.Printf("字符串 \"%s\" 是否以hello开头：%t \n", s, prefix)
	fmt.Printf("字符串 \"%s\" 是否以~结尾：%t \n", s, suffix)

	//4.7.2 字符串包含关系

	any := strings.ContainsAny(s, "o")
	fmt.Printf("字符串 \"%s\" 是否包含o：%t \n", s, any)

	//4.7.3 判断子字符串或字符在父字符串中出现的位置（索引)
	s2 := "o"
	indexAny := strings.IndexAny(s, s2)
	fmt.Printf("%d\n", indexAny)

	index := strings.LastIndex(s, s2)
	fmt.Printf("%d\n", index)

	indexRune := strings.IndexRune(s, 'w') //unicode码值r在s中第一次出现的位置，不存在则返回-1
	fmt.Printf("%d\n", indexRune)

	//4.7.4 字符串替换
	/**
	Replace 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串，如果 n = -1 则替换所有字符串 old 为字符串 new
	*/
	replace := strings.Replace(s, "world", "shanghai", -1)
	fmt.Printf("%v\n", replace)

	//4.7.5 统计字符串出现次数
	count := strings.Count(s, "l")
	fmt.Printf("%v\n", count)

	//4.7.6 重复字符串
	repeat := strings.Repeat(s, 2)
	fmt.Printf("%v\n", repeat)

	//4.7.7 修改字符串大小写
	lower := strings.ToLower(s)
	upper := strings.ToUpper(s)
	fmt.Printf("%v\n", lower)
	fmt.Printf("%v\n", upper)

	//4.7.8 修剪字符串
	/**
	剔除字符串开头和结尾的空白符号
	*/
	space := strings.TrimSpace("  1234567  ")
	fmt.Printf("%v\n", space)
	/**
	剔除指定字符
	*/
	trim := strings.Trim(s, "h")
	fmt.Printf("%v\n", trim)

	//4.7.9 分割字符串
	/**
	利用空白作为分隔符将字符串分割为若干块，并返回一个 slice 。如果字符串只包含空白符号，返回一个长度为 0 的 slice
	*/
	fields := strings.Fields(s)
	for i := range fields {
		fmt.Printf("%v\n", fields[i])
	}
	fmt.Printf("%v\n", fields)

	/**
	自定义分割符号对字符串分割
	*/
	split := strings.Split(s, "o")
	for i := range fields {
		fmt.Printf("%v\n", split[i])
	}
	fmt.Printf("%v\n", split)

	//4.7.10 拼接 slice 到字符串
	/**
	Join 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串
	*/
	join := strings.Join(split, "|")
	fmt.Printf("%v\n", join)

	//4.7.11 从字符串中读取内容
	reader := strings.NewReader(s)
	i := reader.Len()
	fmt.Printf("%v\n", i)
	//reader.ReadAt()
}
