// Package contextexample demonstrates using context to control go routines
package contextexample

import (
	"context"
	"fmt"
	"time"
)

// GetContextType reurns the type of a context
//
// The example shows type of:
//	context.Background()			Type: *context.emptyCtx
//	context.WithCancel(ctx)			Type: *context.cancelCtx
//	context.WithValue(ctx, key1, value1)	Type: *context.valueCtx
// The type of the error from context with cancel:
//	context.WithCancel(ctx)			Type: context.CancelFunc
// is provided by a different function -- GetContextErrorType(err)
func GetContextType(ctx context.Context) string {
	return fmt.Sprintf("%T", ctx)
}

// GetContextErrorType returns the type of a context error
func GetContextErrorType(err context.CancelFunc) string {
	return fmt.Sprintf("%T", err)
}

// GetContextError returns the error from a context
func GetContextError(ctx context.Context) string {
	return fmt.Sprintf("%v", ctx.Err())
}

// CancelAndReturnErrorContext assigns a context to context with cancel:
//	ctx, cancel := context.WithCancel(ctx)
// cancels the context:
//	cancel()
// and returnns the value of the error as a string:
//	return fmt.Sprintf("%v", ctx.Err())
func CancelAndReturnErrorContext(ctx context.Context) string {
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	return fmt.Sprintf("%v", ctx.Err())
}

func CancelGoRoutineWithContext() []int {
	ctx, cancel := context.WithCancel(context.Background())
	var numbers []int
	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				numbers = append(numbers, n)
			}
		}
	}()
	time.Sleep(time.Second * 2)
	cancel()
	return numbers
}