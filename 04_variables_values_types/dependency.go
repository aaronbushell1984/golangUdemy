package hello

import "rsc.io/quote"

func HelloDependency() string {
    return quote.Hello()
}