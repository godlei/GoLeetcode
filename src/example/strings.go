package example

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//字符串操作示例
func StringExample() {
	//string --> int
	var str string
	var n int
	var n64 int64
	var f float64

	str = "28"
	n, _ = strconv.Atoi(str)
	fmt.Println("string --> int")
	fmt.Println(n)
	fmt.Println(reflect.TypeOf(n))
	//int --> string'
	fmt.Println("int --> string")
	n = 33
	str = strconv.Itoa(n)
	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))

	//string --> int64
	fmt.Println("string --> int64")
	str = "123"
	n64, _ = strconv.ParseInt(str, 10, 64)
	fmt.Println(n64)
	fmt.Println(reflect.TypeOf(n64))

	//int64 --> string
	fmt.Println("int64 --> string")
	n64 = 12312
	str = strconv.FormatInt(n64, 10)
	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))

	//string --> float
	fmt.Println("string --> float64")
	str = "12.34"
	f, _ = strconv.ParseFloat(str, 10)
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))

	//float --> string
	fmt.Println("float64 --> string")
	f = 123.456
	str = strconv.FormatFloat(f, 'f', -1, 64)
	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))

	//字符串比较
	fmt.Println("字符串比较")
	str1 := "abc"
	str2 := "adc"
	fmt.Println(strings.Compare(str1, str2))

	//字符串包含  (s1是否包含s2)
	fmt.Println("字符串包含")
	str1 = "hello,word"
	str2 = "o"
	fmt.Println(strings.Contains(str1, str2))

	//查找位置
	fmt.Println(strings.Index(str1, str2))
	fmt.Println(strings.LastIndex(str1, str2))

	//统计给定子串sep的出现次数, sep为空时, 返回1 + 字符串的长度
	fmt.Println("子串出现个数查询")
	str = "godwang,gogogo"
	fmt.Println(strings.Count(str, "go"))
	fmt.Println(strings.Count(str, ""))

	//重复s字符串count次, 最后返回新生成的重复的字符串
	fmt.Println("构建重复的字符串")
	fmt.Println(strings.Repeat("wang", 4))

	//字符串替换
	str = "hello,godwang,godwang,godwang"
	str = strings.Replace(str, "wang", "lei", -1) //n表示替换的次数，小于0表示全部替换
	fmt.Println(str)

	//删除两边的特殊字符
	str = "//Users//Documents/GOPatch/src/MyGO/config/TestString/"
	str = strings.Trim(str, "/") //Users//Documents/GOPatch/src/MyGO/config/TestString
	fmt.Println(str)

	//删除两边空格
	str = "  sss  sss    s    "
	fmt.Println(strings.TrimSpace(str))

	//前缀 后缀 HasPrefix HasSuffix
	str = "hello world"
	fmt.Println(strings.HasPrefix(str, "he"))
	fmt.Println(strings.HasSuffix(str, "world"))

	//字符串分割
	str = "nice&to&meet&&you"
	strArry := strings.Split(str, "&")
	fmt.Println(strArry, len(strArry))

	//字符串拼接
	strArry = []string{"key1=val1", "key2=val2", "key3=val3"}
	str = strings.Join(strArry, "&")
	fmt.Println(str)
}
