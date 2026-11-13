package largenumspell

func ones(digit rune) string {
	switch digit {
	case '1':
		return "one"
	case '2':
		return "two"
	case '3':
		return "three"
	case '4':
		return "four"
	case '5':
		return "five"
	case '6':
		return "six"
	case '7':
		return "seven"
	case '8':
		return "eight"
	case '9':
		return "nine"
	default:
		return ""
	}
}

func teens(digit rune) string {
	switch digit {
	case '0':
		return "ten"
	case '1':
		return "eleven"
	case '2':
		return "twelve"
	case '3':
		return "thirteen"
	case '4':
		return "fourteen"
	case '5':
		return "fifteen"
	case '6':
		return "sixteen"
	case '7':
		return "seventeen"
	case '8':
		return "eighteen"
	case '9':
		return "nineteen"
	default:
		return ""
	}
}

func tens(digit rune) string {
	switch digit {
	case '2':
		return "twenty"
	case '3':
		return "thirty"
	case '4':
		return "fourty"
	case '5':
		return "fifty"
	case '6':
		return "sixty"
	case '7':
		return "seventy"
	case '8':
		return "eighty"
	case '9':
		return "ninety"
	default:
		return ""
	}
}

func doNumbersSpellout(numbers []rune, buffer *[]string) {
	if len(numbers) == 1 {
		one_result := ones(numbers[0])
		if len(one_result) > 0 {
			*buffer = append(*buffer, one_result)
		}
	} else {
		if numbers[0] == '1' {
			part := teens(numbers[1])
			if len(part) > 0 {
				*buffer = append(*buffer, part)
			}
		} else {
			first_part := tens(numbers[0])
			second_part := ones(numbers[1])
			temp := 0
			if len(first_part) > 0 {
				temp += 1
			}
			if len(second_part) > 0 {
				temp += 2
			}
			switch temp {
			case 3:
				*buffer = append(*buffer, first_part+"-"+second_part)
			case 2:
				*buffer = append(*buffer, second_part)
			case 1:
				*buffer = append(*buffer, first_part)
			}
		}
	}
}

func numbersSpellout(numbers []rune, buffer *[]string) {
	var temp []rune
	if len(numbers) == 3 {
		part := ones(numbers[0])
		if len(part) > 0 {
			*buffer = append(*buffer, part, "hundred")
		}
		temp = numbers[1:]
	} else {
		temp = numbers
	}
	doNumbersSpellout(temp, buffer)
}
