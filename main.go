package main

import (
	"log"
	"strings"
)

func PrintStringRecursion(s string) {
	log.Println(s)
	if s == "" {
		return
	}
	s = s[1:] // skip first position and reassign to same string variable
	PrintStringRecursion(s)
}

func StringRevWithoutSplChars(strInput string) {
	processInput1 := strInput
	revChars := []string{}
	placeHolders := make([]string, len(strInput))

	for len(processInput1) > 0 {
		index := len(processInput1) - 1
		s := processInput1[index:]
		if (s >= "a" && s <= "z") || (s >= "A" && s <= "Z") {
			revChars = append(revChars, s)
			placeHolders[index] = ""
			processInput1 = processInput1[:index]
			continue
		}
		placeHolders[index] = s
		processInput1 = processInput1[:index]
	}

	for key, val := range placeHolders {
		if val == "" {
			placeHolders[key] = revChars[0]
			revChars = revChars[1:]
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

	wordsArray := strings.Fields(sentence)
	wordsString := ""
	for _, val := range wordsArray {
		wordsArray2 := strings.Split(val, "")
		singleRevWord := ""
		reverseArr := []string{}

		for len(wordsArray2) > 0 {
			lastindex := len(wordsArray2) - 1
			b := wordsArray2[lastindex:]
			wordsArray2 = wordsArray2[:lastindex]
			reverseArr = append(reverseArr, b[0])
		}

		for i := 0; i < len(reverseArr); i++ {
			singleRevWord = singleRevWord + reverseArr[i]
		}

		wordsString = wordsString + " " + singleRevWord
	}
	log.Println("op", wordsString)
}

func PalindromeString(str string) {
	for i, j := len(str)-1, 0; i > 0; i, j = i-1, j+1 {
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
	}
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

//0 1 1 2 3 5 8 13 21 34 55 89 114
func FibonacciRecursive(fiboRange, prev, next, answer int) {
	if fiboRange == 0 {
		return
	}
	answer = prev + next
	prev = next
	next = answer
	log.Print("answer", answer)
	FibonacciRecursive(fiboRange-1, prev, next, answer)
}

func FindDupStringChars(str string) {
	strArr := strings.Split(str, "")
	dupMap := make(map[string]int)
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
	for len(str) > 0 {
		lastindex := len(str) - 1
		b := str[lastindex:]
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
			if numbers[i] > numbers[j] {
				temp := numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = temp
			}
			counter++
		}
	}
	log.Println(numbers)
}

func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right, pivot := 0, len(a)-1, len(a)/2

	a[pivot], a[right] = a[right], a[pivot] // swap

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left] // swap

	QuickSort(a[:left])
	QuickSort(a[left+1:])

	return a
}

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
				counterflag++
				break
			}
		}
	}
	if counterflag == len(str1) {
		log.Println("anagrams", counterflag)
		return
	}
	log.Println("not anagrams", counterflag)
}

func GetSubstringCombinations(str string) {
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
			if num[i] != num[j] {
				temp := num[i]
				num[i] = num[j]
				num[j] = temp
				log.Println(num)
			}
		}
	}
}

func MaxOfThreeNum(num1, num2, num3 int) {
	max := num1
	if max < num2 {
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
	for num1 < num2 {
		num1 := num1 * num1
		log.Println("print", num1, num2)

		if num2 == num1 {
			log.Println("power of each other", num1, num2)
			return
		}
	}
	log.Println("Not power of each other", num1, num2)
}

func main() {
	FindNumPowerOfOtherNum(2, 126)
}
