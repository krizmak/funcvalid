package funcvalid

import (
	"errors"
	"regexp"

	"golang.org/x/exp/constraints"
)

// The generic validator function type.
type Validator[T any] func(inp T) error

// Factory function with a paramter that returns a validator, that
// validates if the input value equals to the parameter.
func Eq[T comparable](pattern T) Validator[T] {
	return func(inp T) error {
		if inp == pattern {
			return nil
		}
		return errors.New("error: Eq")
	}
}

// Factory function with a parameter that returns a validator, that
// validates if the input value less than the parameter.
func Lt[T constraints.Ordered](pattern T) Validator[T] {
	return func(inp T) error {
		if inp < pattern {
			return nil
		}
		return errors.New("error: Lt")
	}
}

// Factory function with a parameter that returns a validator, that
// validates if the input value less than the parameter.
func Gt[T constraints.Ordered](pattern T) Validator[T] {
	return func(inp T) error {
		if inp > pattern {
			return nil
		}
		return errors.New("error: Gt")
	}
}

// Factory function with a regexp string parameter that returns a validator, that
// validates if the input value matches to the regexp.
func Regexp(pattern string) Validator[string] {
	return func(inp string) error {
		matched, err := regexp.MatchString(pattern, inp)
		if (err == nil) && (matched) {
			return nil
		}
		return errors.New("error: Regexp")
	}
}

// Factory function with a regexp parameter that returns a validator, that
// validates if the input value matches to the regexp.
func RegexpRE(pattern *regexp.Regexp) Validator[string] {
	return func(inp string) error {
		if pattern.MatchString(inp) {
			return nil
		}
		return errors.New("error: Regexp")
	}
}

// Factory function with a parameter that returns a validator, that
// validates if the length of the input array or string equals to the parameter.
func LenEq[T string | []T](length int) Validator[T] {
	return func(inp T) error {
		if len(inp) == length {
			return nil
		}
		return errors.New("error: LenEq")
	}
}

// Factory function with two parameters that returns a validator, that
// validates if the length of the input array or string is between the two parameters.
func LenBw[T string | []T](min int, max int) Validator[T] {
	return func(inp T) error {
		if (min <= len(inp)) && (len(inp) <= max) {
			return nil
		}
		return errors.New("error: LenBw")
	}
}

// Factory function with a parameter that returns a validator, that
// validates if the length of the input array or string less than the parameter.
func LenLt[T string | []T](length int) Validator[T] {
	return func(inp T) error {
		if len(inp) < length {
			return nil
		}
		return errors.New("error: LenLt")
	}
}

// Factory function with a parameter that returns a validator, that
// validates if the length of the input array or string greater than the parameter.
func LenGt[T string | []T](length int) Validator[T] {
	return func(inp T) error {
		if len(inp) > length {
			return nil
		}
		return errors.New("error: LenGt")
	}
}

// Factory function with a number of parameters that returns a validator, that validates
// if the input is one of the values in the parameters.
func OneOf[T comparable](elems ...T) Validator[T] {
	return func(inp T) error {
		for _, e := range elems {
			if e == inp {
				return nil
			}
		}
		return errors.New("error: OneOf")
	}
}

// Factory function with a map parameter that returns a validator, that validates
// if the input is one of the keys in the map.
func KeyIn[K comparable, V any](validmap map[K]V) Validator[K] {
	return func(inp K) error {
		if _, ok := validmap[inp]; ok {
			return nil
		}
		return errors.New("error: IsKeyIn")
	}
}

// Factory function with a map parameter that returns a validator, that validates
// if the input is one of the values in the map.
func ValueIn[K comparable, V comparable](validmap map[K]V) Validator[V] {
	return func(inp V) error {
		for _, value := range validmap {
			if value == inp {
				return nil
			}
		}
		return errors.New("error: IsValueIn")
	}
}

// Factory function with a string parameter that returns a validator, that always return
// error with the message in the parameter.
func ErrorValidator[T any](error_msg string) Validator[T] {
	return func(inp T) error {
		return errors.New("error: " + error_msg)
	}
}

// Factory function that takes variable number of validators, and returns a validator
// that validates if all the parameter validators are valid.
func Not[T any](validator Validator[T]) Validator[T] {
	return func(inp T) error {
		if err := validator(inp); err != nil {
			return nil
		}
		return errors.New("error: Not")
	}
}

// Factory function that takes variable number of validators, and returns a validator
// that validates if all the parameter validators are valid.
func And[T any](validators ...Validator[T]) Validator[T] {
	return func(inp T) error {
		for _, v := range validators {
			if err := v(inp); err != nil {
				return err
			}
		}
		return nil
	}
}

// Factory function that takes variable number of validators, and returns a validator
// that validates if any of the parameter validators are valid.
func Or[T any](validators ...Validator[T]) Validator[T] {
	return func(inp T) error {
		for _, v := range validators {
			if err := v(inp); err == nil {
				return nil
			}
		}
		return errors.New("error: Or")
	}
}

// Helper function that takes variable number of errors or nils, and returns an err
// if any of the params is not nil.
func AnyErr(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

// Interface with a single function that validates the type. A simple example
// struc that implements the Validable interface:
//
//	   type LoginReqData struct {
//		      Username string
//		      Password string
//	     }
//
//	   func (l LoginReqData) validate() error {
//	   	return fv.AnyErr(
//			      fv.LenBw[string](1, 30)(l.Username), //Username length should be between 1 and 30
//			      fv.LenBw[string](7, 32)(l.Password)) //Password length should be between 7 and 32
//	   }
type Validable interface {
	validate() error
}
