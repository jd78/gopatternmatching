# Go Pattern Matching

Pattern Matching for Golang inspired by Hudo's ["Just another pattern matching for C#"](https://github.com/hudo/PatternMatch)

## Examples

```go

result, err := patternmatch.ResultMatch[int, string](70).
    When(func(x int) bool { return x > 100 }, func() string { return "> 100" }).
    When(func(x int) bool { return x > 50 }, func() string { return "> 50" }).
    When(func(x int) bool { return x > 10 }, func() string { return "> 10" }).
    Result()
            
// Result will be "> 50", nil

result, err := patternmatch.ResultMatch[int, string](2).
    When(func(x int) bool { return x > 100 }, func() string { return "> 100" }).
    When(func(x int) bool { return x > 50 }, func() string { return "> 50" }).
    When(func(x int) bool { return x > 10 }, func() string { return "> 10" }).
    Result()

// result will be "", error{"pattern not matched"}

maybeResult := patternmatch.ResultMatch[int, string](11).
    When(func(x int) bool { return x > 100 }, func() string { return "> 100" }).
    When(func(x int) bool { return x > 50 }, func() string { return "> 50" }).
    When(func(x int) bool { return x > 10 }, func() string { return "> 10" }).
    MaybeResult()

// result will be Optional[string] -> Some "> 10"

```

Using defaults:

```go

result, err := patternmatch.ResultMatch[int, string](2).
    When(func(x int) bool { return x.(int) > 100 }, func() string { return "> 100" }).
    When(func(x int) bool { return x.(int) > 50 }, func() string { return "> 50" }).
    When(func(x int) bool { return x.(int) > 10 }, func() string { return "> 10" }).
    ResultOrDefault("< 10")
            
// Result will be "< 10", nil

```

Non-conditional pattern matching:

```go
result, err := patternmatch.ResultMatch[interface{}, string](5).
    WhenValue(5, func() string { return "is int 5" }).
    WhenValue("10", func() string { return "is string 10" }).
    WhenValue(true, func() string { return "is bool" }).
    Result()

// Result will be "is int 5", nil
```

Pattern matching that does not return a value but executes an action:

```go

patternmatch.Match(5).
    When(func(i int) bool { return i == 5 }, func() { fmt.Println("is 5") }).
    When(func(i int) bool { return i == 6 }, func() { fmt.Println("is 6") })

patternmatch.Match(5).
    WhenValue(5, func() { fmt.Println("is 5") }).
    WhenValue(6, func() { fmt.Println("is 6") })

patternmatch.Match(10).
    WhenValue(5, func() { fmt.Println("is 5") }).
    WhenValue(6, func() { fmt.Println("is 6") }).
    OtherwisePanic() // will panic if not matched

```
