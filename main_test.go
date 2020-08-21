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

func TestCountOccurrences(t *testing.T) {
	test := []int{1, 2, 3, 1, 2, 3, 4, 5, 6}
	wanted := []int{2, 2, 2, 1, 1, 1, 0, 0, 0}

	result := countOccurrences(test)

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("countOccurrences(%d) = %d; want %d", test, result, wanted)
	}
}
