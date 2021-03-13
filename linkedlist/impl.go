package main

import "fmt"

func main() {

	// Init the LinkedList
	linkedList := LinkedList{Head: &Node{
		Val:  0,
		Next: nil,
	}}

	// Insert at End
	linkedList.InsertAtEnd(&Node{
		Val:  1,
		Next: nil,
	})
	linkedList.InsertAtEnd(&Node{
		Val:  3,
		Next: nil,
	})
	fmt.Println("After Inserting at End")
	linkedList.Print()

	// Insert at Beginning
	linkedList.InsertAtBeginning(&Node{
		Val:  -1,
		Next: nil,
	})
	fmt.Println("After Inserting at Beginning")
	linkedList.Print()

	// Insert at Position
	linkedList.InsertAtPos(&Node{
		Val:  2,
		Next: nil,
	}, 2)
	linkedList.InsertAtPos(&Node{
		Val:  4,
		Next: nil,
	}, 4)
	fmt.Println("After Inserting at Pos")
	linkedList.Print()

	// Delete at End
	linkedList.DeleteAtEnd()
	fmt.Println("After Deleting at End")
	linkedList.Print()

	// Delete at Beginning
	linkedList.DeleteAtBeginning()
	fmt.Println("After Deleting at Beginning")
	linkedList.Print()

	// Delete at Pos
	linkedList.DeleteAtPos(1)
	fmt.Println("After Deleting at Pos")
	linkedList.Print()

	// Search for a Node
	res := linkedList.Search(1)
	fmt.Print("Searching for Value 1 :- ")
	fmt.Println(res)
	res = linkedList.Search(2)
	fmt.Print("Searching for Value 3 :- ")
	fmt.Println(res)

}
