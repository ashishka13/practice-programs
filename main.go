package main

import (
	"fmt"
	"log"
	"strings"
)

func PrintStringRecursion(s string) {
	log.Println(s)
	if s == "" {
		return
	}
	s = s[1:]               // skip first position and reassign to same string variable
	PrintStringRecursion(s) // send the trimmed variable to same function
}

func StringRevWithoutSplChars(strInput string) {
	processInput1 := strInput
	revChars := []string{}                        // store one one character
	placeHolders := make([]string, len(strInput)) // store only special characters at their original places

	for len(processInput1) > 0 {
		index := len(processInput1) - 1
		s := processInput1[index:]                            // get last character of a word
		if (s >= "a" && s <= "z") || (s >= "A" && s <= "Z") { // check if it is in a z range
			revChars = append(revChars, s)        // if normal character then add it in char array
			placeHolders[index] = ""              // and keep track of special array with assigning empty instead of normal character
			processInput1 = processInput1[:index] // reduce reversing string length
			continue                              // continue beacause we have found the normal char
		}
		placeHolders[index] = s               // special character at original place
		processInput1 = processInput1[:index] // reduce reversing string length
	}
	// at this stage we have 2 arrays. 1) normal letters reversed, 2) special letters without normal letters
	for key, val := range placeHolders {
		if val == "" { // start replacing empty spaces with normal characters
			placeHolders[key] = revChars[0] // append only the first letter of normal array
			revChars = revChars[1:]         // reduce array length by 1 only
		}
	}

	log.Println(strings.Split(strInput, ""))
	log.Println(placeHolders)
}

func PrintStringClosure(a string) func() bool {
	log.Println(a)

	return func() bool {
		if a == "" {
			log.Println("code was here")
			return false
		}
		a = a[1:]
		log.Println("inside func", a)
		return true
	}

	/*
			streps to use
			extract a value into a variable
		and just write it like: retVal()
		done, it will operate on the variable once.

		if you want a infinite for loop with breaking condition,
		then give a condition to inside closure func
		which will return false at some break point.

			retVal := PrintStringClosure("ajsjnoiasogjnasdgjan")
		retVal()
			for retVal() {
			}
	*/
}

func ReverseWordsInString(sentence string) {
	log.Println("ip ", sentence)

	wordsArray := strings.Fields(sentence) // crete array of every word in sentence
	wordsString := ""                      // create a var to store reversed words one by one
	for _, val := range wordsArray {       // iterate over sentence words array
		singleWord := val        // use this variable for operations, not the original one
		singleRevWord := ""      // store single reversed word
		reverseArr := []string{} // store single reversed letter in array

		for len(singleWord) > 0 { // word reversal code
			index := len(singleWord) - 1
			b := singleWord[index:]
			singleWord = singleWord[:index]
			reverseArr = append(reverseArr, b) // append reversed letter
		}

		for i := 0; i < len(reverseArr); i++ {
			singleRevWord = singleRevWord + reverseArr[i] // create single reversed word from reversed array
		}

		wordsString = wordsString + " " + singleRevWord
	} // for loop will run according to number of words
	log.Println("op", wordsString)
}

func PalindromeString(str string) {
	for i, j := len(str)-1, 0; i > 0; i, j = i-1, j+1 { // check equal distance letters. center letter will also get checked.
		ival := string(str[i])
		jval := string(str[j])

		if ival != jval {
			log.Println("not palindrome")
			return
		}
	}
	log.Println("palindrome present")
}

func PalindromeNumber(inputNum int) {
	num := inputNum
	palindrome := 0
	digit := 0

	if num > 9 { // single digit is always palindrome
		for num > 0 {
			digit = num % 10
			palindrome = (palindrome * 10) + digit
			num = num / 10
		}
		if inputNum == palindrome {
			log.Println("palindrome")
		} else {
			log.Println("not")
		}
	} else {
		log.Println("not")
	}
}

func PrimeOrNot(num int) {
	if num == 2 {
		log.Println("prime")
		return
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			log.Println("not prime")
			return
		}
	}
	log.Println("prime")
}

func ArmstrongNum(num int) {
	n := num
	digit := 0
	addition := 0

	for n > 0 {
		digit = n % 10
		n = n / 10
		addition = addition + (digit * digit * digit)
	}
	if addition == num {
		log.Println("armstrong number")
	} else {
		log.Println("not Armstrong")
	} // 153 Armstrong example. 1^3 + 5^3 + 3^3 = 153
}

func ReverseNumber(num int) {
	rev := 0
	digit := 0

	for num > 0 {
		digit = num % 10
		rev = (rev * 10) + digit
		num = num / 10
	}

	log.Println(rev)
}

