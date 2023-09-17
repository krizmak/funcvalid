package funcvalid_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	fv "github.com/krizmak/funcvalid"
)

func TestFuncValid(t *testing.T) {
	err := fv.Eq("test")("test")
	assert.Equal(t, err, nil)

	err = fv.Eq(2.3)(2.3)
	assert.Equal(t, err, nil)

	validator := fv.And(fv.LenGt[string](1), fv.LenLt[string](5))
	err = validator("testelek")
	assert.NotEqual(t, err, nil)
	err = validator("t")
	assert.NotEqual(t, err, nil)
	err = validator("test")
	assert.Equal(t, err, nil)

	assert.Equal(t, fv.Regexp("b.*")("beta"), nil)
	assert.Equal(t, fv.Or(fv.Regexp("a.*"), fv.Regexp("b.*"))("alma"), nil)
	assert.NotEqual(t, fv.Or(fv.Regexp("d.*"), fv.Regexp("c.*"))("beta"), nil)

	assert.Equal(t, fv.Eq[uint](32)(32), nil)


	assert.Equal(t, fv.OneOf(1,2,3)(1), nil)
	assert.NotEqual(t, fv.OneOf(1,2,3)(4), nil)
	
}
