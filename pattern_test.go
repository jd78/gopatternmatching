package patternmatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConditionWithAction(t *testing.T) {

	tables := []struct {
		input    int
		expected string
	}{
		{150, "> 100"},
		{70, "> 50"},
		{20, "> 10"},
		{2, ""},
	}

	for _, table := range tables {
		result := ResultMatch[int, string](table.input).
			When(func(x int) bool { return x > 100 }, func() string { return "> 100" }).
			When(func(x int) bool { return x > 50 }, func() string { return "> 50" }).
			When(func(x int) bool { return x > 10 }, func() string { return "> 10" }).
			ResultOrDefault("")

		assert.Equal(t, table.expected, result)
	}
}

func TestWhenValue(t *testing.T) {
	result, err := ResultMatch[interface{}, string](5).
		WhenValue(5, func() string { return "is int 5" }).
		WhenValue("10", func() string { return "is string 10" }).
		WhenValue(true, func() string { return "is bool" }).
		Result()

	assert.NoError(t, err)
	assert.Equal(t, "is int 5", result)
}

func TestResultOrDefault(t *testing.T) {
	result := ResultMatch[interface{}, string](6).
		WhenValue(5, func() string { return "is int 5" }).
		WhenValue("10", func() string { return "is string 10" }).
		WhenValue(true, func() string { return "is bool" }).
		ResultOrDefault("is default")

	assert.Equal(t, "is default", result)
}

func TestErrorIfNoResult(t *testing.T) {
	res, err := ResultMatch[int, string](5).
		WhenValue(6, func() string { return "is int 6" }).
		Result()
	assert.Equal(t, "", res)
	assert.Error(t, err)
}

func TestActionWhenValue(t *testing.T) {
	assert.True(t, Match(5).WhenValue(5, func() {}).WhenValue(6, func() {}).isMatched)
}

func TestAction(t *testing.T) {
	called := false
	f := func() { called = true }
	Match(5).When(func(i int) bool { return true }, f)
	assert.True(t, called)
}

func TestOtherwisePanic(t *testing.T) {
	assert.Panics(t, func() { Match(5).WhenValue(6, func() {}).OtherwisePanic() })
}

func TestIsMatched(t *testing.T) {

	b := Match(5).
		When(func(i int) bool { return true }, func() {}).
		When(func(i int) bool { return true }, func() {}).isMatched

	assert.True(t, b)
}

func TestMaybe(t *testing.T) {
	maybeOpt := ResultMatch[int, string](0).
		WhenValue(5, func() string { return "5" }).
		WhenValue(0, func() string { return "is 0" }).
		MaybeResult()

	assert.True(t, maybeOpt.IsSome())
	assert.Equal(t, "is 0", maybeOpt.Get())
}
