package main

import (
	"fmt"
	"math"
	"strings"
)

// Code Challenges in Go
// This file contains several programming challenges to practice Go skills.
// Each challenge is described in comments. Try to implement the functions below.

// Challenge 1: Reverse a String
// Write a function that takes a string as input and returns the string reversed.
// Example: reverseString("hello") should return "olleh"
//
//	func reverseString(s string) string {
//	    // Your code here
//	}
func ReverseString(s string) string {
	var result []string
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, string(s[i]))
	}
	return strings.Join(result, "")
}

// Challenge 2: Factorial
// Write a recursive function to calculate the factorial of a number.
// Example: factorial(5) should return 120
//
//	func factorial(n int) int {
//	    // Your code here
//	}
func Factorial(n int) int {
	result := 1
	for i := 1; i < n; i++ {
		result *= i
	}
	return result
}

// Challenge 3: Check if Prime
// Write a function that checks if a number is prime.
// A prime number is greater than 1 and has no positive divisors other than 1 and itself.
// Example: isPrime(7) should return true, isPrime(4) should return false
//
//	func isPrime(n int) bool {
//	    // Your code here
//	}
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Challenge 4: Fibonacci Sequence
// Write a function to return the nth Fibonacci number.
// The Fibonacci sequence: 0, 1, 1, 2, 3, 5, 8, 13, ...
// Example: fibonacci(6) should return 8
//
//	func fibonacci(n int) int {
//	    // Your code here
//	}
func fibonacci(n int) int {
	// if n <= 0 {
	// 	return 0
	// }
	// if n == 1 {
	// 	return 1
	// }
	// return fibonacci(n-1) + fibonacci(n-2)
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b

}

// Challenge 5: Count Vowels
// Write a function that counts the number of vowels (a, e, i, o, u) in a string.
// Case insensitive.
// Example: countVowels("Hello World") should return 3
//
//	func countVowels(s string) int {
//	    // Your code here
//	}
func countVowels(s string) string {
	var result []string
	vowel := "aeiouAEIOU"

	for _, char := range s {
		if strings.ContainsAny(vowel, string(char)) {
			result = append(result, string(char))
		}
	}

	// for _, char := range s {
	// 	switch char {
	// 	case 'a','e','i','o','u','A','E','I','O','U':
	// 		count++
	// 	}
	// }
	return strings.Join(result, " ")
}

// Challenge 6: Find Maximum in Array
// Write a function that finds the maximum value in a slice of integers.
// Example: findMax([]int{1, 3, 2, 5, 4}) should return 5
//
//	func findMax(nums []int) int {
//	    // Your code here
//	}
func findMax(nums []int) int {
	b := nums[0]
	for _, num := range nums {
		if num > b {
			b = num
		}
	}
	return b
}

// Challenge 7: Check Palindrome
// Write a function that checks if a string is a palindrome (reads the same forwards and backwards).
// Ignore case and non-alphanumeric characters.
// Example: isPalindrome("A man, a plan, a canal: Panama") should return true
//
//	func isPalindrome(s string) bool {
//	    // Your code here
//	}
func alphaNumeric(r byte) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isPalindrom(s string)bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !alphaNumeric(s[left]) {
			left++
		}
		for left < right && !alphaNumeric(s[right]) {
			right--
		}
		if left < right && strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}
		left--
		right++
	}
	return true
}



// Challenge 8: Binary Search
// Implement binary search on a sorted slice of integers.
// Return the index if found, -1 if not.
// Example: binarySearch([]int{1, 2, 3, 4, 5}, 3) should return 2
// func binarySearch(nums []int, target int) int {
//     // Your code here
// }

// Challenge 9: Sum of Digits
// Write a function that calculates the sum of digits of a number.
// Example: sumOfDigits(123) should return 6
// func sumOfDigits(n int) int {
//     // Your code here
// }

// Challenge 10: Merge Two Sorted Arrays
// Write a function that merges two sorted slices into one sorted slice.
// Example: merge([]int{1, 3, 5}, []int{2, 4, 6}) should return []int{1, 2, 3, 4, 5, 6}
// func merge(a, b []int) []int {
//     // Your code here
// }

func ReturnIsPrime(p []int) []int {
	var result []int

	for _, num := range p {
		if isPrime(num) {
			result = append(result, num)
		}
	}
	return result

}

func main() {
	fmt.Println("=== Testing Challenges ===\n")

	// Test 1: Reverse String
	fmt.Println("Test 1: Reverse String")
	str := "hello"
	fmt.Printf("Input: %s\n", str)
	fmt.Printf("Output: %s\n\n", ReverseString(str))

	// Test 2: Factorial
	fmt.Println("Test 2: Factorial")
	num := 5
	fmt.Printf("Input: %d\n", num)
	fmt.Printf("Output: %d\n\n", Factorial(num))

	// Test 3: Check Prime
	fmt.Println("Test 3: Check if Prime")
	primeTest := 7
	fmt.Printf("Is %d prime? %v\n", primeTest, isPrime(primeTest))
	fmt.Printf("Is 4 prime? %v\n\n", isPrime(4))

	// Test 4: Fibonacci
	fmt.Println("Test 4: Fibonacci Sequence")
	fmt.Printf("6th Fibonacci number: %d\n", fibonacci(6))
	fmt.Printf("Fibonacci sequence (1-10): ")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println("\n")

	// Test 5: Filter Primes from List
	fmt.Println("Test 5: Filter Primes from List")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Input: %v\n", numbers)
	fmt.Printf("Output: %v\n", ReturnIsPrime(numbers))
	// Test 6: Countn Vowels
	fmt.Println("\nTest 6: Count Vowels")
	vowels := "Hello World  am in my room"
	fmt.Printf("Input: %s\n", vowels)
	fmt.Printf("Output: %s\n", countVowels(vowels))

	// Challenge 6: Find Maximum in Array
	fmt.Println(" Challenge 6: Find Maximum in Array")
	fmt.Printf("Input: %v\n", []int{1, 23, 4, 7, 45})
	fmt.Printf("output: %d\n", findMax([]int{1, 23, 4, 7, 45}))

	// // Challenge 7: Check Palindrome
	fmt.Println("Challenge 7: Check Palindrome")
	fmt.Printf("Input: %s\n", "A man, a plan, a canal: Panama")
	fmt.Printf("output: %v\n", isPalindrome("A man, a plan, a canal: Panama"))

}
