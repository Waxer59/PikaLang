package interpreter_env

/*
 * First return value is the function itself.
 * Second return value is true if the function exists.
 */
func IsNativeFunction(name string) (NativeFunction, bool) {
	function, ok := NativeFunctions[name]

	return function, ok
}
