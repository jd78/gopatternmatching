package patternmatch

import (
	"errors"

	"github.com/google/go-cmp/cmp"
	gooptional "github.com/jd78/go-optional"
)

type match[T any] struct {
	input     T
	isMatched bool
}

type matchResult[T, K any] struct {
	input     T
	output    K
	isMatched bool
}

// Match condition and run an action that does not return a type
func Match[T any](input T) match[T] {
	return match[T]{input, false}
}

// when param is "func(input T) bool" used to match a condition
// a param is "func()" used to run an action
func (m match[T]) When(f func(input T) bool, a func()) match[T] {
	if m.isMatched {
		return m
	}

	if f(m.input) {
		a()
		m.isMatched = true
	}

	return m
}

// val param is "T" used to exact match the input with the passed condition
// a param is "func()" used to run an action
func (m match[T]) WhenValue(val T, a func()) match[T] {
	if m.isMatched {
		return m
	}

	if cmp.Equal(m.input, val) {
		a()
		m.isMatched = true
	}

	return m
}

// OtherwisePanic panic if the pattern is not matched, optionally used at the end of the pattern matching
func (m match[T]) OtherwisePanic() {
	if !m.isMatched {
		panic("pattern not matched")
	}
}

// ResultMatch matches conditions and run a function that returns a type
func ResultMatch[T, K any](input T) matchResult[T, K] {
	var zeroK K
	return matchResult[T, K]{input, zeroK, false}
}

// when param is "func(input T) bool" used to match a condition
// a param is "func() K" used to run an action
func (m matchResult[T, K]) When(f func(input T) bool, a func() K) matchResult[T, K] {
	if m.isMatched {
		return m
	}

	if f(m.input) {
		m.output = a()
		m.isMatched = true
	}

	return m
}

// val param is "T" used to exact match the input with the passed condition
// a param is "K" used to run an action
func (m matchResult[T, K]) WhenValue(val T, a func() K) matchResult[T, K] {
	if m.isMatched {
		return m
	}

	if cmp.Equal(m.input, val) {
		m.output = a()
		m.isMatched = true
	}

	return m
}

// Get the result from the pattern
func (m matchResult[T, K]) Result() (K, error) {
	var zeroK K
	if cmp.Equal(m.output, zeroK) {
		return zeroK, errors.New("pattern not matched")
	}
	return m.output, nil
}

// Get the optional result from the pattern or throws if not matched
func (m matchResult[T, K]) MaybeResult() gooptional.Optional[K] {
	var zeroK K
	if cmp.Equal(m.output, zeroK) {
		return gooptional.Error[K](errors.New("pattern not matched"))
	}
	return gooptional.Some[K](m.output)
}

// Get the result from the pattern or the passed default value
func (m matchResult[T, K]) ResultOrDefault(def K) K {
	var zeroK K
	if cmp.Equal(m.output, zeroK) {
		return def
	}
	return m.output
}
