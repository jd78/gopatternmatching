# Go Pattern Matching

[![Build Status](https://travis-ci.org/jd78/gopatternmatching.svg?branch=master)](https://travis-ci.org/jd78/gopatternmatching)

Pattern Matching for Golang inspired from Hudo's "Just another pattern matching for c#" (https://github.com/hudo/PatternMatch)

#### Examples  

```go

result := patternmatch.ResultMatch(70).
		When(func(x interface{}) bool { return x.(int) > 100 }, func() interface{} { return "> 100" }).
		When(func(x interface{}) bool { return x.(int) > 50 }, func() interface{} { return "> 50" }).
		When(func(x interface{}) bool { return x.(int) > 10 }, func() interface{} { return "> 10" }).
            	Result().(string)
            
//Result will be "> 50"

result := patternmatch.ResultMatch(2).
		When(func(x interface{}) bool { return x.(int) > 100 }, func() interface{} { return "> 100" }).
		When(func(x interface{}) bool { return x.(int) > 50 }, func() interface{} { return "> 50" }).
		When(func(x interface{}) bool { return x.(int) > 10 }, func() interface{} { return "> 10" }).
            	Result().(string)

//If unmatched Result will panic

```

Using defaults

```go

result := patternmatch.ResultMatch(2).
		When(func(x interface{}) bool { return x.(int) > 100 }, func() interface{} { return "> 100" }).
		When(func(x interface{}) bool { return x.(int) > 50 }, func() interface{} { return "> 50" }).
		When(func(x interface{}) bool { return x.(int) > 10 }, func() interface{} { return "> 10" }).
            	ResultOrDefault("< 10").(string)
            
//Result will be "< 10", with ResultOrDefault you'll never panic

```

No conditional pattern matching

```go
result := patternmatch.ResultMatch(5).
		WhenValue(5, func() interface{} { return "is int 5" }).
		WhenValue("10", func() interface{} { return "is string 10" }).
		WhenValue(true, func() interface{} { return "is bool" }).
        	Result()
```

Pattern matching that does not return a value but executes an action

```go

patternmatch.Match(5)
    .When(func(i interface{}) bool { i == 5 }, func() { fmt.Println("is 5") })
    .When(func(i interface{}) bool { i == 6 }, func() { fmt.Println("is 6") })

patternmatch.Match(5)
    .WhenValue(5, func() { fmt.Println("is 5") })
    .WhenValue(5, func() { fmt.Println("is 6") })

patternmatch.Match(10)
    .WhenValue(5, func() { fmt.Println("is 5") })
    .WhenValue(5, func() { fmt.Println("is 6") })
    .OtherwiseThrow() //will throw if not matched

```
