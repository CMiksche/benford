# benford

[![Build Status](https://cloud.drone.io/api/badges/CMiksche/benford/status.svg)](https://cloud.drone.io/CMiksche/benford)
[![Go Report Card](https://goreportcard.com/badge/github.com/CMiksche/benford)](https://goreportcard.com/report/github.com/CMiksche/benford)

Check if a number array confirms to the Benford's Law

## Install

    go get -u github.com/CMiksche/benford

## Usage

Look at the tests for good examples.

E.g.:

````go
import (
    "github.com/CMiksche/benford"
)

func example() {
    benfordResult := benford.CalcBenfords([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
    chiSquared := benfordResult.ChiSquared
    // 0.40105320411553363
    dist := benfordResult.Dist
    // distribution array of the numbers 1 to 9 in the input number array
}
````

The CalcBenfords().ChiSquared result is a float and describes how well Benford's Law was matched. Lower is better.

CalcBenfords().Dist returns the distribution of the numbers 1 to 9 in the input array.

## Development

Install pre-commit hooks:

    pre-commit install

Test program:

    go test -v

## See also

* https://blog.m5e.de/post/benchmarking-go-crystal-python-and-js/ Benchmarking Benford's Law in different Programming Languages
* https://github.com/CMiksche/benfordslaw A JavaScript / TypeScript npm Package implementation of this benford's law check
* https://github.com/CMiksche/benfordslaw.cr A Crystal implementation of this benford's law check
