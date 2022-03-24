package main

import (
	"fmt"
	"palindrome/palindrome"
)

func main() {
	//result := "A"
	///palindrome.IsPalindrome(result)
	fmt.Println(palindrome.IsPalindrome("AAAB"))
	fmt.Println(palindrome.IsPalindrome(10001))
}
