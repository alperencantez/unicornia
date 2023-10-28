<div style="text-align:center">
<h1> 🦄 Unicornia</h1>
A list of utility functions to make your Go development easier. Inspired by Lodash, written in Go.

**This is the version 1.0 & feel free to contribute.**

</div>

## Nice to know

- Godoc is available for all the methods. You can get the detailed information about a specific function just by hovering on it.
- While creating Unicornia, I tried to avoid extra 3rd party packages because I wanted to keep it simple and steady.

### Strings

```go
uuid := unicornia.StringUuid() // FD98536E-63B6-404F-A1CD-45BDBE9E2864
toSnakeCase := unicornia.StringToSnakeCase("unicornia is great") // unicornia_is_great
truncate := unicornia.StringTruncate("never gonna give you up", 8) // never go...
repeat := unicornia.StringRepeat("go", 4) // gogogogo
startsWith := unicornia.StringStartsWith("unicornia", "un") // true
endsWith := unicornia.StringEndsWith("unicornia", "xo") // false
isAlphaNum := unicornia.StringIsAlphaNum("42") // true
slug := unicornia.StringSlugify("unicornia is a utility library FOR Go")
// unicornia-is-a-utility-library-for-go
```

### Arrays

```go
shuffled := unicornia.ArrayShuffle([]any{1, 2, 3, "unicornia is great"})
// [3 1 unicornia is great 2]
includes := unicornia.ArrayIncludes([]int{1, 2, 3}, 5) // false
findIndex := unicornia.ArrayFindIndex([]int{1, 2, 3}, 2) // 1
reversedArray := unicornia.ArrayReverse([]int{1, 2, 3, 4, 5}) // [5 4 3 2 1]
sortedArray := unicornia.ArrayIntSort([]int{1, 2, 3, 4}, "desc") // [4 3 2 1]
intersection := unicornia.ArrayIntersection([]int{1, 2, 3, 4}, []int{1, 6, 3,7}) // [1 3]
```

### Math

```go
inRange := unicornia.MathInRange(42, 1, 83) // true
mean := unicornia.MathFindMean(30, 40, 56) // 42
```

### Utils

```go
isCreditCard := unicornia.UtilsIsCreditCard("4225319580287") // true visa
```

## LICENSE

This project is licensed under the <a href="https://github.com/alperencantez/unicornia/blob/v1.0/LICENSE">MIT License.</a>

<p style="text-align:center; color: gray">initial commit on 2023-10-28</p>
