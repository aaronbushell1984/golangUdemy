module github.com/aaronbushell1984/golangUdemy

go 1.18

require rsc.io/quote v1.5.2

require (
	github.com/yuin/goldmark v1.4.13 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/tools v0.1.12 // indirect
)

// List - go list -m all - to list all dependencies
// Upgrade Latest - go get rsc.io/sampler - to get latest version
// Upgrade Specific - go get rsc.io/sampler@v1.3.1 - to get exact version
// Remove Unused - go mod tidy - clear unused dependencies
require (
	golang.org/x/text v0.3.7 // indirect
	rsc.io/sampler v1.3.1 // indirect
)
