package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("assemblerCode.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sliceOfVar := []string{}
	sliceOfValues := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if strings.Contains(scanner.Text(), "DB") {
			temp := strings.Split(scanner.Text(), " ")
			sliceOfVar = append(sliceOfVar, temp[0])
			sliceOfValues = append(sliceOfValues, temp[2])
			fmt.Println(sliceOfVar, sliceOfValues)
		} else if strings.Contains(scanner.Text(), "INC") || strings.Contains(scanner.Text(), "DEC") {
			temp := strings.Split(scanner.Text(), " ")
			for i := range sliceOfVar {
				if sliceOfVar[i] == temp[1] {
					tempVar, err := strconv.Atoi(sliceOfValues[i])
					if err != nil {
						log.Fatal(err)
					}
					switch temp[0] {
					case "INC":
						tempVar++
					case "DEC":
						tempVar--
					}
					sliceOfValues[i] = strconv.Itoa(tempVar)
					break
				}
			}
		} else if strings.Contains(scanner.Text(), "ADD") || strings.Contains(scanner.Text(), "SUB") || strings.Contains(scanner.Text(), "MUL") ||
			strings.Contains(scanner.Text(), "DIV") {
			var firstVar, secondVar, index int
			temp := strings.Split(scanner.Text(), " ")
			for i := range sliceOfVar {
				if sliceOfVar[i] == temp[1] {
					firstVar, err = strconv.Atoi(sliceOfValues[i])
					if err != nil {
						log.Fatal(err)
					}
					index = i
				}
				if sliceOfVar[i] == temp[2] {
					secondVar, err = strconv.Atoi(sliceOfValues[i])
					if err != nil {
						log.Fatal(err)
					}
				}
			}
			switch temp[0] {
			case "ADD":
				sliceOfValues[index] = strconv.Itoa(firstVar + secondVar)
			case "SUB":
				sliceOfValues[index] = strconv.Itoa(firstVar - secondVar)
			case "MUL":
				sliceOfValues[index] = strconv.Itoa(firstVar * secondVar)
			case "DIV":
				sliceOfValues[index] = strconv.Itoa(firstVar / secondVar)
			}
			fmt.Println("After ops:", sliceOfVar, sliceOfValues)
		}
	}

	for i := range sliceOfVar {
		fmt.Println(sliceOfVar[i], sliceOfValues[i])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
