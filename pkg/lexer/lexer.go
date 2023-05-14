package lexer

import (
	"pika/pkg/lexer/lexerTypes"
	"pika/pkg/lexer/lexerUtils"
	"strings"
)

func Tokenize(line string) []lexerTypes.Token {
	tokens := []lexerTypes.Token{}
	src := strings.Split(line, "")

	for len(src) > 0 {
		tokenStr := src[0]
		// Check if token is skippable
		if lexerUtils.IsSkippable(tokenStr) {
			src = src[1:]
			continue
		}

		// Check for keywords
		if keyword, ok := lexerTypes.KEYWORDS[tokenStr]; ok {
			tokens = append(tokens, lexerTypes.Token{Type: keyword, Value: tokenStr})
			continue
		}

		// Check for number
		if lexerUtils.IsInt(tokenStr) {
			num, rest := lexerUtils.ExtractInt(src)
			tokens = append(tokens, lexerTypes.Token{Type: lexerTypes.Number, Value: num})
			src = rest
			continue
		}

		// Check for alpha
		if lexerUtils.IsAlpha(tokenStr) {
			alpha, rest := lexerUtils.ExtractAlpha(src)
			alphaType := lexerTypes.Identifier

			if keyword, ok := lexerUtils.IsKeyword(alpha); ok {
				alphaType = keyword
			}

			tokens = append(tokens, lexerTypes.Token{Type: alphaType, Value: alpha})

			src = rest
			continue
		}

		// Check for operators
		switch tokenStr {
		case "+", "-", "*", "/", "%":
			tokens = append(tokens, lexerTypes.Token{Type: lexerTypes.BinaryOperator, Value: tokenStr})
		case "=":
			tokens = append(tokens, lexerTypes.Token{Type: lexerTypes.Equals, Value: tokenStr})
		case ";":
			tokens = append(tokens, lexerTypes.Token{Type: lexerTypes.SemiColon, Value: tokenStr})
		case "(":
			tokens = append(tokens, lexerTypes.Token{Type: lexerTypes.RightParen, Value: tokenStr})
		case ")":
			tokens = append(tokens, lexerTypes.Token{Type: lexerTypes.LeftParen, Value: tokenStr})
		case "{":
			tokens = append(tokens, lexerTypes.Token{Type: lexerTypes.LeftBrace, Value: tokenStr})
		case "}":
			tokens = append(tokens, lexerTypes.Token{Type: lexerTypes.RightBrace, Value: tokenStr})
		case ",":
			tokens = append(tokens, lexerTypes.Token{Type: lexerTypes.Comma, Value: tokenStr})
		case ":":
			tokens = append(tokens, lexerTypes.Token{Type: lexerTypes.Colon, Value: tokenStr})
		default:
			tokens = append(tokens, lexerTypes.Token{Type: lexerTypes.Identifier, Value: tokenStr})
		}

		// Remove token
		src = src[1:]
	}
	tokens = append(tokens, lexerTypes.Token{Type: lexerTypes.EOF, Value: "EndOfFile"})
	return tokens
}
