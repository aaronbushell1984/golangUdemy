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
	fmt.Println(err)
	type key int
	type value string
	var key1 key = 1
	var value1 value = "banana"
	ctxVal := context.WithValue(ctx, key1, value1)
	fmt.Println(GetContextType(ctxVal))
	// Output:
	// *context.emptyCtx
	// *context.cancelCtx
	// 0x4a8d20
	// *context.valueCtx
}

func ExampleGetContextErrorBeforeCancel() {
	ctx := context.Background()
	fmt.Println(GetContextErrorBeforeCancel(ctx))
	ctxCan, err := context.WithCancel(ctx)
	fmt.Println(GetContextType(ctxCan))
	fmt.Println(err)
	// Output:
	// <nil>
	// *context.cancelCtx
	// 0x4a8d20
}

func ExampleGetContextErrorType() {
	ctx := context.Background()
	_, err := context.WithCancel(ctx)
	fmt.Println(GetContextErrorType(err))
	// Output:
	// context.CancelFunc
}


func ExampleGetContextErrorAfterCancel() {
	ctx := context.Background()
	fmt.Println(GetContextErrorAfterCancel(ctx))
	// Output:
	// context canceled
}



