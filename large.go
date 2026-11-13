package largenumspell

import "math"

func basicLargeNumberName(i int) string {
	switch i {
	case 1:
		return "thousand"
	case 2:
		return "million"
	case 3:
		return "billion"
	case 4:
		return "trillion"
	case 5:
		return "quadrillion"
	case 6:
		return "quintillion"
	case 7:
		return "sextillion"
	case 8:
		return "septillion"
	case 9:
		return "octillion"
	case 10:
		return "nonillion"
	default:
		return ""
	}
}

func getMainPart(i int) string {
	switch i {
	case 1:
		return "decillion"
	case 2:
		return "vigintillion"
	case 3:
		return "trigintillion"
	case 4:
		return "quadragintillion"
	case 5:
		return "quinquagintillion"
	case 6:
		return "sexagintillion"
	case 7:
		return "septuagintillion"
	case 8:
		return "octogintillion"
	case 9:
		return "nonagintillion"
	case 10:
		return "centillion"
	default:
		return ""
	}
}

func getExtraPart(i int) string {
	switch i {
	case 1:
		return "un"
	case 2:
		return "duo"
	case 3:
		return "tre"
	case 4:
		return "quattuor"
	case 5:
		return "quin"
	case 6:
		return "sex"
	case 7:
		return "septen"
	case 8:
		return "octo"
	case 9:
		return "novem"
	default:
		return ""
	}
}

func advancedLargeNumberName(i int) string {
	temp := i - 1
	main := int(math.Floor(float64(temp) / 10))
	extra := temp % 10
	return getExtraPart(extra) + getMainPart(main)
}

func getLargeNumberName(i int) string {
	if i < 11 {
		return basicLargeNumberName(i)
	}
	return advancedLargeNumberName(i)
}
