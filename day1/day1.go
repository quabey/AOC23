package day1

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
    file, _ := os.Open("day1input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum += calcWord(line)
		nums = nil
	}
	fmt.Println(sum)
}

func calcWord(word string) (calc int) {
	for index, char := range word {
		if unicode.IsNumber(char) {
			nums = append(nums, char)
		} 
		switch char {
		case 'o': 
			if (len(word) >= index + 3) && word[index:index + 3] == "one"{
				nums = append(nums, '1')
			}
		case 't':
			if (len(word) >= index + 3) &&  word[index:index + 3] == "two" { 
				nums = append(nums, '2')
			} else if (len(word) >= index + 5) && word[index:index + 5] == "three" { 
				nums = append(nums, '3')
			}
		case 'f':
			if (len(word) >= index + 4) &&  word[index:index + 4] == "four" {
				nums = append(nums, '4')
			} else if (len(word) >= index + 4) && word[index:index + 4] == "five" {
				nums = append(nums, '5')
			}
		case 's':
			if (len(word) >= index + 3) && word[index:index + 3] == "six" {
				nums = append(nums, '6')
			} else if (len(word) >= index + 5) && word[index:index + 5] == "seven" {
				nums = append(nums, '7')
			}
		case 'e':
			if (len(word) >= index + 5) && word[index:index + 5] == "eight" {
				nums = append(nums, '8')
			}
		case 'n':
			if (len(word) >= index + 4) && word[index:index + 4] == "nine" {
				nums = append(nums, '9')
			}
		}
	}
	calc, _ = strconv.Atoi(string(nums[0]) + string(nums[len(nums) -1]))
	fmt.Println(word, nums, calc)
	nums = nil
	return 
}