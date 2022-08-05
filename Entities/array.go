package Entities

type ArrayStruct struct {
	Array []int
}

func (A *ArrayStruct) GetArray() []int {
	return A.Array
}

func (A *ArrayStruct) SetArray() {
	A.Array = []int{45, 23, 92, 12, 53}
}
