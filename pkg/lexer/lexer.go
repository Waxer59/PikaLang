package lexer

import (
	"errors"

	compilerErrors "github.com/Waxer59/PikaLang/internal/errors"
	"github.com/Waxer59/PikaLang/pkg/lexer/token_type"
)

func Tokenize(input string) ([]token_type.Token, error) {
	tokens := []token_type.Token{}
	src := []rune(input)
	substract := func(i int) rune {
		if len(src) <= 0 {
			return 0
		}
		char := src[0]
		src = src[i:]
		return char
	}

	nextChar := func() rune {
		if len(src) <= 1 {
			return 0
		}
		return src[1]
	}

	for len(src) > 0 {
		tokenChar := src[0]

		// Check if token is skippable
		if IsSkippable(tokenChar) {
			substract(1)
			continue
		}

		// Check for number
		if IsInt(tokenChar) {
			num, rest := ExtractNum(src)
			tokens = append(tokens, token_type.Token{Type: token_type.Number, Value: num})
			src = rest
			continue
		}

		// Check for alpha
		if IsAlpha(tokenChar) {
			alpha, rest := ExtractIdentifier(src)
			alphaType := token_type.Identifier

			if keyword, ok := IsKeyword(alpha); ok {
				alphaType = keyword
			}

			tokens = append(tokens, token_type.Token{Type: alphaType, Value: alpha})

			src = rest
			continue
		}

		// Check for operators
		switch tokenChar {
		case '+':
			if nextChar() == '=' {
				substract(2) // consume ' += '
				tokens = append(tokens, token_type.Token{Type: token_type.PlusEquals, Value: string(tokenChar) + "="})
				continue
			}

			if nextChar() == '+' {
				substract(2) // consume ' ++ '
				tokens = append(tokens, token_type.Token{Type: token_type.Increment, Value: "++"})
				continue
			}

			tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: string(tokenChar)})
		case '%':
			if nextChar() == '=' {
				substract(2) // consume ' %= '
				tokens = append(tokens, token_type.Token{Type: token_type.ModuleEquals, Value: string(tokenChar) + "="})
				continue
			}
			tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: string(tokenChar)})
		case '-':
			if nextChar() == '=' {
				substract(2) // consume ' -= '
				tokens = append(tokens, token_type.Token{Type: token_type.MinusEquals, Value: string(tokenChar) + "="})
				continue
			}

			if nextChar() == '-' {
				substract(2) // consume ' -- '
				tokens = append(tokens, token_type.Token{Type: token_type.Decrement, Value: "--"})
				continue
			}

			tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: string(tokenChar)})
		case '*':
			if nextChar() == '=' {
				substract(2) // consume ' *= '
				tokens = append(tokens, token_type.Token{Type: token_type.TimesEquals, Value: string(tokenChar) + "="})
				continue
			}

			if nextChar() == '*' { // Check for power
				substract(2) // consume ' ** '
				if src[0] == '=' {
					substract(1) // advance '='
					tokens = append(tokens, token_type.Token{Type: token_type.PowerEquals, Value: "**="})
					continue
				}
				tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: "**"})
				continue
			}

			tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: string(tokenChar)})
		case '/':
			switch nextChar() {
			case '/':
				substract(2) // consume ' // '
				for len(src) > 0 && src[0] != '\n' {
					substract(1)
					if len(src) <= 0 {
						break
					}
				}
			case '*':
				substract(2) // consume /*
				for len(src) > 0 && src[0] != '*' && nextChar() != '/' {
					substract(1)
					if len(src) <= 1 { // if the comment is not terminated
						return nil, errors.New(string(compilerErrors.ErrSyntaxUnterminatedMultilineComment))
					}
				}
				substract(2) // consume */
			case '=':
				substract(2) // consume '/='
				tokens = append(tokens, token_type.Token{Type: token_type.DivideEquals, Value: "/="})
			default:
				tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: string(tokenChar)})
			}
		case '=':
			switch nextChar() {
			case '=':
				substract(2) // consume '=='
				tokens = append(tokens, token_type.Token{Type: token_type.EqualEqual, Value: "=="})
				continue
			case '>':
				substract(2) // consume '=>'
				tokens = append(tokens, token_type.Token{Type: token_type.Arrow, Value: "=>"})
				continue
			default:
				tokens = append(tokens, token_type.Token{Type: token_type.Equals, Value: string(tokenChar)})
			}
		case '!':
			if nextChar() == '=' {
				substract(2) // consume '!='
				tokens = append(tokens, token_type.Token{Type: token_type.NotEqual, Value: "!="})
				continue
			}
			tokens = append(tokens, token_type.Token{Type: token_type.Bang, Value: string(tokenChar)})
		case '>':
			if nextChar() == '=' {
				substract(2) // consume '>='
				tokens = append(tokens, token_type.Token{Type: token_type.GreaterEqual, Value: ">="})
				continue
			}
			tokens = append(tokens, token_type.Token{Type: token_type.Greater, Value: string(tokenChar)})
		case '<':
			if nextChar() == '=' {
				substract(2) // consume '<='
				tokens = append(tokens, token_type.Token{Type: token_type.LessEqual, Value: "<="})
				continue
			}
			tokens = append(tokens, token_type.Token{Type: token_type.Less, Value: string(tokenChar)})
		case ';':
			tokens = append(tokens, token_type.Token{Type: token_type.Semicolon, Value: string(tokenChar)})
		case '?':
			tokens = append(tokens, token_type.Token{Type: token_type.QuestionMark, Value: string(tokenChar)})
		case '(':
			tokens = append(tokens, token_type.Token{Type: token_type.LeftParen, Value: string(tokenChar)})
		case ')':
			tokens = append(tokens, token_type.Token{Type: token_type.RightParen, Value: string(tokenChar)})
		case '{':
			tokens = append(tokens, token_type.Token{Type: token_type.LeftBrace, Value: string(tokenChar)})
		case '}':
			tokens = append(tokens, token_type.Token{Type: token_type.RightBrace, Value: string(tokenChar)})
		case '[':
			tokens = append(tokens, token_type.Token{Type: token_type.LeftBracket, Value: string(tokenChar)})
		case ']':
			tokens = append(tokens, token_type.Token{Type: token_type.RightBracket, Value: string(tokenChar)})
		case ',':
			tokens = append(tokens, token_type.Token{Type: token_type.Comma, Value: string(tokenChar)})
		case ':':
			tokens = append(tokens, token_type.Token{Type: token_type.Colon, Value: string(tokenChar)})
		case '.':
			if IsInt(nextChar()) { // Check for decimal numbers as .123 == 0.123
				num, rest := ExtractNum(src)
				tokens = append(tokens, token_type.Token{Type: token_type.Number, Value: num})
				src = rest
				continue
			}
			tokens = append(tokens, token_type.Token{Type: token_type.Dot, Value: string(tokenChar)})
		case '"':
			tokens = append(tokens, token_type.Token{Type: token_type.DoubleQoute, Value: string(tokenChar)}) // Append double qoute
			substract(1)

			var str string
			for len(src) > 0 && src[0] != '"' {
				if src[0] == '\\' {
					substract(1)
					str += string(substract(1))
					continue
				}
				str += string(substract(1))
			}

			tokens = append(tokens, token_type.Token{Type: token_type.StringLiteral, Value: str})
			tokens = append(tokens, token_type.Token{Type: token_type.DoubleQoute, Value: string(tokenChar)})
		case '\'':
			tokens = append(tokens, token_type.Token{Type: token_type.SingleQoute, Value: string(tokenChar)})
		case '|':
			if nextChar() == '|' {
				substract(2) // consume '||'
				tokens = append(tokens, token_type.Token{Type: token_type.Or, Value: "||"})
			}
		case '&':
			if nextChar() == '&' {
				substract(2) // consume '&&'
				tokens = append(tokens, token_type.Token{Type: token_type.And, Value: "&&"})
			}
		default:
			tokens = append(tokens, token_type.Token{Type: token_type.Identifier, Value: string(tokenChar)})
		}
		substract(1)
	}
	tokens = append(tokens, token_type.Token{Type: token_type.EOF, Value: "EndOfFile"})
	return tokens, nil
}
