package dynamic_programming

import (
	"fmt"
	"testing"
)

func TestCutRod(t *testing.T) {
	priceList := make(map[int]int)
	priceList[1] = 1
	priceList[2] = 5
	priceList[3] = 8
	priceList[4] = 9
	priceList[5] = 10
	priceList[6] = 17
	priceList[7] = 17
	priceList[8] = 20
	priceList[9] = 24
	priceList[10] = 30
	for i := 0; i < 20; i++ {
		rod := CutRod(i, priceList)
		fmt.Printf("长度为%d的钢条最多可以出售%d元。\n", i, rod)
	}
}
