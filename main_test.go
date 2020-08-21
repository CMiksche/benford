package benford

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
	wanted := []float64{
		0.18181818181818182, 0.13636363636363635,
		0.09090909090909091, 0.18181818181818182,
		0.045454545454545456, 0.13636363636363635,
		0.045454545454545456, 0.09090909090909091,
		0.09090909090909091}

	result := countDistribution(test)

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("countDistribution(%d) = %g; want %g", test, result, wanted)
	}
}

func TestCalcBenford(t *testing.T) {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	wanted := 0.40105320411553363

	result := CalcBenfords(test).chiSquared

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("CalcBenfords(%d) = %v; want %v", test, result, wanted)
	}
}

func TestCalcBenfordWithZeroes(t *testing.T) {
	test := []int{0, 1, 0, 2, 3, 4, 0, 5, 6, 7, 8, 9, 0}
	wanted := 0.40105320411553363

	result := CalcBenfords(test).chiSquared

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("CalcBenfords(%d) = %v; want %v", test, result, wanted)
	}
}

func TestCalcBenfordBackwards(t *testing.T) {
	test := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	wanted := 0.40105320411553363

	result := CalcBenfords(test).chiSquared

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("CalcBenfords(%d) = %v; want %v", test, result, wanted)
	}
}

func TestCalcBenfordTwoDigit(t *testing.T) {
	test := []int{12, 25, 38, 40, 54, 63, 75, 83, 96}
	wanted := 0.40105320411553363

	result := CalcBenfords(test).chiSquared

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("CalcBenfords(%d) = %v; want %v", test, result, wanted)
	}
}

func TestCalcBenfordThreeDigit(t *testing.T) {
	test := []int{122, 256, 382, 402, 546, 632, 757, 832, 964}
	wanted := 0.40105320411553363

	result := CalcBenfords(test).chiSquared

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("CalcBenfords(%d) = %v; want %v", test, result, wanted)
	}
}

func TestCalcBenfordBiggerArray(t *testing.T) {
	test := []int{1, 1, 1, 1, 1, 1, 1, 1,
		2, 2, 2, 2,
		3, 3, 3,
		4, 4,
		5, 5,
		6, 6,
		7, 7,
		8,
		9}
	wanted := 0.019868294035033682

	result := CalcBenfords(test).chiSquared

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("CalcBenfords(%d) = %v; want %v", test, result, wanted)
	}
}

func TestCalcBenfordRandom(t *testing.T) {
	test := []int{3, 6, 6, 9, 4, 9, 1, 1, 8, 4, 7, 2, 5, 3, 1, 4, 8, 4, 2, 2, 6, 1}
	wanted := 0.3034231881809182

	result := CalcBenfords(test).chiSquared

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("CalcBenfords(%d) = %v; want %v", test, result, wanted)
	}
}

func TestCalcBenfordBiggerRandom(t *testing.T) {
	test := []int{1, 2, 2, 2, 5, 6, 3, 8, 2, 4, 0, 2, 5, 4, 6, 6, 3, 2, 7, 5, 7, 8, 3, 2, 9, 6, 4, 1, 3, 2, 6, 9, 4, 1, 4, 7, 4, 9, 2, 1, 4, 9, 2, 6, 8, 3, 8, 9, 3, 8, 3, 5}
	wanted := 0.32533042485042685

	result := CalcBenfords(test).chiSquared

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("CalcBenfords(%d) = %v; want %v", test, result, wanted)
	}
}

func TestCalcBenfordWithDist(t *testing.T) {
	test := []int{3, 6, 6, 9, 4, 9, 1, 1, 8, 4, 7, 2, 5, 3, 1, 4, 8, 4, 2, 2, 6, 1}
	wanted := []float64{
		0.18181818181818182, 0.13636363636363635,
		0.09090909090909091, 0.18181818181818182,
		0.045454545454545456, 0.13636363636363635,
		0.045454545454545456, 0.09090909090909091,
		0.09090909090909091}

	result := CalcBenfords(test).dist

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("CalcBenfords(%d) = %v; want %v", test, result, wanted)
	}
}
