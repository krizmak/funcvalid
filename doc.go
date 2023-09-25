/*
Package funcvalid is a tiny validator package in functional style.

Unlike the usual go validators, it does not rely on the tag feature of the structs, because they
don't benefit from the type safety that golang types ensures.

So funcvalid package:

  - provides a type-safe validation framework
  - in a relatively dense format.

The main concept is that a validator is simply a function that takes an input and returns an error if
the validation is unsuccessful or nil, if it is successful.

	  // Define the validator function type for ints
	  type IntValidator func(input int) error

		// Here we create a validator function for ints (i.e. its type is IntValidator)
		func IntEq42(input int) error {
			if (input == 42) {
				return nil
			}
			return errors.New(fmt.Sprintf("Input %d doesn't answer to life the universe and everything.")
		}

		// Now we can use the validator this way:
		IntEq42(42)  // -> nil
		IntEq42(0)   // -> error

It doesn't seem very useful until now, because in most of the cases the validators need a parameter, too.

This package provides factory functions, that may take arguments and return a validator. It works a bit like
currying in functional languages hence the name. The idea comes from this excellent [curry post].

If it still sounds a bit cryptic, let's see the following example:

	// We create a general validator generator (factory function) that takes a parameter, and it will return
	// the actual validator function:
	func IntEq(pattern int) IntValidator {
		// Here we create the validator function (i.e. its type is IntValidator)
		// as it's a closure, it can use the pattern parameter, too.
		return func(input int) error {
			if (input == pattern) {
				return nil
			}
			return errors.New(fmt.Sprintf("Input %d doesn't answer to life the universe and everything.")
		}
	}

	// Now the usage is like:
	IntEq(42)(42)  // -> nil
	IntEq(42)(0)   // -> error

The actual factory functions are even a bit more general by using generic types. The package also provides
Not, Or and And factory functions with that you can chain other validators (see funcvalid.go).

The package also contains built-in validators from the [validator] package, that is the best of the
tag based validators (see validator_builtin.go).

For further details see the documentation of the functions, and have a functional fun in Go!

[curry post]: https://medium.com/@meeusdylan/function-currying-in-go-a88672d6ebcf
[validator]: https://github.com/go-playground/validator
*/
package funcvalid
