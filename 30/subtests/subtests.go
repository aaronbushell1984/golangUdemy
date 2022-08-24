// Package subtests demonstrates grouping tests into subtests
package subtests

// Hello returns a language dependant prefix followed by provided name
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

// greetingPrefix sets a "Hello, " prefix based on provided language. Defaulting to Hello if language is not recognised.
func greetingPrefix(language string) (prefix string) {
	//goland:noinspection SpellCheckingInspection
	const (
		defaultPrefix = "Hello, "
		spanishPrefix = "Hola, "
		frenchPrefix  = "Bonjour, "
		spanish       = "Spanish"
		french        = "French"
	)
	prefix = defaultPrefix
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		return defaultPrefix
	}
	return
}
