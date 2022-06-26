// Package dependency demonstrates syntax used to import a package
package dependency

import "rsc.io/quote"

// HelloDependency returns "Hello World" using rsc.io/quote.Hello
func HelloDependency() string {
	return quote.Hello()
}
