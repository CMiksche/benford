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

func TestCountOccurrences2(t *testing.T) {
	test := []int{3, 6, 6, 9, 4, 9, 1, 1, 8, 4, 7, 2, 5, 3, 1, 4, 8, 4, 2, 2, 6, 1}
	wanted := []int{4, 3, 2, 4, 1, 3, 1, 2, 2}

	result := countOccurrences(test)

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("countOccurrences(%d) = %d; want %d", test, result, wanted)
	}
}

func TestCountDistribution(t *testing.T) {
	test := []int{4, 3, 2, 4, 1, 3, 1, 2, 2}
	wanted := []float32{0.18181819, 0.13636364, 0.09090909, 0.18181819, 0.045454547, 0.13636364, 0.045454547, 0.09090909, 0.09090909}

	result := countDistribution(test)

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("countDistribution(%d) = %g; want %g", test, result, wanted)
	}
}
