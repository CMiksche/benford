# benford

Check the benford's law

## Install

    go get -u github.com/CMiksche/benford
    
## Usage

Look at the tests for good examples.

E.g.:

    import (
        "github.com/CMiksche/benford"
    )

    func example() {
        benfordResult := benford.CalcBenfords([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
        chiSquared := benfordResult.chiSquared
        // 0.40105320411553363
        dist := benfordResult.dist
        // distribution array of the numbers 1 to 9 in the input number array
    }

The ChiSquared result is a float and describes how well Benford's Law was matched. Lower is better.

The Probability describes how relevant ChiSquared is. It should be >= 0.9

getDist() returns the distribution of the numbers.

## Development

Install pre-commit hooks:

    pre-commit install

Test program:

    go test -v

## See also

https://github.com/CMiksche/benfordslaw A JavaScript / TypeScript npm Package implementation of this benford's law check
