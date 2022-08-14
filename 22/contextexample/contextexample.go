// Package contextexample demonstrates using context to control go routines
package contextexample

import (
	"context"
	"fmt"
)

func GetContextType(ctx context.Context) string {
	return fmt.Sprintf("%T", ctx)
}

func GetContextErrorBeforeCancel(ctx context.Context) string {
	return fmt.Sprintf("%v", ctx.Err())
}

func GetContextErrorAfterCancel(ctx context.Context) string {
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	return fmt.Sprintf("%v", ctx.Err())
}