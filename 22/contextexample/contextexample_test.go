package contextexample

import (
	"context"
	"fmt"
)

func ExampleGetContextType() {
	ctx := context.Background()
	fmt.Println(GetContextType(ctx))
	ctxCan, _ := context.WithCancel(ctx)
	fmt.Println(GetContextType(ctxCan))
	type key int
	type value string
	var key1 key = 1
	var value1 value = "banana"
	ctxVal := context.WithValue(ctx, key1, value1)
	fmt.Println(GetContextType(ctxVal))
	// Output:
	// *context.emptyCtx
	// *context.cancelCtx
	// *context.valueCtx
}

func ExampleGetContextErrorBeforeCancel() {
	ctx := context.Background()
	fmt.Println(GetContextErrorBeforeCancel(ctx))
	ctxCan, _ := context.WithCancel(ctx)
	fmt.Println(GetContextType(ctxCan))
	// Output:
	// <nil>
	// *context.cancelCtx
}

func ExampleGetContextErrorAfterCancel() {
	ctx := context.Background()
	fmt.Println(GetContextErrorAfterCancel(ctx))
	// Output:
	// context canceled
}



