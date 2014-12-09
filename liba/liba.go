package liba

import "github.com/rsc/go-get-issue-9224-lib/libb"

func Foo() string {
	return "foo"
}

func FooBar() string {
	return Foo() + libb.Bar()
}
