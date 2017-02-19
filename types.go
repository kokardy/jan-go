package main

import (
	"strconv"
)

const (
	GS1TYPE_ITEM = "01"
)

type Code interface {
	CheckDigit() string
	CheckDigitOk() bool
}

func CalcCheckDigit(code string) string {
	count := 0
	for i := 0; i < len(code)-1; i++ {
		d, _ := strconv.Atoi(string(code[i]))
		if i%2 == 1 {
			d *= 3
		}
		count += d
	}
	count %= 10
	digit := 10 - count
	digit %= 10
	return strconv.Itoa(digit)
}

type JAN string

func (j JAN) CheckDigit() string {
	return CalcCheckDigit(string(j))
}

func (j JAN) CheckDigitOK() bool {
	digit := j.CheckDigit()
	if string(j[12]) == digit {
		return true
	}
	return false
}

type GS1 string

func (g GS1) CheckDigit() string {
	return CalcCheckDigit(string(g))
}
func (g GS1) CheckDigitOK() bool {
	digit := g.CheckDigit()
	if string(g[13]) == digit {
		return true
	}
	return false
}

func (g GS1) ToJAN() JAN {
	pre_jan := JAN(string(g[3:16]))
	check_digit := pre_jan.CheckDigit()
	return JAN(string(g[3:15]) + check_digit)
}