func PrimeOrNotInRange(num1, num2 int) {
	if num1 < 2 || num2 < 2 {
		log.Println("Numbers must be greater than 2.")
		return
	}
	for num1 <= num2 {
		isPrime := true
		for i := 2; i <= num1/2; i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			log.Println(num1)
		}
		num1++
	}
}

//0 1 1 2 3 5 8 13 21 34 55 89 114
func Fibonacci(foboRange int) {
	prev, next, answer := 0, 1, 0

	log.Print(prev)
	log.Println(next)
	for i := 0; i < foboRange; i++ {
		answer = prev + next
		prev = next
		next = answer
		log.Print(answer)
	}
}

func FibonacciRecursive(fiborange int) {
	fibonacciRecursiveInner(fiborange, 0, 1, 0)
}

//0 1 1 2 3 5 8 13 21 34 55 89 114
func fibonacciRecursiveInner(fiboRange, prev, next, answer int) {
	_ = answer
	if fiboRange == 0 {
		return
	}
	answer = prev + next
	prev = next
	next = answer
	log.Print("answer", answer)
	fibonacciRecursiveInner(fiboRange-1, prev, next, answer)
}

func FindDupStringChars(str string) {
	strArr := strings.Split(str, "")
	dupMap := make(map[string]int) // use map to store unique values with its count
	for _, val := range strArr {
		if _, ok := dupMap[val]; !ok {
			dupMap[val] = 1
		} else {
			log.Println("duplicate emelent ", val)
		}
	}
}

func FindUniqueChars(str string) {
	strArr := strings.Split(str, "")
	dupMap := make(map[string]bool)
	for _, val := range strArr {
		if ok := dupMap[val]; !ok {
			log.Println("unique emelent ", val)
			dupMap[val] = true
		}
	}
}

func ReverseStringNormal(str string) {
	revStr := ""
	for len(str) > 0 { // basic reverse string code
		lastindex := len(str) - 1
		b := str[lastindex:] // locates last position of string array
		str = str[:lastindex]
		revStr = revStr + b
	}
	log.Println(revStr)
}

func BubbleSort(numbers []int) {
	// 6 5 7 1 3 4
	counter := 0
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] { // number swap code
				temp := numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = temp
			}
			counter++
		}
	}
	log.Println(numbers)
}

func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left, right, pivot := 0, len(nums)-1, len(nums)/2
	// left = leftmost num, right = rightmost num, pivot = central num

	nums[pivot], nums[right] = nums[right], nums[pivot] // swap pivote and right positions

	for i := range nums {
		if nums[i] < nums[right] { // check if i'th position is less than swapped rightmost
			nums[left], nums[i] = nums[i], nums[left] // swap left and i'th position
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left] // swap code

	QuickSort(nums[:left])
	QuickSort(nums[left+1:])

	return nums
}

// equal length and number of letters but order can be different
func AnagramString(str1, str2 string) {
	if len(str1) != len(str2) {
		log.Println("length not same not anagrams")
		return
	}
	strArr1 := strings.Split(str1, "")
	strArr2 := strings.Split(str2, "")
	counterflag := 0

	for i := 0; i < len(strArr1); i++ {
		for j := 0; j < len(strArr2); j++ {
			if strArr1[i] == strArr2[j] {
				counterflag++ // count to keep track of matched letters
				break         // if just single letter is found, break j loop and go to next i loop
			}
		}
	}
	if counterflag == len(str1) { // because equal number of chars and length
		log.Println("anagrams", counterflag)
		return
	}
	log.Println("not anagrams", counterflag)
}

func GetStringCombinations(str string) {
	strarr := strings.Split(str, "")
	for i := 0; i < len(strarr); i++ {
		for j := i + 1; j < len(strarr); j++ {
			if strarr[i] != strarr[j] {
				temp := strarr[i]
				strarr[i] = strarr[j]
				strarr[j] = temp
				a := ""
				for _, val := range strarr {
					a = a + val
				}
				log.Println(a)
			}
		}
	}
}

func ShuffleNumbers(num []int) {
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i] != num[j] { // use swap logic to swap numbers mutiple times in for loop
				temp := num[i]
				num[i] = num[j]
				num[j] = temp
				log.Println(num)
			}
		}
	}
}

func MaxOfThreeNum(num1, num2, num3 int) {
	max := num1     // assume 1st num is max.
	if max < num2 { // if not then swap it with other
		max = num2
	}
	if max < num3 {
		max = num3
	}
	log.Println("max", max)
}

func SumOfDigits(num int) {
	sum := 0
	for num > 0 {
		digit := num % 10
		sum = sum + digit
		num = num / 10
	}
	log.Println(sum)
}

