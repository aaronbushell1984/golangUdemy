package capturereturns

import "fmt"

// fmt.Println - func Println(a ...any) (n int, err error)
// returns (n int, err error)
func CaptureReturnsPrint() {
	n, err := fmt.Println("Capturing return values from fmt.Println:")
	fmt.Println(n)
	fmt.Println(err)
}

// fmt.Println - func Println(a ...any) (n int, err error)
// return of err can be discarded with _ (blank identifier)
func CaptureReturnsVoidErrPrint() {
	n, _ := fmt.Println("Capturing return values from fmt.Println:")
	fmt.Println(n)
	//fmt.Println(err) this no throws compile error because err is _
}