package largenumspell

import (
	"errors"
	"math"
	"regexp"
	"unicode"
)

var (
	numberFormat = regexp.MustCompile(`([+-]?\d+)(\.\d+)?([eE][+-]?\d+)?`)
	ErrorFormat  = errors.New("wrong numeric string format")
)

func getPartsLen(str []rune) int {
	return int(math.Ceil(float64(len(str)) / 3.0))
}

func getLeftSide(str []rune) []string {
	var level, temp, quantity int
	ret := make([]string, 0, 100)
	partsLen := getPartsLen(str)
	total := 0
	mod := len(str) % 3
	if mod == 0 {
		temp = 3
	} else {
		temp = mod
	}
	for i := 1; i < partsLen; i++ {
		level = partsLen - i
		quantity = len(ret)
		numbersSpellout(str[total:temp], &ret)
		if len(ret) > quantity {
			ret = append(ret, getLargeNumberName(level))
		}
		total = temp
		temp += 3
	}
	if len(ret) > 0 {
		numbersSpellout(str[total:temp], &ret)
	} else {
		zero := true
		for _, i := range str[total:temp] {
			if i != '0' {
				zero = false
				break
			}
		}
		if zero {
			ret = append(ret, "zero")
		} else {
			numbersSpellout(str[total:temp], &ret)
		}
	}
	return ret
}

func getRightSide(str []rune) []string {
	var quantity int
	ret := make([]string, 0, 100)
	partsLen := getPartsLen(str) + 1
	total := 0
	lastPart := false
	for i := 1; i < partsLen; i++ {
		temp := total + 3
		if temp > len(str) {
			temp = len(str)
			lastPart = true
		}
		quantity = len(ret)
		numbersSpellout(str[total:temp], &ret)
		if len(ret) > quantity {
			if lastPart {
				part := getLargeNumberName(i - 1)
				otherPart := "ten"
				if temp-total > 1 {
					otherPart = "hundred"
				}
				if len(part) > 0 {
					ret = append(ret, otherPart+"-"+part+"th")
				} else {
					ret = append(ret, otherPart+"th")
				}
			} else {
				ret = append(ret, getLargeNumberName(i)+"th")
			}
		}
		total += 3
	}
	return ret
}

func parseSide(match []rune, offset int, parser func([]rune) []string) []string {
	var result []string
	partOffset := offset
	if !unicode.IsDigit(match[partOffset]) {
		if match[partOffset] == '+' {
			result = append(result, "plus")
		} else if match[partOffset] == '-' {
			result = append(result, "minus")
		}
		partOffset++
	}
	temp := parser(match[partOffset:])
	if len(temp) > 0 && len(result) > 0 {
		result = append(result, temp...)
	} else {
		result = temp
	}
	return result
}

func Spellout(str string) ([]string, error) {
	result := make([]string, 0, 100)
	var temp []string
	match := numberFormat.FindStringSubmatch(str)
	if match == nil {
		return nil, ErrorFormat
	}
	temp = parseSide([]rune(match[1]), 0, getLeftSide)
	if len(temp) > 0 {
		result = append(result, temp...)
	}
	if len(match[2]) > 0 {
		temp = parseSide([]rune(match[2]), 1, getRightSide)
		if len(temp) > 0 {
			result = append(result, "and")
			result = append(result, temp...)
		}
	}
	if len(match[3]) > 0 {
		temp = parseSide([]rune(match[3]), 1, getLeftSide)
		if len(temp) > 0 {
			result = append(result, "exponent")
			result = append(result, temp...)
		}
	}
	return result, nil
}
