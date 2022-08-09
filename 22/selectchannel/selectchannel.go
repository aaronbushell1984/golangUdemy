// Package selectchannel demonstrates collecting values from different channels in go
package selectchannel

// ConsumeMultipleReadOnlyChannels allows consumption of odd and even numbers 0-9 onto two read only channels
//
// The third channel passes 10 as final value
//
// NB. See example for consuming these channels using select
func ConsumeMultipleReadOnlyChannels() (<-chan int, <-chan int, <-chan int) {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go func(){
		for i := 0; i < 10; i++ {
			if i % 2 == 0 {
				even <- i
			}
			if i % 2 != 0 {
				odd <- i
			}
		}
		quit <- 10
	}()
	return even, odd, quit
}