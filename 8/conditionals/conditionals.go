// TODO package docs
package conditionals

// TODO func docs
func ReturnSuppliedNameIfOverThreeCharactersCase(name string) string {
	switch {
	case len(name) < 4:
		return "Name must be over three characters"
	}
	return name
}

// TODO func docs
func ReturnSuppliedNameIfOverThreeCharactersInitializationStatement(name string) string {
	if minimum := 4; len(name) < minimum {
		return "Name must be over three characters"
	}
	return name
}
