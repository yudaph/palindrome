package main

import (
	"fmt"
	"palindrome/palindrome"
)

func main() {
	fmt.Println(palindrome.IsPalindrome("AAAB"))
	fmt.Println(palindrome.IsPalindrome("ABABA"))
	fmt.Println(palindrome.IsPalindrome(10001))
	fmt.Println(palindrome.IsPalindrome(10110))
}
