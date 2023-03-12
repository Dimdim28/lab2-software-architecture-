package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
)

func (s *MySuite) TestComputeHandler(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("- 6 2"),
		Output: b,
	}
	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(b.String(), Equals, "6 - 2")
}

func (s *MySuite) TestComputeHandlerHard(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("/ 5 * + 9 8 3"),
		Output: b,
	}
	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(b.String(), Equals, "5 / ((9 + 8) * 3)")
}

func (s *MySuite) TestComputeHandlerError(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("14 88"),
		Output: b,
	}
	err := handler.Compute()

	c.Assert(err, NotNil)
}
