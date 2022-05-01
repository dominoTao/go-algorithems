package main

import (
	"fmt"
	"github.com/uber/h3-go"
	//"github.com/uber/h3-go/v3"
	"testing"
)

func TestH3(t *testing.T) {
	exampleFromGeo()
}

func exampleFromGeo() {

	coord := h3.GeoCoord{
		Longitude: 116.397128,
		Latitude:  39.916527,
	}
	a := h3.FromGeo(coord, 8)
	fmt.Printf("h3.FromGeo(coord, 8)  %T， 字面量%v， 十六进制%x\n", a, a, a)
	//distances := h3.KRingDistances(    0x8928308280fffff, 8)
	distances := h3.KRingDistances(613363266611052543, 8)
	fmt.Printf("%T\n", distances)
	for i, distance := range distances {
		fmt.Printf("index:%d, type:%T, value:%v\n", i,distance, distance)
		for i2, index := range distance {
			fmt.Printf("\tindex:%d, type:%T, value:%v\n", i2,index, index)
		}
	}
	geo := h3.GeoCoord{
		Latitude:  37.775938728915946,
		Longitude: -122.41795063018799,
	}
	resolution := 9
	fmt.Printf("%#x\n", h3.FromGeo(geo, resolution))
	// Output:
	// 0x8928308280fffff
}
