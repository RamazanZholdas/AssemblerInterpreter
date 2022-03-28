package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//починить проверку хекс числа и парсить тока нужную часть хекс числа так как при декларации хекс числа в конце ставят h
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
			//fmt.Println(sliceOfVar, sliceOfValues)
		} else if strings.Contains(scanner.Text(), "INC") || strings.Contains(scanner.Text(), "DEC") {
			temp := strings.Split(scanner.Text(), " ")
			for i := range sliceOfVar {
				if sliceOfVar[i] == temp[1] {
					isHex := false
					parseThis := sliceOfValues[i]
					if parseThis, err = IsHex(sliceOfValues[i]); err == nil {
						isHex = true
					}
					tempVar, err := strconv.Atoi(parseThis)
					if err != nil {
						log.Fatal(err)
					}
					switch temp[0] {
					case "INC":
						tempVar++
					case "DEC":
						tempVar--
					}
					if isHex {
						sliceOfValues[i] = DecToHex(tempVar)
					} else {
						sliceOfValues[i] = strconv.Itoa(tempVar)
					}
					break
				}
			}
		} else if strings.Contains(scanner.Text(), "ADD") || strings.Contains(scanner.Text(), "SUB") ||
			strings.Contains(scanner.Text(), "MUL") || strings.Contains(scanner.Text(), "DIV") ||
			strings.Contains(scanner.Text(), "MOV") || strings.Contains(scanner.Text(), "XCHG") {
			var firstVar, secondVar, index int
			var isHex1, isHex2 bool
			temp := strings.Split(scanner.Text(), " ")
			for i := range sliceOfVar {
				if sliceOfVar[i] == temp[1] {
					isHex1 = false
					parseThis := sliceOfValues[i]
					if parseThis, err = IsHex(sliceOfValues[i]); err == nil {
						isHex1 = true
					}
					firstVar, err = strconv.Atoi(parseThis)
					if err != nil {
						log.Fatal(err)
					}
					index = i
				} else if sliceOfVar[i] == temp[2] {
					isHex2 = false
					parseThis := sliceOfValues[i]
					if parseThis, err = IsHex(sliceOfValues[i]); err == nil {
						isHex2 = true
					}
					secondVar, err = strconv.Atoi(parseThis)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
			switch temp[0] {
			case "ADD":
				if isHex1 && isHex2 {
					sliceOfValues[index] = DecToHex(firstVar + secondVar)
				} else {
					sliceOfValues[index] = strconv.Itoa(firstVar + secondVar)
				}
			case "SUB":
				if isHex1 && isHex2 {
					sliceOfValues[index] = DecToHex(firstVar - secondVar)
				} else {
					sliceOfValues[index] = strconv.Itoa(firstVar - secondVar)
				}
			case "MUL":
				if isHex1 && isHex2 {
					sliceOfValues[index] = DecToHex(firstVar * secondVar)
				} else {
					sliceOfValues[index] = strconv.Itoa(firstVar * secondVar)
				}
			case "DIV":
				if isHex1 && isHex2 {
					sliceOfValues[index] = DecToHex(firstVar / secondVar)
				} else {
					sliceOfValues[index] = strconv.Itoa(firstVar / secondVar)
				}
			case "MOV":
				secondVar = firstVar
				if isHex1 && isHex2 {
					sliceOfValues[index] = DecToHex(secondVar)
				} else {
					sliceOfValues[index] = strconv.Itoa(secondVar)
				}
			case "XCHG":
				if isHex1 && isHex2 {
					sliceOfValues[index] = DecToHex(firstVar)
					sliceOfValues[index+1] = DecToHex(secondVar)
				} else {
					sliceOfValues[index] = strconv.Itoa(secondVar)
					sliceOfValues[index+1] = strconv.Itoa(firstVar)
				}
			}
			//fmt.Println("After ops:", sliceOfVar, sliceOfValues)
		}
	}

	fmt.Println(sliceOfValues)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func IsHex(num string) (string, error) {
	dec, err := strconv.Atoi(num)
	if err != nil {
		return "", err
	}
	if strings.Contains(num, "h") {
		hex := strconv.FormatInt(int64(dec), 16)
		return hex, nil
	}
	return "", errors.New("invalid hex")
}

func HexToDec(hex string) (string, error) {
	num, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(num)), nil
}

/*
func BinToDec(bin string) (string, error) {
	num, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(num)), nil
}*/

func DecToHex(dec int) string {
	hex := strconv.FormatInt(int64(dec), 16)
	return hex
}

/*
func DecToBin(dec int) string {
	bin := strconv.FormatInt(int64(dec), 2)
	return bin
}*/
