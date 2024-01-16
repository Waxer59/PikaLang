package lexer

import (
	"errors"
	compilerErrors "github.com/Waxer59/PikaLang/internal/errors"
	"github.com/Waxer59/PikaLang/pkg/lexer/internal/utils"
	"github.com/Waxer59/PikaLang/pkg/lexer/token_type"
)

func Tokenize(input string) ([]token_type.Token, error) {
	var tokens []token_type.Token
	src := []rune(input)
	subtract := func(i int) rune {
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
		if utils.IsSkippable(tokenChar) {
			subtract(1)
			continue
		}

		// Check for number
		if utils.IsInt(tokenChar) {
			num, rest := utils.ExtractNum(src)
			tokens = append(tokens, token_type.Token{Type: token_type.Number, Value: num})
			src = rest
			continue
		}

		// Check for alpha
		if utils.IsAlpha(tokenChar) {
			alpha, rest := utils.ExtractIdentifier(src)
			alphaType := token_type.Identifier

			if keyword, ok := utils.IsKeyword(alpha); ok {
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
				subtract(2) // consume ' += '
				tokens = append(tokens, token_type.Token{Type: token_type.PlusEquals, Value: string(tokenChar) + "="})
				continue
			}

			if nextChar() == '+' {
				subtract(2) // consume ' ++ '
				tokens = append(tokens, token_type.Token{Type: token_type.Increment, Value: "++"})
				continue
			}

			tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: string(tokenChar)})
		case '%':
			if nextChar() == '=' {
				subtract(2) // consume ' %= '
				tokens = append(tokens, token_type.Token{Type: token_type.ModuleEquals, Value: string(tokenChar) + "="})
				continue
			}
			tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: string(tokenChar)})
		case '-':
			if nextChar() == '=' {
				subtract(2) // consume ' -= '
				tokens = append(tokens, token_type.Token{Type: token_type.MinusEquals, Value: string(tokenChar) + "="})
				continue
			}

			if nextChar() == '-' {
				subtract(2) // consume ' -- '
				tokens = append(tokens, token_type.Token{Type: token_type.Decrement, Value: "--"})
				continue
			}

			tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: string(tokenChar)})
		case '*':
			if nextChar() == '=' {
				subtract(2) // consume ' *= '
				tokens = append(tokens, token_type.Token{Type: token_type.TimesEquals, Value: string(tokenChar) + "="})
				continue
			}

			if nextChar() == '*' { // Check for power
				subtract(2) // consume ' ** '
				if src[0] == '=' {
					subtract(1) // advance '='
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
				subtract(2) // consume ' // '
				for len(src) > 0 && src[0] != '\n' {
					subtract(1)
					if len(src) <= 0 {
						break
					}
				}
			case '*':
				subtract(2) // consume /*
				for len(src) > 0 && src[0] != '*' && nextChar() != '/' {
					subtract(1)
					if len(src) <= 1 { // if the comment is not terminated
						return nil, errors.New(compilerErrors.ErrSyntaxUnterminatedMultilineComment)
					}
				}
				subtract(2) // consume */
			case '=':
				subtract(2) // consume '/='
				tokens = append(tokens, token_type.Token{Type: token_type.DivideEquals, Value: "/="})
			default:
				tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: string(tokenChar)})
			}
		case '=':
			switch nextChar() {
			case '=':
				subtract(2) // consume '=='
				tokens = append(tokens, token_type.Token{Type: token_type.EqualEqual, Value: "=="})
				continue
			case '>':
				subtract(2) // consume '=>'
				tokens = append(tokens, token_type.Token{Type: token_type.Arrow, Value: "=>"})
				continue
			default:
				tokens = append(tokens, token_type.Token{Type: token_type.Equals, Value: string(tokenChar)})
			}
		case '!':
			if nextChar() == '=' {
				subtract(2) // consume '!='
				tokens = append(tokens, token_type.Token{Type: token_type.NotEqual, Value: "!="})
				continue
			}
			tokens = append(tokens, token_type.Token{Type: token_type.Bang, Value: string(tokenChar)})
		case '>':
			if nextChar() == '=' {
				subtract(2) // consume '>='
				tokens = append(tokens, token_type.Token{Type: token_type.GreaterEqual, Value: ">="})
				continue
			}
			tokens = append(tokens, token_type.Token{Type: token_type.Greater, Value: string(tokenChar)})
		case '<':
			if nextChar() == '=' {
				subtract(2) // consume '<='
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
			if utils.IsInt(nextChar()) { // Check for decimal numbers as .123 == 0.123
				num, rest := utils.ExtractNum(src)
				tokens = append(tokens, token_type.Token{Type: token_type.Number, Value: num})
				src = rest
				continue
			}
			tokens = append(tokens, token_type.Token{Type: token_type.Dot, Value: string(tokenChar)})
		case '"':
			tokens = append(tokens, token_type.Token{Type: token_type.DoubleQuote, Value: string(tokenChar)})

			if utils.IsAlpha(nextChar()) {
				subtract(1) // consume '"'
				str, rest := utils.ExtractString(src)
				tokens = append(tokens, token_type.Token{Type: token_type.StringLiteral, Value: str})
				src = rest
				continue
			}
		case '\'':
			tokens = append(tokens, token_type.Token{Type: token_type.SingleQuote, Value: string(tokenChar)})
		case '|':
			if nextChar() == '|' {
				subtract(2) // consume '||'
				tokens = append(tokens, token_type.Token{Type: token_type.Or, Value: "||"})
			}
		case '&':
			if nextChar() == '&' {
				subtract(2) // consume '&&'
				tokens = append(tokens, token_type.Token{Type: token_type.And, Value: "&&"})
			}
		default:
			tokens = append(tokens, token_type.Token{Type: token_type.Identifier, Value: string(tokenChar)})
		}
		subtract(1)
	}
	tokens = append(tokens, token_type.Token{Type: token_type.EOF, Value: "EndOfFile"})
	return tokens, nil
}
