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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	passports := make([]map[string]string, 0)
	passport := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, passport)
			passport = make(map[string]string)
			continue
		}
		for _, i := range strings.Split(line, " ") {
			key := strings.Split(i, ":")[0]
			passport[key] = strings.Split(i, ":")[1]
		}
	}
	passports = append(passports, passport)
	fmt.Printf("Problem 1: %v\n", problem1(passports))
	fmt.Printf("Problem 2: %v\n", problem2(passports))
}

func problem1(passports []map[string]string) int {
	fmt.Println(len(passports))
	count := 0
	for _, passport := range passports {
		cidPressent := false
		keys := 0
		for key := range passport {
			if key == "cid" {
				cidPressent = true
			}
			keys++
		}
		if keys == 8 || (keys == 7 && !cidPressent) {
			count++
		}
	}
	return count
}

func problem2(passports []map[string]string) int {
	fmt.Println(len(passports))
	count := 0
	for _, passport := range passports {
		cidPressent := false
		inValid := false
		keys := 0
		for key := range passport {
			if key == "cid" {
				cidPressent = true
			}
			switch key {
			case "byr":
				byr, _ := strconv.Atoi(passport[key])
				if !(1920 <= byr && byr <= 2002) {
					inValid = true
					break
				}
			case "iyr":
				iyr, _ := strconv.Atoi(passport[key])
				if !(2010 <= iyr && iyr <= 2020) {
					inValid = true
					break
				}
			case "eyr":
				eyr, _ := strconv.Atoi(passport[key])
				if !(2020 <= eyr && eyr <= 2030) {
					inValid = true
					break
				}
			case "hgt":
				if passport[key][len(passport[key])-2:] == "cm" {
					hgt, _ := strconv.Atoi(passport[key][:len(passport[key])-2])
					if !(150 <= hgt && hgt <= 193) {
						inValid = true
						break
					}
				} else if passport[key][len(passport[key])-2:] == "in" {
					hgt, _ := strconv.Atoi(passport[key][:len(passport[key])-2])
					if !(59 <= hgt && hgt <= 76) {
						inValid = true
						break
					}
				} else {
					inValid = true
					break
				}
			case "hcl":
				if passport[key][0] != '#' || len(passport[key]) != 7 {
					inValid = true
					break
				} else {
					for _, i := range passport[key][1:] {
						if !(('0' <= i && i <= '9') || ('a' <= i && i <= 'f')) {
							inValid = true
							break
						}
					}
				}
			case "ecl":
				valid := false
				for _, p := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
					if passport[key] == p {
						valid = true
						break
					}
				}
				if !valid {
					inValid = true
					break
				}
			case "pid":
				if len(passport[key]) != 9 {
					inValid = true
					break
				}
				for _, i := range passport[key] {
					if !('0' <= i && i <= '9') {
						inValid = true
						break
					}
				}
			}
			keys++
		}
		if !inValid && (keys == 8 || (keys == 7 && !cidPressent)) {
			count++
		}
	}
	return count
}
