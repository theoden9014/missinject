package a

type Stringer interface {
	String() string
}

type iStringer struct{}

func (i *iStringer) String() string { return "" }

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
