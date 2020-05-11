/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */

// @lc code=start

type DLinkedNode struct{
	key,val int
	pre,next *DLinkedNode
}
type LRUCache struct {
	cur, capacity int
	head,tail *DLinkedNode
	hash map[int]*DLinkedNode
}


func Constructor(capacity int) LRUCache {
	head,tail:=&DLinkedNode{},&DLinkedNode{}
	head.next = tail
	tail.pre = head

	return LRUCache{
		capacity:capacity,
		cur:0,
		head:head,
		tail:tail,
		hash:make(map[int]*DLinkedNode),
	}
}
func(this *LRUCache) addnode(n *DLinkedNode){
	n.next = this.head.next
	n.pre = this.head
	this.head.next.pre = n
	this.head.next = n
}
func(this *LRUCache) remove(n *DLinkedNode){
	n.pre.next = n.next
	n.next.pre = n.pre
}
func(this *LRUCache) movetohead(n *DLinkedNode){
	this.remove(n)
	this.addnode(n)
}

func (this *LRUCache) Get(key int) int {
	n,ok:= this.hash[key]
	if ok{
		this.movetohead(n)
		return n.val
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	n,ok:=this.hash[key]
	//存在,更新value并移动到头部
	if ok{
		n.val = value
		this.movetohead(n)
		return
	}
	
	n=&DLinkedNode{key:key,val:value}
	this.hash[key]=n
	//有容量,添加到头部
	if this.cur < this.capacity{
		this.addnode(n)
		this.cur++
		return
	}
	//没容量,删除节点并添加新节点到头部
	delete(this.hash,this.tail.pre.key)
	this.remove(this.tail.pre)
	this.addnode(n)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

