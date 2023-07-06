package interpreter_env

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type NativeFunction func(args []RuntimeValue, env Environment) RuntimeValue

var NativeFunctions = map[string]NativeFunction{
	"print": func(args []RuntimeValue, env Environment) RuntimeValue {
		for _, arg := range args {
			fmt.Println(arg.GetValue(), " ")
		}
		return NullVal{
			Type:  Null,
			Value: nil,
		}
	},
	"printe": func(args []RuntimeValue, env Environment) RuntimeValue {
		for _, arg := range args {
			s := fmt.Sprintf("%v", arg.GetValue())
			color.Red(s)
		}
		return NullVal{
			Type:  Null,
			Value: nil,
		}
	},
	"len": func(args []RuntimeValue, env Environment) RuntimeValue {
		if len(args) < 1 {
			return NaNVal{
				Type:  Number,
				Value: "NaN",
			}
		}

		arg, ok := args[0].GetValue().(string)
		if !ok {
			return NullVal{
				Type:  Null,
				Value: nil,
			}
		}
		return NumberVal{
			Type:  Number,
			Value: float64(len(arg)),
		}
	},
	"typeof": func(args []RuntimeValue, env Environment) RuntimeValue {
		if len(args) < 1 {
			return NaNVal{
				Type:  Number,
				Value: "NaN",
			}
		}

		return StringVal{
			Type:  String,
			Value: string(args[0].GetType()),
		}
	},
	"string": func(args []RuntimeValue, env Environment) RuntimeValue {
		if len(args) < 1 {
			return StringVal{
				Type:  String,
				Value: "",
			}
		}

		switch args[0].GetType() {
		case Null:
			return StringVal{
				Type:  String,
				Value: "null",
			}
		case Object:
			return StringVal{
				Type:  String,
				Value: "object",
			}
		default:
			s := fmt.Sprintf("%v", args[0].GetValue())
			return StringVal{
				Type:  String,
				Value: s,
			}
		}
	},
	"num": func(args []RuntimeValue, env Environment) RuntimeValue {
		if len(args) < 1 {
			return NaNVal{
				Type:  Number,
				Value: "NaN",
			}
		}

		i, err := strconv.ParseFloat(args[0].GetValue().(string), 64)

		if err != nil {
			return NaNVal{
				Type:  Number,
				Value: "NaN",
			}
		}

		return NumberVal{
			Type:  Number,
			Value: i,
		}
	},
	"bool": func(args []RuntimeValue, env Environment) RuntimeValue {
		if len(args) < 1 {
			return BooleanVal{
				Type:  Boolean,
				Value: false,
			}
		}

		result := false
		switch v := args[0].GetValue().(type) {
		case bool:
			result = v
		case int, float64, float32:
			result = v != 0.0
		case string:
			result = v != ""
		case nil:
			result = false
		default:
			result = true
		}

		return BooleanVal{
			Type:  Boolean,
			Value: result,
		}
	},
	"isNaN": func(args []RuntimeValue, env Environment) RuntimeValue {
		if len(args) < 1 {
			return BooleanVal{
				Type:  Boolean,
				Value: false,
			}
		}

		return BooleanVal{
			Type:  Boolean,
			Value: args[0].GetValue() == "NaN" && args[0].GetType() == Number,
		}
	},
	"isNull": func(args []RuntimeValue, env Environment) RuntimeValue {
		if len(args) < 1 {
			return BooleanVal{
				Type:  Boolean,
				Value: false,
			}
		}

		return BooleanVal{
			Type:  Boolean,
			Value: args[0].GetValue() == nil,
		}
	},
	"randNum": func(args []RuntimeValue, env Environment) RuntimeValue {
		if len(args) <= 1 {
			return NaNVal{
				Type:  Number,
				Value: "NaN",
			}
		}

		min := int(args[0].GetValue().(float64))
		max := int(args[1].GetValue().(float64))

		if min > max {
			return NaNVal{
				Type:  Number,
				Value: "NaN",
			}
		}
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(max-min+1) + min
		return NumberVal{
			Type:  Number,
			Value: float64(num),
		}
	},
	"prompt": func(args []RuntimeValue, env Environment) RuntimeValue {
		if len(args) < 1 {
			return StringVal{
				Type:  String,
				Value: "",
			}
		}

		fmt.Print(args[0].GetValue())
		var input string
		fmt.Scanln(&input)
		return StringVal{
			Type:  String,
			Value: input,
		}
	},
	"pow": func(args []RuntimeValue, env Environment) RuntimeValue {
		if len(args) <= 1 {
			return NaNVal{
				Type:  Number,
				Value: "NaN",
			}
		}

		base := args[0].GetValue().(float64)
		exponent := args[1].GetValue().(float64)
		result := math.Pow(base, exponent)
		return NumberVal{
			Type:  Number,
			Value: result,
		}
	},
	"toUpperCase": func(args []RuntimeValue, env Environment) RuntimeValue {
		if len(args) < 1 {
			return StringVal{
				Type:  String,
				Value: "",
			}
		}

		str := args[0].GetValue().(string)
		result := strings.ToUpper(str)
		return StringVal{
			Type:  String,
			Value: result,
		}
	},
	"toLowerCase": func(args []RuntimeValue, env Environment) RuntimeValue {
		if len(args) < 1 {
			return StringVal{
				Type:  String,
				Value: "",
			}
		}

		str := args[0].GetValue().(string)
		result := strings.ToLower(str)
		return StringVal{
			Type:  String,
			Value: result,
		}
	},
}
