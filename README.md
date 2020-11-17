# missinject
the missinject is a Linter to prevent forgetting to inject dependency.
When initializing the structure, the missinject check whether the structure field of the interface type has a value.

Work as follows.
```go
func f() {
	type A struct {
		string   string
		String   string
		stringer Stringer
		Stringer Stringer
	}

	_ = &A{}

	_ = &A{"", "", nil, nil}

	_ = &A{
		stringer: &iStringer{},
		Stringer: nil,
	}

	_ = &A{ // want "find missing inject: stringer"
		string:   "",
		String:   "",
		Stringer: &iStringer{},
	}

	_ = &A{ // want "find missing inject: stringer" "find missing inject: Stringer"
		string: "",
	}
}
```

## Installation
```
go get -u github.com/theoden9014/missinject/cmd/missinject
```
