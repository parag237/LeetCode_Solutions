package main

import (
	"Practice/code/LinkedList"
	"Practice/utils"
	"fmt"
)

func main() {
	// Create a sample palindrome linked list [1,2,2,1]
	palindrome := utils.CreateLinkedList([]int{1, 2, 2, 1})

	// Create a sample non-palindrome linked list [1,2,3,4]
	nonPalindrome := utils.CreateLinkedList([]int{1, 2, 3, 4})

	// Call the IsPalindrome method and print results
	fmt.Println("Is palindrome linked list a palindrome?", LinkedList.IsPalindrome(palindrome))
	fmt.Println("Is non-palindrome linked list a palindrome?", LinkedList.IsPalindrome(nonPalindrome))
}
