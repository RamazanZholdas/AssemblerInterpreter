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
		if strings.Contains(scanner.Text(), "ret") {
			fmt.Println("Resutl of function:")
			for i := range sliceOfVar {
				fmt.Println(sliceOfVar[i], sliceOfValues[i])
			}
			sliceOfVar = []string{}
			sliceOfValues = []string{}
		} else if strings.Contains(scanner.Text(), "DB") {
			temp := strings.Split(scanner.Text(), " ")
			sliceOfVar = append(sliceOfVar, temp[0])
			sliceOfValues = append(sliceOfValues, temp[2])
		} else if strings.Contains(scanner.Text(), "INC") || strings.Contains(scanner.Text(), "DEC") {
			temp := strings.Split(scanner.Text(), " ")
			for i := range sliceOfVar {
				if sliceOfVar[i] == temp[1] {
					var tempVar int
					var itsHex, itsBin bool
					if strings.Contains(sliceOfValues[i], "h") {
						tempVar = hexToDec(sliceOfValues[i][:len(sliceOfValues[i])-1])
						itsHex = true
					} else if strings.Contains(sliceOfValues[i], "b") {
						tempVar = binToDec(sliceOfValues[i][:len(sliceOfValues[i])-1])
						itsBin = true
					} else {
						tempVar, err = strconv.Atoi(sliceOfValues[i])
						if err != nil {
							log.Fatal(err)
						}
					}
					switch temp[0] {
					case "INC":
						tempVar++
					case "DEC":
						tempVar--
					}
					if itsHex {
						sliceOfValues[i] = decToHex(strconv.Itoa(tempVar)) + "h"
					} else if itsBin {
						sliceOfValues[i] = decToBin(strconv.Itoa(tempVar)) + "b"
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
			var result, result2 string
			var itsHex, itsBin, xchg bool
			temp := strings.Split(scanner.Text(), " ")
			for i := range sliceOfVar {
				if sliceOfVar[i] == temp[1] {
					if strings.Contains(sliceOfValues[i], "h") {
						firstVar = hexToDec(sliceOfValues[i][:len(sliceOfValues[i])-1])
						itsHex = true
					} else if strings.Contains(sliceOfValues[i], "b") {
						firstVar = binToDec(sliceOfValues[i][:len(sliceOfValues[i])-1])
						itsBin = true
					} else {
						firstVar, err = strconv.Atoi(sliceOfValues[i])
						if err != nil {
							log.Fatal(err)
						}
					}
					if err != nil {
						log.Fatal(err)
					}
					index = i
				}
				if sliceOfVar[i] == temp[2] {
					if strings.Contains(sliceOfValues[i], "h") {
						secondVar = hexToDec(sliceOfValues[i][:len(sliceOfValues[i])-1])
					} else if strings.Contains(sliceOfValues[i], "b") {
						secondVar = binToDec(sliceOfValues[i][:len(sliceOfValues[i])-1])
					} else {
						secondVar, err = strconv.Atoi(sliceOfValues[i])
						if err != nil {
							log.Fatal(err)
						}
					}
					if err != nil {
						log.Fatal(err)
					}
				}
			}
			switch temp[0] {
			case "ADD":
				result = strconv.Itoa(firstVar + secondVar)
			case "SUB":
				result = strconv.Itoa(firstVar - secondVar)
				fmt.Println(result)
			case "MUL":
				result = strconv.Itoa(firstVar * secondVar)
			case "DIV":
				result = strconv.Itoa(firstVar / secondVar)
			case "MOV":
				result = strconv.Itoa(secondVar)
			case "XCHG":
				xchg = true
				result = strconv.Itoa(secondVar)
				result2 = strconv.Itoa(firstVar)
			}

			if itsHex {
				sliceOfValues[index] = decToHex(result) + "h"
			} else if itsBin {
				sliceOfValues[index] = decToBin(result) + "b"
			} else if xchg {
				if strings.Contains(result, "h") {
					sliceOfValues[index] = decToHex(result) + "h"
					sliceOfValues[index+1] = decToHex(result2) + "h"
				} else if strings.Contains(result, "b") {
					sliceOfValues[index] = decToBin(result) + "b"
					sliceOfValues[index+1] = decToBin(result2) + "b"
				} else {
					sliceOfValues[index] = result
					sliceOfValues[index+1] = result2
				}
			} else {
				sliceOfValues[index] = result
			}
		}
	}

	fmt.Println()
	fmt.Println("Result:")
	for i := range sliceOfVar {
		fmt.Println(sliceOfVar[i], sliceOfValues[i])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func hexToDec(hex string) int {
	dec, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(dec)
}

func binToDec(bin string) int {
	dec, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(dec)
}

func decToHex(dec string) string {
	hex, err := strconv.ParseInt(dec, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", hex)
}

func decToBin(dec string) string {
	bin, err := strconv.ParseInt(dec, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%b", bin)
}
