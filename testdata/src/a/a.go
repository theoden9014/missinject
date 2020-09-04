package a

type itr interface {
	impl()
}

type impl struct{}

func (i *impl) impl() {}

type A struct {
	string string
	String string
	itr    itr
	Itr    itr
}

func main() {
	_ = &A{}

	_ = &A{
		string: "",
		String: "",
		itr:    &impl{},
		Itr:    &impl{},
	}

	_ = &A{ // want "find missing inject: itr"
		string: "",
		String: "",
		Itr:    &impl{},
	}

	_ = &A{ // want "find missing inject: itr" "find missing inject: Itr"
		string: "",
	}
}
