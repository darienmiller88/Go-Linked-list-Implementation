package main

import (
	"LinkedList/List"
	"fmt"
)

func main(){
	list := List.LinkedList{}

	list.PushBack(5)
	list.PushBack(6)
	list.PushBack(7)
	list.PushFront(26)

	fmt.Println("list len:", list.Length())
	list.PrintList()
	backValue, err := list.Back()

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("last value:", backValue)
	}

	frontValue, err := list.Front()

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("first value:", frontValue)
	}

	list.Delete(6)
	list.PrintList()
	fmt.Println("list len:", list.Length())

	result := list.Find(7)

	if !result{
		fmt.Println(7, "was not found")
	}else{
		fmt.Println(7, "was found")
	}

	result = list.Find(44)

	if !result{
		fmt.Println(44, "was not found")
	}else{
		fmt.Println(44, "was found")
	}

	list.PopBack()
	list.PrintList()

	list.PopFront()
	list.PrintList()

	list.Delete(5)
	list.PrintList()
}