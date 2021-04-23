package main

import "fmt"

type node struct {
	key  int
	val  int
	pre  *node
	next *node
}

type LRUCache struct {
	head *node
	tail *node
	c    int
	data map[int]*node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{data: make(map[int]*node, capacity), c: capacity}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.data[key]
	if !ok {
		return -1
	} else if v != this.head { // 移动到头结点
		v.pre.next = v.next
		if v != this.tail {
			v.next.pre = v.pre
		} else {
			this.tail = v.pre
		}
		v.next = this.head
		this.head.pre = v
		v.pre = nil
		this.head = v
	}
	return v.val
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.data[key]
	if !ok { // 如果这个值不存在
		tmp := &node{
			key: key,
			val: value,
		}
		if len(this.data) == 0 { // 如果map为空，链表为空
			this.head = tmp
			this.tail = tmp
			this.data[key] = tmp
		} else if len(this.data) == this.c { // 如果map满了
			this.del()
			this.Put(key, value) // 删除元素后，再次添加进来
		} else {
			tmp.next = this.head
			this.head.pre = tmp
			this.head = tmp
			this.data[key] = tmp
		}
	} else {
		v.val = value
		if v != this.head { // 移动到头结点
			v.pre.next = v.next
			if v != this.tail {
				v.next.pre = v.pre
			} else {
				this.tail = v.pre
			}
			v.next = this.head
			this.head.pre = v
			v.pre = nil
			this.head = v
		}
	}
}

func (this *LRUCache) del() {
	delete(this.data, this.tail.key)
	if this.tail.pre != nil { // 删除map中的元素后，同时删除链表尾节点
		pre := this.tail.pre
		this.tail.pre = nil
		pre.next = nil
		this.tail = pre
	} else {
		this.head = nil
		this.tail = nil
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**
可以使用map + link，来解决此题
map[key]*Node
如果一个变量被缓存，判断map当中有没有该变量缓存，如果有，更新结果，并将这个节点移动到链表头部
如果map当中不存在，分两种情况：
1. map容量满了，删除链表尾节点，并在map当中删除尾节点的值，将新值存入map，放入链表头部
2. map没有满，则直接存入map，插入到链表头部

如果读取一个变量，判断map当中有没有，有的话， 直接读取该值，并将节点移动到链表头部，返回该值
*/

func main() {
	l := Constructor(2)
	l.Put(2, 1)
	l.Put(1, 1)
	l.Put(2, 3)
	l.Put(4, 3)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(2))
}
