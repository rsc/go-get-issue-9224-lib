package liba

func Foo() string {
	return "foo"
}

func Bar() string {
	return "bar"
}

func FooBar() string {
	return Foo() + Bar()
}
