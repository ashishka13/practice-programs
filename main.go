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

func StringRev(strInput string) {
	strProcess := make([]string, 0)
	strProcess2 := make([]string, len(strInput))

	for i, j := len(strInput)-1, 0; i >= 0; i, j = i-1, j+1 {
		if (strInput[i] >= 'A' && strInput[i] <= 'Z') || (strInput[i] >= 'a' && strInput[i] <= 'z') {
			stringVal := string(strInput[i])
			strProcess = append(strProcess, stringVal)
		}

		if (strInput[j] >= 'A' && strInput[j] <= 'Z') || (strInput[j] >= 'a' && strInput[j] <= 'z') {
		} else {
			stringVal := string(strInput[j])
			strProcess2[j] = stringVal
		}
	}

	position2 := 0
	op2 := ""
	for i := 0; i < (len(strProcess2)); i++ {
		if strProcess2[i] == "" {
			strProcess2[i] = strProcess[position2]
			position2++
		}
		op2 = op2 + strProcess2[i]
	}

	log.Println(strInput)
	log.Println(op2)
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

func main() {
	ArmstrongNum(163)
}
