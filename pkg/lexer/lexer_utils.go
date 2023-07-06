package lexer

import (
	"pika/pkg/lexer/token_type"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func IsSkippable(char rune) bool {
	return slices.Contains(token_type.SkippableChars, char)
}

func NextChar(src *[]rune) string {
	if len(*src) <= 0 {
		return ""
	}

	char := (*src)[0]
	if len(*src)-1 > 0 {
		*src = (*src)[1:]
	} else {
		*src = []rune{}
	}

	return string(char)
}

/*  FirstReturn: Number extracted
 * 	SecondReturn: Rest of the string
 */
func ExtractInt(src []rune) (string, []rune) {
	if len(src) <= 0 {
		return "", src
	}

	isNegative := src[0] == '-'
	num := ""

	if isNegative {
		num = "-"
		NextChar(&src)
	}

	for len(src) > 0 && (IsInt(src[0]) || src[0] == '.') {
		num += NextChar(&src)
	}

	return num, src
}

/*  FirstReturn: String extracted
 * 	SecondReturn: Rest of the string
 */
func ExtractIdentifier(src []rune) (string, []rune) {
	if len(src) <= 0 || !IsAlpha(src[0]) {
		return "", src
	}

	var str = ""

	for len(src) > 0 && (IsAlpha(src[0]) || slices.Contains(token_type.AllowedIdentifierChars, src[0])) {
		str += NextChar(&src)
	}

	return str, src
}

func IsAlpha(char rune) bool {
	return strings.ToUpper(string(char)) != strings.ToLower(string(char)) || slices.Contains(token_type.AllowedIdentifierCharsWithFirst, char)
}

/*  FirstReturn: Keyword type
 * 	SecondReturn: Is the keyword valid
 */
func IsKeyword(src string) (token_type.TokenType, bool) {
	keywordType, ok := token_type.KEYWORDS[src]

	return keywordType, ok
}

func IsInt(char rune) bool {
	val, err := strconv.ParseFloat(string(char), 64)
	return err == nil && val >= 0
}
