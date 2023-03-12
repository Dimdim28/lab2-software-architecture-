package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPrefixToInfix(c *C) {
	result, err := PrefixToInfix("+ 2 2")
	c.Assert(result, Equals, "2 + 2")

	result, err = PrefixToInfix("/ 5 * + 9 8 3")
	c.Assert(result, Equals, "5 / ((9 + 8) * 3)")

	result, err = PrefixToInfix("+ 1 * 3 / - 4 / 6 2 - 9 7")
	c.Assert(result, Equals, "1 + 3 * ((4 - 6 / 2) / (9 - 7))")

	result, err = PrefixToInfix(" +++ 12 +++")
	c.Assert(err, ErrorMatches, "invalid input expression")
}

func ExamplePrefixToInfix() {
	res, err := PrefixToInfix("+ 2 2")
	if err == nil {
		fmt.Println(res)
	} else {
		panic(err)
	}

	// Output:
	// 2 + 2
}
