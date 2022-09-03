package main

import (
	"fmt"
)

type Hoge struct{ value int }

func (h Hoge) HogeA(v int) {
	h.value = v + 5
}

func (h *Hoge) HogeB(v int) {
	h.value = v + 5
}

func main() {
	var h Hoge
	h.HogeA(1)
	fmt.Printf("the output of HogeA(not pointer type):%v\n", h.value)

	h.HogeB(0)
	fmt.Printf("the output of HogeB(pointer type):%v\n", h.value)
}
