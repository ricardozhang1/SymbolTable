package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Key  int
	Val  interface{}
	Next *Node
	Prev *Node
}

type SymbolList struct {
	head *Node
	size int
}

func NewNode(key int, value interface{}) *Node {
	return &Node{Key: key, Val: value}
}

func NewSymbilList() *SymbolList {
	return &SymbolList{}
}

func (s *SymbolList) GetByKey(key int) (value interface{}, err error) {
	headNode := s.head
	for headNode != nil {
		if headNode.Key == key {
			value, err = headNode.Val, nil
			return
		}
		headNode = headNode.Next
	}
	value, err = -1, errors.New("can not find this key")
	return
}

func (s *SymbolList) Put(key int, value interface{}) {
	headNode := s.head
	for headNode != nil {
		if headNode.Key == key {
			headNode.Val = value
			return
		}
		headNode = headNode.Next
	}

	// 符号表中没有键为Key的键值对
	newNode := NewNode(key, value)
	if s.head == nil {
		s.head = newNode
	} else {
		newNode.Next = s.head
		s.head.Prev = newNode
		s.head = newNode
	}
	s.size++
}

func (s *SymbolList) Delete(key int) (bool, error) {
	headNode := s.head
	for headNode != nil {
		if headNode.Key == key {
			if headNode.Prev == nil {
				//删除第一个节点
				fmt.Println("11111", headNode.Val)
				s.head = headNode.Next
				s.size--
				return true, nil
			} else if headNode.Next == nil {
				//删除最后一个节点
				headNode.Prev.Next = nil
				s.size--
				return true, nil
			} else {
				//删除中间节点
				fmt.Println("22222", headNode.Val)
				fmt.Println(headNode.Prev.Val, headNode.Next.Next.Val)
				headNode.Prev.Next = headNode.Next
				headNode.Next.Prev = headNode.Prev
				s.size--
				return true, nil
			}
		}
		headNode = headNode.Next
	}
	return false, errors.New("do not exist key")
}

func (s *SymbolList) GetSize() int {
	return s.size
}

func (s *SymbolList) PrintList() {
	headNode := s.head
	for headNode != nil {
		fmt.Println(headNode.Val)
		headNode = headNode.Next
	}
}

func main() {
	sl := NewSymbilList()

	sl.Put(1, "aa")
	sl.Put(2, "bb")
	sl.Put(3, "cc")
	sl.Put(4, "dd")
	sl.Put(5, "ee")
	sl.Put(6, "ff")
	sl.Put(7, "gg")

	// sl.PrintList()
	// dd := sl.GetSize()
	// fmt.Println("Size: ", dd)

	// sl.Put(4, "xx")
	// sl.Put(2, "uu")
	sl.Put(100, "qqqqqq")
	sl.Put(101, "pppppp")
	sl.PrintList()

	// k4, v4 := sl.GetByKey(4)
	// fmt.Println("Get 4 value: ", k4, v4)
	// k6, v6 := sl.GetByKey(6)
	// fmt.Println("Get 6 value: ", k6, v6)
	// k99, v99 := sl.GetByKey(99)
	// fmt.Println("Get 99 value: ", k99, v99)

	// k3, v3 := sl.Delete(3)
	// fmt.Println("Delete 3: ", k3, v3)
	// fmt.Println("=========================")
	// sl.PrintList()

	_, _ = sl.Delete(4)
	_, _ = sl.Delete(5)
	_, _ = sl.Delete(6)
	_, _ = sl.Delete(7)

	// fmt.Println("=========================")
	// sl.PrintList()
	_, _ = sl.Delete(101)

	// fmt.Println("=============1111============")
	// sl.PrintList()
	_, _ = sl.Delete(1)

	fmt.Println("=========================")

	// _, _ = sl.Delete(2)
	sl.PrintList()

	
	

	// fmt.Println("=========================")
	// sl.PrintList()

}
