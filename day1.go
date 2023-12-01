package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var nums []rune
var sum int = 0

func main() {
    scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		sum += calcWord(line)
		fmt.Println(calcWord(line), sum)
		nums = nil
	}
}

func calcWord(word string) (calc int) {
	for _, char := range word {
		if unicode.IsNumber(char) {
			nums = append(nums, char)
		} 
	}
	calc, _ = strconv.Atoi(string(nums[0]) + string(nums[len(nums) -1]))
	nums = nil
	return 
}