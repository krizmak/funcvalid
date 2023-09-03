package funcvalid

import (
	"errors"
	"regexp"

	"golang.org/x/exp/constraints"
)

type  Validator[T any] func(inp T) error

func Eq[T comparable](pattern T) Validator[T] {
   return func(inp T) error {
     if (inp != pattern) {
        return errors.New("StrEq error")
     }
     return nil
   }
}

func Lt[T constraints.Ordered](pattern T) Validator[T] {
   return func(inp T) error {
     if (inp < pattern) {
        return errors.New("StrEq error")
     }
     return nil
   }
}

func Regexp(pattern string) Validator[string] {
   return func(inp string) error {
      matched, err := regexp.MatchString(pattern, inp)
      if (err != nil) || (!matched) {
         return errors.New("Regexp error")
      }
      return nil
   }
}

func LenEq[T string | []T](length int) Validator[T] {
   return func (inp T) error {
      if (len(inp) == length) {
         return nil
      }
      return errors.New("length error: Eq")
   }
}

func LenBw[T string | []T](min int, max int) Validator[T] {
   return func (inp T) error {
      if (min <= len(inp)) && (len(inp) <= max) {
         return nil
      }
      return errors.New("length error: Bw")
   }
}

func LenLt[T string | []T](length int) Validator[T] {
   return func (inp T) error {
      if (len(inp) < length) {
         return nil
      }
      return errors.New("length error: Lt")
   }
}

func LenGt[T string | []T](length int) Validator[T] {
   return func (inp T) error {
      if (len(inp) > length) {
        return nil
      }
      return errors.New("length error: Gt")
   }
}

func And[T any](validators ...Validator[T]) Validator[T] {
   return func (inp T) error {
      for _, v := range validators {
        if err := v(inp); err != nil {
          return err
        }
      }
      return nil
   }
}

func Or[T any](validators ...Validator[T]) Validator[T] {
   return func (inp T) error {
      for _, v := range validators {
        if err := v(inp); err == nil {
          return nil
        }
      }
      return errors.New("All validator failed in StrOr")
   }
}

func AnyErr(errs ...error) error {
   for _, err := range errs {
      if err != nil {
        return err
      }
    }
    return nil
}

type Validable interface {
   validate() error
}