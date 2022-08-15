package contextexample

import (
	"context"
	"fmt"
)

func ExampleGetContextType() {
	ctx := context.Background()
	fmt.Println(GetContextType(ctx))
	ctxCan, err := context.WithCancel(ctx)
	fmt.Println(GetContextType(ctxCan))
	fmt.Println(GetContextErrorType(err))
	type key int
	type value string
	var key1 key = 1
	var value1 value = "banana"
	ctxVal := context.WithValue(ctx, key1, value1)
	fmt.Println(GetContextType(ctxVal))
	// Output:
	// *context.emptyCtx
	// *context.cancelCtx
	// context.CancelFunc
	// *context.valueCtx
}

func ExampleGetContextError() {
	ctx := context.Background()
	fmt.Println(GetContextError(ctx))
	ctxCan, err := context.WithCancel(ctx)
	fmt.Println(GetContextType(ctxCan))
	fmt.Println(GetContextErrorType(err))
	// Output:
	// <nil>
	// *context.cancelCtx
	// context.CancelFunc
}

func ExampleGetContextErrorType() {
	ctx := context.Background()
	_, err := context.WithCancel(ctx)
	fmt.Println(GetContextErrorType(err))
	// Output:
	// context.CancelFunc
}


func ExampleCancelAndReturnErrorContext() {
	ctx := context.Background()
	fmt.Println(ctx.Err())
	fmt.Println(CancelAndReturnErrorContext(ctx))
	// Output:
	// <nil>
	// context canceled
}

func ExampleCancelGoRoutineWithContext() {
	fmt.Println(CancelGoRoutineWithContext())
	// Output:
	// [1 2 3 4 5 6 7 8 9]
}