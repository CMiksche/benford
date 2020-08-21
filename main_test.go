package main

import (
	"reflect"
	"testing"
)

func TestFirstDigit(t *testing.T) {
	test := []int{123, 678, 987}
	wanted := []int{1, 6, 9}

	result := firstDigit(test)

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("firstDigit(%d) = %d; want %d", test, result, wanted)
	}
}

func TestFirstDigitWithZero(t *testing.T) {
	test := []int{123, 0, 678, 987, 0}
	wanted := []int{1, 6, 9}

	result := firstDigit(test)

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("firstDigit(%d) = %d; want %d", test, result, wanted)
	}
}
