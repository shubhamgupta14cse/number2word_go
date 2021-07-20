package main

import (
        "fmt"
)

var lowNames = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

var tensNames = []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

var bigNames = []string{"thousand", "million", "billion"}

func convert999(num int) string {

        s1 := lowNames[num/100] + " hundred"
        s2 := convert99(num % 100);
        if num <= 99 {
                return s2
        }
        s2  = "and " + s2
        if num == 0 {
                return s1
        } else if num == 100 {
                return s1
        } else {
                return s1 + " " + s2
        }
}

func convert99(num int) string {
        if num < 20 {
                return lowNames[num]
        }
        s := tensNames[num/10-2]
        if num == 0 {
                return s
        }
        return s + " " + lowNames[num % 10]
}

func convertNum2Words(num int) string {
        if num < 0 {
                return "no Less tha 0 are not allowed"
        }

        if num <= 999 {
                return convert999(num)
        }

        s := ""
        t := 0
        mod := num % 1000
        for num > 0 {
                if num != 0 {
                        s2 := convert999(num%1000)
                        if t > 0 {
                                s2 = s2 + " " + bigNames[t-1]
                        }
                        if s == "" {
                                s = s2
                        } else {
                                seperator := " "
                                if mod < 99 && t == 1 {
                                        seperator = " and "
                                }
                                s = s2 + seperator + s
                        }
                }
                num /= 1000
                t++
        }
        return s
}

func main() {
        fmt.Println("1 in words : ", convertNum2Words(1))
        fmt.Println("12 in words : ", convertNum2Words(12))
        fmt.Println("100 in words : ", convertNum2Words(100))
        fmt.Println("123 in words : ", convertNum2Words(123))
        fmt.Println("1012 in words : ", convertNum2Words(1012))
        fmt.Println("99718 in words : ", convertNum2Words(99718))
}
