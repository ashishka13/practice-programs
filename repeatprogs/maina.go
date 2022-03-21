package repeatprogs

import (
	"fmt"
	"log"
)

func StringRevWithoutSplChars2(str string) {
	//
}

func AnagramString2(str1, str2 string) {
	if len(str1) != len(str2) {
		log.Println("not anagram")
		return
	}

	m1 := make(map[string]int)
	m2 := make(map[string]int)

	for i := 0; i < len(str1); i++ {
		if _, ok := m1[string(str1[i])]; !ok {
			m1[string(str1[i])] = 1
		} else {
			m1[string(str1[i])] = m1[string(str1[i])] + 1
		}

		if _, ok := m2[string(str2[i])]; !ok {
			m2[string(str2[i])] = 1
		} else {
			m2[string(str2[i])] = m2[string(str2[i])] + 1
		}

	}
	if fmt.Sprint(m1) != fmt.Sprint(m2) {
		log.Println("not anagrams")
		return
	}
	log.Println("anagrams")
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
