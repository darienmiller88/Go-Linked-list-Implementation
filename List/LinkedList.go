package List

import (
	"errors"
	"fmt"
)

type node struct {
	data     int
	next     *node
	previous *node
}

type LinkedList struct {
	len  int
	head *node
	tail *node
}

func (l *LinkedList) PrintList() {
	fmt.Print("[")
	for head := l.head; head != nil; head = head.next {
		fmt.Print(head.data, " -> ")
	}
	fmt.Println("]")
}

func (l *LinkedList) PrintListReverse() {
	fmt.Print("[")
	for tail := l.tail; tail != nil; tail = tail.previous {
		fmt.Print(tail.data, " -> ")
	}
	fmt.Println("]")
}

//PushFront will add a new node to the front of the list in O(1) time.
func (l *LinkedList) PushFront(number int) {
	if l.len == 0 {
		l.head = new(node)
		l.tail = l.head
		l.tail.next     = nil
		l.tail.previous = nil
		l.tail.data     = number
	} else {
		//First, create a new node that will become the new first node.
		newNode := new(node)

		//Next, set the data of the node to the number the user set.
		newNode.data = number

		//Afterwards, connect the next pointer of the new node to the node the head
		newNode.next = l.head

		//Set the previous to nil
		newNode.previous = nil

		//Set the previous pointer of the heads node to the new node
		l.head.previous = newNode

		//Finally, set the head pointer to point at the new node.
		l.head = newNode
	}

	l.len++
}

//PushBack will add a new node to the back of the list in O(1) time.
func (l *LinkedList) PushBack(number int) {
	if l.len == 0 {
		l.tail = new(node)
		l.head = l.tail
		l.tail.data = number
		l.tail.next = nil
	} else {
		//First, set the tail's next to point to a new node
		l.tail.next = new(node)

		//Second, set the previous of the new node to point at the current tail node
		l.tail.next.previous = l.tail

		//Next, set the tail to point to the new
		l.tail = l.tail.next

		//Assign the node the number the user passed.
		l.tail.data = number

		//Finally, set the tail node to nil.
		l.tail.next = nil
	}

	l.len++
}

//Delete will remove the first instance of "number", or return an error if the list is empty.
func (l *LinkedList) Delete(number int) error{
	if l.len == 0{
		return errors.New("List is empty! Cannot delete")
	}else if l.head.data == number{
		l.PopFront()
	}else if l.tail.data == number{
		l.PopBack()
	}else if l.len == 1{
		l.head = nil
		l.tail = nil
	}else{
		tempHead := l.head

		for tempHead != nil {
			if tempHead.data == number {
				//In order to delete in the middle of the list, get a pointer to the node in front and behind
				//of the node that has the value the user is looking for.
				nodeBehind  := tempHead.previous
				nodeInFront := tempHead.next

				//Set the next of the pointer behind the found node to point to the one in front, and vice versa
				nodeBehind.next      = nodeInFront
				nodeInFront.previous = nodeBehind
				
				//Disconnect the found node from the nodes in front and behind it.
				tempHead.next     = nil
				tempHead.previous = nil
				break
			}
	
			tempHead = tempHead.next
		}
	}

	l.len--
	
	return nil
}

//PopBack will remove the first element in the in the list, and will return an error if the list is empty.
func (l *LinkedList) PopFront() error {
	if l.len == 0{
		return errors.New("List is empty! Cannot delete")
	}else{
		tempHead := l.head

		//Set the head to the node in front of it.
		l.head = l.head.next

		//If the head is now nil after moving it to the next node, also set the tail to nil
		if l.head == nil{
			l.tail = l.head
		}else{
			//otherwise, set the previous of the new first node to nil.
			l.head.previous = nil
		}
		
		//Finally, set the next to the old first node to nil to completely disengage it from the list.
		tempHead.next = nil
	}

	l.len--

	return nil
}

//PopBack will remove the last element in the in the list, and will return an error if the list is empty.
func (l *LinkedList) PopBack() error {
	if l.len == 0{
		return errors.New("List is empty! Cannot delete")
	}else{
		tempTail := l.tail

		//Set the tail to the node in behind of it.
		l.tail = l.tail.previous

		//If the tail is now nil after moving it to the next node, also set the head to nil
		if l.tail == nil{
			l.head = l.tail
		}else{
			//otherwise, set the next of the new last node to nil.
			l.tail.next = nil
		}
		
		//Finally, set the previous. to the old last node to nil to completely disengage it from the list.
		tempTail.previous = nil
	}

	l.len--
	
	return nil
}

//Returns the first value of the list, or an error if the list is empty
func (l *LinkedList) Front() (int, error){
	if l.head == nil{
		return 0, errors.New("No values in linked list.")
	}

	return l.head.data, nil
}

//Returns the last value of the list, or an error if the list is empty
func (l *LinkedList) Back() (int, error){
	if l.tail == nil{
		return 0, errors.New("No values in linked list.")
	}
	
	return l.tail.data, nil
}


//Find will search for the number in the linked list, and return true or false if it was found or not.
func (l *LinkedList) Find(number int) bool {
	for head := l.head; head != nil; head = head.next {
		if head.data == number{
			return true
		}
	}

	return false

}

//Returns the length of the linked list.
func (l *LinkedList) Length() int {
	return l.len
}
