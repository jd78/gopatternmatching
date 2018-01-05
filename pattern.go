package patternmatch

import "reflect"

type match struct {
	input     interface{}
	isMatched bool
}

type matchResult struct {
	input     interface{}
	output    interface{}
	isMatched bool
}

// Match condition and run an action that does not return a type
func Match(input interface{}) match {
	return match{reflect.ValueOf(input).Interface(), false}
}

// when param is "func(input interface{}) bool" used to match a condition
// a param is "func()" used to run an action
func (m match) When(f func(input interface{}) bool, a func()) match {
	if m.isMatched {
		return m
	}

	if f(m.input) {
		a()
		m.isMatched = true
	}

	return m
}

// val param is "interface{}" used to exact match the input with the passed condition
// a param is "func()" used to run an action
func (m match) WhenValue(val interface{}, a func()) match {
	if m.isMatched {
		return m
	}

	if m.input == val {
		a()
		m.isMatched = true
	}

	return m
}

// OtherwiseThrow throws if the pattern is not matched, optionally used at the end of the pattern matching
func (m match) OtherwiseThrow() {
	if !m.isMatched {
		panic("pattern not matched")
	}
}

// ResultMatch matches conditions and run a function that returns a type
func ResultMatch(input interface{}) matchResult {
	return matchResult{reflect.ValueOf(input).Interface(), nil, false}
}

// when param is "func(input interface{}) bool" used to match a condition
// a param is "func() interface{}" used to run an action
func (m matchResult) When(f func(input interface{}) bool, a func() interface{}) matchResult {
	if m.isMatched {
		return m
	}

	if f(m.input) {
		m.output = a()
		m.isMatched = true
	}

	return m
}

// val param is "interface{}" used to exact match the input with the passed condition
// a param is "func() interface{}" used to run an action
func (m matchResult) WhenValue(val interface{}, a func() interface{}) matchResult {
	if m.isMatched {
		return m
	}

	if m.input == val {
		m.output = a()
		m.isMatched = true
	}

	return m
}

// Get the result from the pattern or throws if not matched
func (m matchResult) Result() interface{} {
	if m.output == nil {
		panic("pattern not matched")
	}
	return m.output
}

// Get the result from the pattern or the passed default value
func (m matchResult) ResultOrDefault(def interface{}) interface{} {
	if m.output == nil {
		return def
	}
	return m.output
}
