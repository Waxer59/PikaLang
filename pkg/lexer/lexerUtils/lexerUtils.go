package lexerUtils

import (
	"pikalang/pkg/lexer/lexerTypes"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func IsSkippable(char string) bool {
	return slices.Contains(lexerTypes.SkippableChars, char)
}

func NextChar(src *[]string) string {
	char := (*src)[0]
	*src = (*src)[1:]
	return string(char)
}

/*  FirstReturn: Number extracted
 * 	SecondReturn: Rest of the string
 **/
func ExtractInt(src []string) (string, []string) {
	var num = ""

	for len(src) > 0 && IsInt(src[0]) {
		num += NextChar(&src)
	}

	return num, src
}

/*  FirstReturn: String extracted
 * 	SecondReturn: Rest of the string
 **/
func ExtractAlpha(src []string) (string, []string) {
	var str = ""

	for len(src) > 0 && IsAlpha(src[0]) {
		str += NextChar(&src)
	}

	return str, src
}

func IsAlpha(char string) bool {
	return strings.ToUpper(char) != strings.ToLower(char)
}

/*  FirstReturn: Keyword extracted
 * 	SecondReturn: Is the keyword valid
 **/
func IsKeyword(src string) (lexerTypes.TokenType, bool) {
	keyword, ok := lexerTypes.KEYWORDS[src]

	return keyword, ok
}

func IsInt(char string) bool {
	val, err := strconv.Atoi(char)
	return err == nil && val >= 0
}
