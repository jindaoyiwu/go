package main

import "fmt"

type Node struct {
	Val  interface{}
	Next *Node
}

type List struct {
	headNode *Node
}

func (l *List) IsEmpty() bool {
	if (l.headNode != nil) {
		return false
	} else {
		return true
	}
}

func (l *List) Add(data interface{}) {
	node := new(Node)
	node.Val = data
	node.Next = l.headNode
	l.headNode = node
	return
}

func (l *List) Append(data interface{}) {
	node := new(Node)
	node.Val = data
	if (l.IsEmpty() == true) {
		l.headNode = node
	} else {
		cur := l.headNode
		//找到最后的节点
		for cur.Next != nil{
			cur = cur.Next
		}
		cur.Next = node
	}
}

func (l *List) ShowList() {
	if !l.IsEmpty() {
		cur := l.headNode
		for cur.Next != nil{
			fmt.Println(cur.Val)
			cur = cur.Next
		}
		fmt.Println(cur.Val)
	}
}

func (l *List) Reverse() *List{
	ll := new(List)

	if l.IsEmpty(){
		return &List{}
	} else {
		cur := l.headNode
		for cur.Next != nil{
			ll.Add(cur.Val)
			cur = cur.Next
		}
		ll.Add(cur.Val)
	}
	return ll
}


func main() {
	list := List{}
	list.Add(2)
	list.Add(1)
	list.Append(3)
	list.Append(4)
	list.ShowList()

	newlist := list.Reverse()
	newlist.ShowList()
}
