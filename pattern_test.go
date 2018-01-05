package patternmatch

import "testing"
import "github.com/stretchr/testify/assert"

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
		result := ResultMatch(table.input).
			When(func(x interface{}) bool { return x.(int) > 100 }, func() interface{} { return "> 100" }).
			When(func(x interface{}) bool { return x.(int) > 50 }, func() interface{} { return "> 50" }).
			When(func(x interface{}) bool { return x.(int) > 10 }, func() interface{} { return "> 10" }).
			ResultOrDefault("").(string)

		if result != table.expected {
			t.Errorf("Expected %s but got %s", table.expected, result)
		}
	}
}

func TestWhenValue(t *testing.T) {
	result := ResultMatch(5).
		WhenValue(5, func() interface{} { return "is int 5" }).
		WhenValue("10", func() interface{} { return "is string 10" }).
		WhenValue(true, func() interface{} { return "is bool" }).
		Result()

	if result != "is int 5" {
		t.Errorf("Expected is int 5 but got %s", result)
	}
}

func TestPanicIfNoResult(t *testing.T) {
	assert.Panics(t, func() { ResultMatch(5).WhenValue(6, func() interface{} { return "is int 6" }).Result() })
}

func TestActionWhenValue(t *testing.T) {
	assert.True(t, Match(5).WhenValue(5, func() {}).WhenValue(6, func() {}).isMatched)
}

func TestAction(t *testing.T) {
	called := false
	f := func() { called = true }
	Match(5).When(func(i interface{}) bool { return true }, f)
	assert.True(t, called)
}

func TestOtherwiseThrow(t *testing.T) {
	assert.Panics(t, func() { Match(5).WhenValue(6, func() {}).OtherwiseThrow() })
}

func TestIsMatched(t *testing.T) {

	b := Match(5).
		When(func(i interface{}) bool { return true }, func() {}).
		When(func(i interface{}) bool { return true }, func() {}).isMatched

	assert.True(t, b)
}
