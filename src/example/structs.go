package example

/*
	结构相关
*/
import "fmt"

type A struct {
	A1 string
	A2 int
}

type B struct {
	A
	B1 string
}

type C struct {
	C1 string
	B
	A
}

func StructExample() {
	entity := C{}
	entity.B.A1 = "a11"
	entity.B.A2 = 22
	entity.A1 = "a1"
	entity.A2 = 2
	entity.B1 = "b1"
	entity.C1 = "c1"

	fmt.Printf("%v\n", entity)

	//匿名结构
	v := struct {
		V1 string
		V2 string
	}{V1: "v1", V2: "v2"}
	fmt.Println(v)
}
