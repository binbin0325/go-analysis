package main

import (
	"fmt"
	"sync"
)

type Entry struct {
	Value string
}

func main() {
	m := sync.Map{}
	m.Store("1", Entry{Value: "1"})
	m.Store("2", Entry{Value: "2"})
	result, ok := m.Load("1")
	if ok {
		fmt.Println(result)
	}
	m.Load("1")
	m.Load("2")
	m.Delete("1")
	fmt.Println(m)
	m.LoadOrStore("3", Entry{Value: "3"})
}