func FindNumPowerOfOtherNum(num1, num2 int) {
	temp := num1
	counter := 1
	for temp < num2 {
		temp = temp * num1
		counter++
		if num2 == temp {
			log.Println(counter, "power of each other", num1, num2)
			return
		}
	}
	log.Println("Not power of each other", num1, num2)
}

func PanicRecovery() {
	defer handleOutOfBounds()
	a := []int{5, 1, 2, 3}
	for i := 0; i < 10; i++ {
		log.Print(a[i])
	}

	log.Println("after recovery")
}

func handleOutOfBounds() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}

// first closing bracket must match with last opening bracket
func CheckParentheses(str string) {
	myStack := "" // create a stack to add opening and remove closing brackets

	for i := 0; i < len(str); i++ {
		if string(str[i]) == "{" || string(str[i]) == "(" || string(str[i]) == "[" { // check if opening
			myStack = myStack + string(str[i]) // add opening bracket to stack
		} else if myStack == "" && (string(str[i]) == "}" || string(str[i]) == ")" || string(str[i]) == "]") {
			log.Println("no opening bracket for the closing bracket") // if closing bracket comes inbetween, return
			return
		}

		if string(str[i]) == "}" || string(str[i]) == ")" || string(str[i]) == "]" { // check for closing
			index := len(myStack) - 1
			lastOpen := myStack[index:] // pop topmost stack element
			if (string(str[i]) == "}" && lastOpen == "{") || (string(str[i]) == "]" && lastOpen == "[") || (string(str[i]) == ")" && lastOpen == "(") {
				myStack = myStack[:index] // reduce stack after above open close testing
				continue
			} else {
				log.Println("first closing not matching with last opening")
				return
			}
		}
	}
	log.Println("all bracket matched")
}

func CheckPotentialPalindrome(str1 string) bool {
	m1 := make(map[string]int)

	for i := 0; i < len(str1); i++ { // check number of elements and their count
		if _, ok := m1[string(str1[i])]; !ok {
			m1[string(str1[i])] = 1
		} else {
			m1[string(str1[i])] = m1[string(str1[i])] + 1
		}
	}

	log.Println(m1)
	oddRepeat := 0           // if repeated chars occur odd times, it should be > 1. Even times no problem.
	for _, val := range m1 { // example: "ababa", here "a" came 3 times but still palindrome
		if val%2 != 0 {
			oddRepeat++
		}
	}
	if oddRepeat > 1 { // if "f" and "g" came only once (means repeated <1) "gababaf". No palindrome.
		log.Println("more than 1 single chars")
		return false
	}
	return true
}

func RemoveAdjacentDuplicateLetter(str string) {
	mystack := "" // abbaca

	for i := 0; i < len(str); i++ {
		index := len(mystack) - 1
		if mystack == "" {
			mystack = mystack + str[i:i+1] // push first letter into stack
			continue
		}

		// check if last letter matches with i+1th letter
		if mystack[index:] != str[i:i+1] { // means a matches with b, or b matches with b (abbaca example)
			mystack = mystack + str[i:i+1] // if not then push it to stack
			continue
		}

		// if last stack letter matches with i+1th letter
		if mystack[index:] == str[i:i+1] {
			mystack = mystack[:index] // remove that last pushed element
			continue
		}
	}

	log.Println(mystack)
	// 1 abbaca
	// 2 aaca
	// 3 ca
}

func FindSmallestNumber(arr []int) {
	min := arr[0]                   // assume first number is smallest
	for i := 1; i < len(arr); i++ { // if not then make the next min smallest
		if min > arr[i] {
			min = arr[i]
		}
	}
	log.Println(min)
}

func FindSubstringOccurance(str, substr string) {
	// abcdefabcabdaaa abc
	totalCounter := 0 // total number of occurances
	for i := 0; i < len(str); i++ {
		tempI := i         // preserve old i value
		singleCounter := 0 // use this to check if count matches with len(substring). means all the letters match.
		for j := 0; j < len(substr); j++ {
			if substr[j] == str[i] {
				i++ // increment i to check next letter position
				if i == len(str) {
					i = tempI // don't let it go beyond array length to avoid panic
				}
				singleCounter++
				continue
			}
			break
		}
		i = tempI
		if singleCounter == len(substr) {
			totalCounter++ // means j==i loop ran every time till all letters matched
		}
	}
	log.Println(totalCounter)
}

func PrintAllSubstrings(str string) {
	n := len(str)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			printstring := ""
			for k := i; k <= j; k++ {
				printstring = printstring + string(str[k])
			}
			log.Println(printstring)
		}
	}
}

func FindExpectedSumFromArray(arr []int, sumExpected int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if sum := arr[i] + arr[j]; sum == sumExpected {
				log.Println(arr[i], arr[j])
				break
			}
		}
	}

	// [1,2,5,5,8,10,15] find elements with sum=10
}

func main() {
}
