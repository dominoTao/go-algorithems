package main

// 加减优先级高于位移运算符
const (
	IntSize = 32 << (^uint(0) >> 63)
	IntMin = -1 << (IntSize - 1)
	IntMax = 1 << (IntSize - 1) - 1

	int8Min = -1 << 7
	int8Max = 1 << 7 - 1

	uintMax = 1 << (IntSize - 1)
	uint8Max = 1 << 8 - 1
)

//func main() {
//	fmt.Println(IntSize)
//	fmt.Println(IntMin)
//	fmt.Println(IntMax)
//	fmt.Println(int8Min)
//	fmt.Println(int8Max)
//	//fmt.Println(uintMax)
//	fmt.Println(uint8Max)
//	//64
//	//-9223372036854775808
//	//9223372036854775807
//	//-128
//	//127
//	//255
//}
