package util

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/karamaru-alpha/melt/pkg/merrors"
)

func EqualMeltError(t *testing.T, expected, actual error) bool {
	if expected == nil {
		return assert.NoError(t, actual)
	}
	if !assert.Error(t, expected) {
		return false
	}
	if !assert.Error(t, actual) {
		return false
	}

	merrExpected := &merrors.MeltError{}
	if !assert.ErrorAs(t, expected, &merrExpected) {
		return false
	}
	merrActual := &merrors.MeltError{}
	if !assert.ErrorAs(t, actual, &merrActual) {
		return false
	}

	if !assert.Equal(t, merrExpected.ErrorPattern, merrActual.ErrorPattern) {
		return false
	}
	if !assert.Equal(t, merrExpected.Message(), merrActual.Message()) {
		return false
	}
	return true
}
