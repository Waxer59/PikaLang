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
	char := (*src)[0]
	*src = (*src)[1:]
	return string(char)
}

/*  FirstReturn: Number extracted
 * 	SecondReturn: Rest of the string
 */
func ExtractInt(src []rune) (string, []rune) {
	var num = ""

	for len(src) > 0 && IsInt(src[0]) {
		num += NextChar(&src)
	}

	return num, src
}

/*  FirstReturn: String extracted
 * 	SecondReturn: Rest of the string
 */
func ExtractAlpha(src []rune) (string, []rune) {
	var str = ""

	for len(src) > 0 && IsAlpha(src[0]) {
		str += NextChar(&src)
	}

	return str, src
}

func IsAlpha(char rune) bool {
	return strings.ToUpper(string(char)) != strings.ToLower(string(char))
}

/*  FirstReturn: Keyword extracted
 * 	SecondReturn: Is the keyword valid
 */
func IsKeyword(src string) (token_type.TokenType, bool) {
	keyword, ok := token_type.KEYWORDS[src]

	return keyword, ok
}

func IsInt(char rune) bool {
	val, err := strconv.Atoi(string(char))
	return err == nil && val >= 0
}
