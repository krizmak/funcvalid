# funcvalid

Validator in functional style

This package provides a functional style API to validating data.
It's type safe, meaning that a lot of possible errors are already detected at compile time, not runtime.

The basic idea is that validators for any datatype are just functions that return an error or nil.

If you want to construct your own validator, you can:

- construct parametrized validators with a technic called curried functions in functional programming, see https://medium.com/@meeusdylan/function-currying-in-go-a88672d6ebcf
- create is-validators (validator without parameters) by uncurrying exisiting validators
- construct new validator functions from the existing ones, by using special validators that combine other validators (like And, Or, Not). That's the approach to build validators for your complex types (e.g. structs, maps).

## Installation

Use:

    go get github.com/krizmak/funcvalid

Then import it into your code:

    import (fv "github.com/krizmak/funcvalid")

## Usage

See go doc for more details.

## Acknowledment

The package uses some files from the tag based validator package: https://github.com/go-playground/validator based on the kind permission of the author https://github.com/deankarn.

## Similar packages

Just after I had started working on this package, I found https://github.com/go-ozzo/ozzo-validation, that provides similar approach (type-safe validation). However the functional style API is unique to this package, and may have benefit for others, who like this style.

## How to Contribute

Make a pull request.

## License

See https://github.com/krizmak/funcvalid/blob/main/LICENSE (MIT license)

## Maintainers

The package is in an early phase, maintined only by the author https://github.com/krizmak. Any help is appreciated.
