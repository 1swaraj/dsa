package main

func main() {
	tr := Tree{Head: &Node{
		Left:   nil,
		Right:  nil,
		Key:    10,
		Height: 1,
	}}
	head := tr.Head
	head = Insert(head,10)
	head = Insert(head,20)
	head = Insert(head,30)
	head = Insert(head,40)
	head = Insert(head,50)
	head = Insert(head,25)
	preOrder(head)
}
