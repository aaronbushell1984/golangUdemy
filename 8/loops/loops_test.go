package loops

import (
	"fmt"
	"testing"
)

func ExamplePrintNumbersLoop() {
	fmt.Println(PrintNumbersLoop(0, 1))
	fmt.Println(PrintNumbersLoop(0, 10))
	fmt.Println(PrintNumbersLoop(-10, 10))
	// Output:
	// [0 1]
	// [0 1 2 3 4 5 6 7 8 9 10]
	// [-10 -9 -8 -7 -6 -5 -4 -3 -2 -1 0 1 2 3 4 5 6 7 8 9 10]
}

func BenchmarkPrintNumbersLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintNumbersLoop(-10, 10)
	}
}

func ExampleReturnAllTwoBitOptions() {
	fmt.Println(ReturnAllTwoBitOptions())
	// Output:
	// [00 01 10 11]
}

func BenchmarkReturnAllTwoBitOptions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReturnAllTwoBitOptions()
	}
}

func TestReturnNumbersOneToTen(t *testing.T) {
	want := "[1 2 3 4 5 6 7 8 9 10]"
	if got := ReturnNumbersOneToTen(); got != want {
		t.Errorf("ReturnNumbersOneToTen(): got: %q want: %q", got, want)
	}
}

func BenchmarkReturnNumbersOneToTen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReturnNumbersOneToTen()
	}
}

func TestReturnNumbersOneToTenWithBreak(t *testing.T) {
	want := "[1 2 3 4 5 6 7 8 9 10]"
	if got := ReturnNumbersOneToTenWithBreak(); got != want {
		t.Errorf("ReturnNumbersOneToTenWithBreak: got: %q want: %q", got, want)
	}
}

func BenchmarkReturnNumbersOneToTenWithBreak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReturnNumbersOneToTenWithBreak()
	}
}

func ExampleReturnEvenNumbersWithContinue() {
	fmt.Println(ReturnEvenNumbersWithContinue(0, 7))
	fmt.Println(ReturnEvenNumbersWithContinue(-10, 25))
	fmt.Println(ReturnEvenNumbersWithContinue(10, -24))
	fmt.Println(ReturnEvenNumbersWithContinue(10, 10))
	// Output:
	// [2 4 6]
	// [-10 -8 -6 -4 -2 2 4 6 8 10 12 14 16 18 20 22 24]
	// End number must be greater or equal to start number
	// [10]
}

func ExampleReturnAsciiNumbersAsCharacters() {
	fmt.Println(ReturnAsciiNumbersAsCharacters(-10, 25))
	fmt.Println(ReturnAsciiNumbersAsCharacters(33, 122))
	fmt.Println(ReturnAsciiNumbersAsCharacters(10, -24))
	// Output:
	// There are no negative ascii characters
	// [! " # $ % & ' ( ) * + , - . / 0 1 2 3 4 5 6 7 8 9 : ; < = > ? @ A B C D E F G H I J K L M N O P Q R S T U V W X Y Z [ \ ] ^ _ ` a b c d e f g h i j k l m n o p q r s t u v w x y z]
	// End number must be greater or equal to start number
}
