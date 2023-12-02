package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var sum int 

func main() {
	attempt1()
	attempt2()
}

func attempt2() {
    file, _ := os.Open("day2input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()
		line = strings.Split(line, ":")[1]
		line = strings.ReplaceAll(line," ", "")
		line = strings.ReplaceAll(line, "green", "g")
		line = strings.ReplaceAll(line, "red", "r")
		line = strings.ReplaceAll(line, "blue", "b")
		sum += checkLineArray(line)
	}
	fmt.Println(sum)
}

func checkLineArray(line string) int {
	colors := map[rune]int{'r': 1, 'g': 1, 'b': 1}
	re := regexp.MustCompile("[;,]")
	array := re.Split(line, -1)
	for _, item := range array {
		var color rune = rune(item[len(item) - 1]) 
		offset, _ := strconv.Atoi(string(item[0:len(item) - 1]))
		if offset > colors[color] {colors[color] = offset} 
	}
	return colors['r'] * colors['g'] * colors['b']
}


func attempt1() {
    file, _ := os.Open("day2input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()
		line = strings.Split(line, ":")[1]
		line = strings.ReplaceAll(line," ", "")
		line = strings.ReplaceAll(line, "green", "g")
		line = strings.ReplaceAll(line, "red", "r")
		line = strings.ReplaceAll(line, "blue", "b")
		sum += checkLine(line)
		
	}
	fmt.Println(sum)
	sum = 0
}

func checkLine(line string) int {
	colors := map[rune]int{'r': 1, 'g': 1, 'b': 1}

	for i, char := range line {
		if char == 'r' || char == 'g' || char == 'b' {
			offset, _ := strconv.Atoi(string(line[i-1]))
			if i > 1 && line[i-2] == '1' && offset+10 > colors[char] {
				colors[char] = offset + 10
			} else if offset > colors[char] {
				colors[char] = offset
			}
			if i > 1 && line[i-2] == '2' {
				colors[char] = 20
			}
		}
	}
	return colors['r'] * colors['g'] * colors['b']
}