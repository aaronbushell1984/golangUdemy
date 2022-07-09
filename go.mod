module github.com/aaronbushell1984/golangUdemy

go 1.17

require rsc.io/quote v1.5.2

// List - go list -m all - to list all dependencies
// Upgrade Latest - go get rsc.io/sampler - to get latest version
// Upgrade Specific - go get rsc.io/sampler@v1.3.1 - to get exact version
// Remove Unused - go mod tidy - clear unused dependencies
require (
	golang.org/x/text v0.3.7 // indirect
	rsc.io/sampler v1.3.1 // indirect
)
