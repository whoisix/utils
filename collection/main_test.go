package main

import (
	"log"
	"testing"
)

type Foo struct {
	A string
	B int
	C int32
	D uint32
}

var c *collection

func TestMain(m *testing.M) {
	a1 := Foo{A: "a1", B: 3, C: 13, D: 14}
	a2 := Foo{A: "a2", B: 4, C: 23, D: 24}
	a3 := Foo{A: "a3", B: 5, C: 33, D: 34}
	foo := []Foo{a1, a3, a2}
	c = Collect(foo)
	m.Run()
}

func Test_collection(t *testing.T) {
	log.Println("==========slice===========")
	log.Println("B", c.Pluck("B").SliceInt())
	log.Println("C", c.Pluck("C").SliceInt32())
	log.Println("D", c.Pluck("D").SliceUint32())

	log.Println("==========map===========")
	log.Println("B", c.KeyByInt("B"))
	log.Println("C", c.KeyByInt32("C"))
	log.Println("D", c.KeyByUint32("D"))
}
