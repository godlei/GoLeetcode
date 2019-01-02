package example

/*
	map操作相关
*/
import (
	"fmt"
	"sort"
)

func MapExample() {
	m := make(map[string]int)
	m["Tom"] = 81
	m["Aline"] = 72
	m["Monica"] = 75
	m["God"] = 89

	//遍历map
	for key, value := range m {
		fmt.Printf("%s - %d\n", key, value)
	}

	//删除元素
	delete(m, "Tom")

	//排序
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s - %d\n", key, m[key])
	}

	//元素是否存在
	v, ok := m["Tom"]
	if !ok {
		fmt.Println("key不存在")
	} else {
		fmt.Println(v)
	}
}
