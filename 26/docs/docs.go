// Package docs demonstrates the use of package documentation in go
package docs

import "fmt"

func HelloWorld(name string) string {
	return fmt.Sprintf("Hello %v", name)
}