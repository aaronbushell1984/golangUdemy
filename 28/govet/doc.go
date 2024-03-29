/*
This package just houses documentation on go go vet.

At https://pkg.go.dev/cmd/vet:

	"Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string.
	Vet uses heuristics that do not guarantee all reports are genuine problems, but it can find errors not caught by the compilers."

runing command:

	go help vet

gives output:

	usage: go vet [-n] [-x] [-vettool prog] [build flags] [vet flags] [packages]

	Vet runs the Go vet command on the packages named by the import paths.

	For more about vet and its flags, see 'go doc cmd/vet'.
	For more about specifying packages, see 'go help packages'.
	For a list of checkers and their flags, see 'go tool vet help'.
	For details of a specific checker such as 'printf', see 'go tool vet help printf'.

	The -n flag prints commands that would be executed.
	The -x flag prints commands as they are executed.

	The -vettool=prog flag selects a different analysis tool with alternative
	or additional checks.
	For example, the 'shadow' analyzer can be built and run using these commands:

	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
	go vet -vettool=$(which shadow)

	The build flags supported by go vet are those that control package resolution
	and execution, such as -n, -x, -v, -tags, and -toolexec.
	For more about these flags, see 'go help build'.

	See also: go fmt, go fix.
*/
package govet
