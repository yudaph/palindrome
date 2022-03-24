package palindrome

import (
	"strings"

	"golang.org/x/exp/constraints"
)

type signedOrString interface{
	constraints.Signed | string
}

func IsPalindrome[input signedOrString](firstInput input) bool {
	return func(toInterface interface{}) bool {
		var sliceInt []int
		switch newType := toInterface.(type){
			case string:
				return palindrome(strings.Split(newType, ""))
			case int:
				sliceInt = sliceInteger(newType)
			case int8:
				sliceInt = sliceInteger(newType)
			case int16:
				sliceInt = sliceInteger(newType)
			case int32:
				sliceInt = sliceInteger(newType)
			case int64:
				sliceInt = sliceInteger(newType)
		}
		return palindrome(sliceInt)
	}(firstInput)
}

func sliceInteger[signed constraints.Signed](input signed) (slice []int) {
	for input > 0 {
		slice = append(slice, int(input%10))
		input/=10
	}
	return
}

func palindrome[input signedOrString](sliceInput []input) bool {
	length := len(sliceInput)
	for i:=0;i<length/2;i++{
		if sliceInput[i] != sliceInput[length-1-i] {
			return false
		}
	}
	return true
}
