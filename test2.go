package main

import (
	"fmt"
	"math"
)

func s(r float64) float64 {
	var s float64
	s = r * r * math.Pi
	return s
}
func main() {
	var r = 10.0
	result := s(r)      // 计算 9 的平方根
	fmt.Println(result) // 输出：3
}
