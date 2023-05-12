package main

import (
	"fmt"
)

func main() {
	list := NewLinkedList[int]()
	fmt.Printf("Len: %d\n", list.Len())
	list.PrettyPrint()

	list.Add(2)
	list.PrettyPrint()
	list.Add(3)
	list.PrettyPrint()
	fmt.Printf("Len: %d\n", list.Len())

	list.Add(4)
	list.PrettyPrint()
	fmt.Printf("Len: %d\n", list.Len())

	list.RemoveLast()
	list.PrettyPrint()

	list.RemoveFirst()
	list.PrettyPrint()

	list.AddToStart(7)
	list.PrettyPrint()
	fmt.Printf("Len: %d\n", list.Len())

	fmt.Println()
	fmt.Printf("AsArray:: %v\n", list.AsArray())

	list.Clear()
	list.Add(7)
	list.Add(9)
	list.Add(11)
	list.Add(15)
	list.Add(22)
	list.PrettyPrint()
	list.RemoveByIndex(2)
	list.RemoveByIndex(0)
	list.RemoveByIndex(list.Len() - 1)
	list.PrettyPrint()

	fmt.Printf("Item at index 0: %v\n", *list.Find(0))
	fmt.Printf("Item at index 1: %v\n", *list.Find(1))

	list.Insert(58, 1)
	list.PrettyPrint()
	list.Insert(62, 2)
	list.PrettyPrint()
	// list.RemoveByValue(3)
	// fmt.Printf("Len: %d\n", list.Len())
}
