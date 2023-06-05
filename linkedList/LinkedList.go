package linkedList

type Node[T interface{}] struct {
	Value T
	Next  *Node[T]
}

type linkedList[T interface{}] struct {
	Head    *Node[T]
	ListLen int
}

type LinkedList[T interface{}] interface {
	Add(value T)
	AddToStart(value T)
	Insert(value T, index int)
	Len() int
	RemoveFirst()
	RemoveLast()
	RemoveByIndex(index int)
	Clear()
	Find(index int) *T
	PrettyPrint()
	AsArray() []T
}

func NewLinkedList[T interface{}]() LinkedList[T] {
	return &linkedList[T]{}
}
