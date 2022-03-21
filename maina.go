package main

import "log"

func StringRevWithoutSplChars2(str string) {

}

func AnagramString2(str1, str2 string) {
	if len(str1) != len(str2) {
		log.Println("not anagram")
		return
	}

	flag := true
	for i := 0; i < len(str1); i++ {
		if !flag {
			log.Println("not anagram")
			return
		}
		for j := 0; j < len(str2); j++ {
			if str1[i] != str2[j] {
				flag = false
			} else {
				flag = true
				continue
			}
		}

	}
	if flag {
		log.Println("anagram")
	}
}

/*
func PrintStringRecursion(s string) {
func StringRevWithoutSplChars(strInput string) {
func PrintStringClosure(a string) func() bool {
func ReverseWordsInString(sentence string) {
func PalindromeString(str string) {
func PalindromeNumber(inputNum int) {
func PrimeOrNot(num int) {
func ArmstrongNum(num int) {
func ReverseNumber(num int) {
func PrimeOrNotInRange(num1, num2 int) {
func Fibonacci(foboRange int) {
func FibonacciRecursive(fiboRange, prev, next, answer int) {
func FindDupStringChars(str string) {
func FindUniqueChars(str string) {
func ReverseStringNormal(str string) {
func BubbleSort(numbers []int) {
func QuickSort(a []int) []int {
func AnagramString(str1, str2 string) {
func GetSubstringCombinations(str string) {
func ShuffleNumbers(num []int) {
func MaxOfThreeNum(num1, num2, num3 int) {
func SumOfDigits(num int) {
func FindNumPowerOfOtherNum(num1, num2 int) {
*/
