package main

import "fmt"

type Data interface {
}

type Node struct {
	Data Data
	Next *Node
}

type List struct {
	headNode *Node
}

//从头部添加节点
func (this *List) add(data Data) {
	//先申请一个节点
	node := Node{
		Data: data,
		Next: nil,
	}
	//这个节点指向链表的头部
	node.Next = this.headNode
	//链表的头部重新指向这个节点
	this.headNode = &node
}

//从尾部添加一个节点
func (this *List) append(data Data) {
	//先申请一个节点
	node := Node{
		Data: data,
		Next: nil,
	}
	if this.IsEmpty() {
		this.headNode = &node
	} else {
		cur := this.headNode
		for cur.Next != nil{
			cur = cur.Next
		}
		cur.Next = &node
	}
}

//获取链表的长度
func (this *List) length() int {
	cur := this.headNode
	count := 0
	for cur.Next != nil  {
		cur = cur.Next
		count++
	}
	return count
}

//判断链表是否为空
func (this *List) IsEmpty() bool {
	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

//遍历链表
func (this *List) ShowList() {
	if !this.IsEmpty(){
		cur := this.headNode
		for {
			fmt.Println(cur.Data)
			if cur.Next != nil{
				cur = cur.Next
			} else {
				break
			}
		}
	}
}

func (this *List) insert(index int, data Data) {
	
}

func main() {
	list := List{}
	list.add(1)
	list.add(2)
	list.add(4)
	list.add(5)
	//从尾部添加一个
	list.append(23)
	list.append(24)
	list.ShowList()

	fmt.Println(list.length())

}
