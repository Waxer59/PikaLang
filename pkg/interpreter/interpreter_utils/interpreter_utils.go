package interpreter_utils

import "pika/pkg/interpreter/interpreter_nativeFns"

/*
 * First return value is the function itself.
 * Second return value is true if the function exists.
 */
func IsNativeFunction(name string) (interpreter_nativeFns.NativeFunction, bool) {
	function, ok := interpreter_nativeFns.NativeFunctions[name]

	return function, ok
}
