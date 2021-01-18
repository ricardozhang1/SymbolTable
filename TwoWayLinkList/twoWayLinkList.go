package main

import (
	"fmt"
)

//链表的一个节点
type ListNode struct {
	prev *ListNode
	next *ListNode
	value interface{}
}

//创建一个节点
func NewListNode(value interface{}) (listNode *ListNode) {
	listNode = &ListNode{
		// prev: nil,
		// next: nil,
		value: value,
	}
	return
}

//当前节点的下一个节点
func (l *ListNode) NextNode() (next *ListNode) {
	next = l.next
	return
}

//当前节点的前一个节点
func (l *ListNode) PrevNode() (prev *ListNode) {
	prev = l.prev
	return
}

//获取节点的值
func (l *ListNode) GetValue() (value interface{}) {
	if l == nil {
		return
	}
	value = l.value
	return
}

//定义一个链表
type LinkList struct{
	head *ListNode  // 表头节点
	tail *ListNode  // 表尾节点
	len int  // 链表的长度
}

//创建一个空链表
func NewList() (list *LinkList) {
	list = &LinkList{
		head: nil,
		tail: nil,
		len: 0,
	}
	return
}

//返回链表头结点
func (nl *LinkList) GetHead() (head *ListNode) {
	head = nl.head
	return
}

//返回链表尾结点
func (nl *LinkList) GetTail() (tail *ListNode) {
	tail = nl.tail
	return
}

//返回链表长度
func (nl *LinkList) Length() (length int) {
	length = nl.len
	return
}

//在链表的右边插入一个元素
func (nl *LinkList) RPush(value interface{}) {
	node := NewListNode(value)

	//链表为空的时候
	if nl.Length() == 0 {
		nl.head = node
		nl.tail = node
	} else {
		tail := nl.tail
		tail.next = node
		node.prev = tail

		nl.tail = node
	}
	nl.len++
	return
}

//在链表头部插入一个元素
func (nl *LinkList) LPush(value interface{}) {
	node := NewListNode(value)

	//当链表为空时
	if nl.Length() == 0 {
		nl.head = node
		nl.tail = node
	} else {
		nl.head.prev = node
		node.next = nl.head
		nl.head = node
	}
	nl.len++
	return
}

//从链表左边取出一个元素
func (nl *LinkList) LPop() (node *ListNode) {
	//数据为空
	if nl.len == 0 {
		return
	}

	node = nl.head

	if node.next == nil {
		//链表未空
		nl.head = nil
		nl.tail = nil
	} else {
		nl.head = node.next
		node.next = nil
		nl.head.prev = nil
	}
	nl.len--
	return
}

//丛链表右边取出一个元素
func (nl *LinkList) RPop() (node *ListNode) {
	//数据为空的时候
	if nl.len == 0 {
		return nil
	}
	node = nl.tail
	
	if node.prev == nil {
		nl.head = nil
		nl.tail = nil
	} else {
		nl.tail = node.prev
		nl.tail.next = nil
		node.prev = nil
	}
	nl.len--
	return
}

//向索引i出插入元素
func (nl *LinkList) InsertByIndex(index int, value interface{}) {
	if index < 0 || index > nl.len {
		return
	}

	node := NewListNode(value)
	currentNode := nl.SelectByIndex(index)

	currentNode.next.prev = node
	node.next = currentNode.next
	currentNode.next = node
	node.prev = currentNode
	nl.len++
	return
}

//删除索引i处的节点
func (nl *LinkList) DeleteByIndex(index int) {
	if index < 0 || index > nl.Length() {
		return
	}

	currentNode := nl.SelectByIndex(index)
	prevNode := currentNode.prev
	nextNode := currentNode.next

	prevNode.next = nextNode
	nextNode.prev = prevNode
	nl.len--
	return
}

//通过索引查找节点
//查不到节点则返回空nill
func (nl *LinkList) SelectByIndex(index int) (node *ListNode) {
	if index > nl.Length() {
		return
	}
	//索引为负数则表尾开始查找
	if index < 0 {
		index = (-index) - 1
		node = nl.tail
		for {
			if node == nil {
				return
			}

			if index == 0 {
				return
			}

			node = node.prev
			index--
		}
	} else {
		node = nl.head
		for ; index > 0 && node != nil; index-- {
			node = node.next
		}
	}
	return
}

//返回指定区间的元素
func (nl *LinkList) Range(start, stop int) (nodes []*ListNode) {
	nodes = make([]*ListNode, 0)

	//转为自然数
	if start > stop || start < 0 || stop > nl.len {
		return nil
	}

	//区间内元素个数
	rangeLen := stop - start + 1
	startNode := nl.SelectByIndex(start)
	for i:=0; i<rangeLen; i++ {
		if startNode == nil {
			break
		}

		nodes = append(nodes, startNode)
		startNode = startNode.next
	}
	return
}

//链表遍历
func (nl *LinkList) Print() {
	fmt.Println("TwoWayLinkList size: ", nl.Length())
	if nl.len == 0 {
		return
	}

	node:=nl.head
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

func main() {
	sl := NewList()

	for i:=0; i<6; i++ {
		sl.RPush(i)
	}

	sl.Print()
	fmt.Println("============================================")

	node := sl.LPop()
	fmt.Println("左侧元素1: ", node.value)
	node = sl.LPop()
	fmt.Println("左侧元素2: ", node.value)
	node = sl.LPop()
	fmt.Println("左侧元素3: ", node.value)

	sl.Print()
	fmt.Println("============================================")

	v := sl.SelectByIndex(1)
	fmt.Println(v.value)
	v = sl.SelectByIndex(2)
	fmt.Println(v.value)
	fmt.Println("============================================")

	sl.LPush(200)
	sl.LPush(100)
	sl.Print()
	fmt.Println("============================================")

	nodeList := sl.Range(0, 3)
	for _, v := range nodeList {
		fmt.Println(v.value)
	}
	fmt.Println("============================================")

	// node = sl.RPop()
	// fmt.Println("右侧元素1: ", node.value)
	// node = sl.RPop()
	// fmt.Println("右侧元素2: ", node.value)
	// node = sl.RPop()
	// fmt.Println("右侧元素3: ", node.value)
	// node = sl.RPop()
	// fmt.Println("右侧元素4: ", node.value)
	// node = sl.RPop()
	// fmt.Println("右侧元素5: ", node.value)

	// sl.Print()
	// fmt.Println("============================================")
	ret := sl.SelectByIndex(3)
	fmt.Println(ret.value)


	sl.InsertByIndex(3, 888)
	sl.InsertByIndex(2, 666)
	sl.Print()
	fmt.Println("============================================")

	sl.DeleteByIndex(1)
	sl.Print()
	fmt.Println("============================================")
}



