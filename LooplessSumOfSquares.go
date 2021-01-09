package main

import (
	"fmt"
)

/*
 CHALLENGE DETAILS
Description
    We want you to calculate the sum of squares of given integers, excluding any negatives.
    The first line of the input will be an integer N (1 <= N <= 100), indicating the number of test cases to follow.
    Each of the test cases will consist of a line with an integer X (0 < X <= 100), followed by another line consisting of X number of space-separated integers Yn (-100 <= Yn <= 100).
    For each test case, calculate the sum of squares of the integers, excluding any negatives, and print the calculated sum in the output.
    Note: There should be no output until all the input has been received.
    Note 2: Do not put blank lines between test cases solutions.
    Note 3: Take input from standard input, and output to standard output.

Rules
    Write your solution using Go Programming Language
    Your source code must be a single file (package main)
    Do not use any for statement
    You may only use standard library packages


Sample Input:
2
4
3 -1 1 14
5
9 6 -53 32 16

Sample Output:
206
1397
*/

func main() {
	numOfTestCases := 0

	fmt.Scan(&numOfTestCases)
	output := make([]int, numOfTestCases)
	processSquares(numOfTestCases, output)
	printSum(0, output)
}

//recursive function to process the input
func processSquares(i int, output []int){
	if i == 0{
		return
	}
	rowIterations := 0
	fmt.Scan(&rowIterations)

	output[len(output) - i] = addPositiveSquares(rowIterations)
	processSquares(i - 1, output)
}

//recursively add iter# of integers from the stdin, excluding negative integers
func addPositiveSquares(iter int) int{
	num := 0
	if iter != 0{ //protect from hanging in the case of no integers provided
		fmt.Scan(&num)
	}
	if iter <= 1 { //recursion finished
		return num * num
	} else if num < 0{ //exclude negatives from summation
		return addPositiveSquares(iter - 1)
	} else {
		return num*num + addPositiveSquares(iter - 1) //add and continue recursion
	}
}

//recursively print a slice of integers
func printSum(i int, output []int){
	if i == len(output){
		return
	}
	fmt.Println(output[i])
	printSum(i + 1, output)
}