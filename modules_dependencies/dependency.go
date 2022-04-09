package modules_dependencies

import "rsc.io/quote"

func HelloDependency() string {
    return quote.Hello()
}