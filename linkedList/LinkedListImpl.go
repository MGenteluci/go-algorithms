package main

import "fmt"

func (l *linkedList[T]) Add(value T) {
	node := &Node[T]{Value: value}
	temp := l.Head
	if temp == nil {
		l.AddToStart(value)
	} else {
		for {
			if temp.Next == nil {
				temp.Next = node
				l.increaseLen()
				break
			}
			temp = temp.Next
		}
	}
}

func (l *linkedList[T]) AddToStart(value T) {
	node := &Node[T]{Value: value}

	if l.Head == nil {
		l.Head = node
		return
	} else {
		node.Next = l.Head
		l.Head = node
	}
	l.increaseLen()
}

func (l *linkedList[T]) Insert(value T, index int) {
	newNode := &Node[T]{Value: value}
	temp := l.Head
	for i := 0; i < index; i++ {
		if i == index - 1 {
			newNode.Next = temp.Next
			temp.Next = newNode
			break
		}
		temp = temp.Next
	}
	l.increaseLen()
}

func (l *linkedList[T]) Len() int {
	return l.ListLen
}

func (l *linkedList[T]) increaseLen() {
	l.ListLen += 1
}

func (l *linkedList[T]) decreaseLen() {
	l.ListLen -= 1
}

func (l *linkedList[T]) RemoveFirst() {
	l.Head = l.Head.Next
	l.decreaseLen()
}

func (l *linkedList[T]) RemoveLast() {
	l.removeLast(l.Head)
	l.decreaseLen()
}

func (l *linkedList[T]) RemoveByIndex(index int) {
	if index == 0 {
		l.RemoveFirst()
		return
	}

	temp := l.Head
	for i := 0; i < index - 1; i++ {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
	l.decreaseLen()
}

func (l *linkedList[T]) removeLast(node *Node[T]) {
	if node.Next.Next == nil {
		node.Next = nil
		return
	}
	
	l.removeLast(node.Next)
}

func (l *linkedList[T]) Clear() {
	l.Head = nil
	l.ListLen = 0
}

func (l *linkedList[T]) Find(index int) *T {
	if l.Head == nil {
		return nil
	}

	node := l.Head
	for i := 0; i < index; i++ {
		if node.Next == nil {
			return nil
		}

		node = node.Next
	}

	return &node.Value
}

func (l *linkedList[T]) PrettyPrint() {
	temp := l.Head
	if temp == nil {
		return
	}

	prettyPrint := ""
	for {
		if temp.Next != nil {
			prettyPrint = prettyPrint + fmt.Sprintf("%v -> ", temp.Value)
			temp = temp.Next
			continue
		}
		
		prettyPrint = prettyPrint + fmt.Sprintf("%v", temp.Value)
		break
	}
	fmt.Println(prettyPrint)
}

func (l *linkedList[T]) AsArray() []T {
	temp := l.Head
	if temp == nil {
		return nil
	}

	res := []T{}
	for {
		if temp.Next != nil {
			res = append(res, temp.Value)
			temp = temp.Next
			continue
		}
		
		res = append(res, temp.Value)
		break
	}

	return res
}
