package main

import (
	"log"
)

func main() {
	retVal := PrintStringClosure("ajsjnoiasogjnasdgjan")

	for retVal() {
	}
}

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
