package lexer

import (
	"pika/pkg/lexer/token_type"
	"strconv"

	"golang.org/x/exp/slices"
)

func IsSkippable(char rune) bool {
	return slices.Contains(token_type.SkippableChars, char)
}

func NextChar(src *[]rune) string {
	char := (*src)[0]
	*src = (*src)[1:]
	return string(char)
}

/*  FirstReturn: Number extracted
 * 	SecondReturn: Rest of the string
 */
func ExtractInt(src []rune) (string, []rune) {
	var num = ""

	for len(src) > 0 && (IsInt(src[0]) || src[0] == '.') {
		num += NextChar(&src)
	}

	return num, src
}

/*  FirstReturn: String extracted
 * 	SecondReturn: Rest of the string
 */
func ExtractIdentifier(src []rune) (string, []rune) {
	var str = ""

	for len(src) > 0 && IsIdentifier(src[0]) {
		str += NextChar(&src)
	}

	return str, src
}

func IsIdentifier(char rune) bool {
	return slices.Contains(token_type.AllowedIdentifierChars, char)
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
