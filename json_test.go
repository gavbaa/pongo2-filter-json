package pongo2filterjson

import (
	"github.com/flosch/pongo2"
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	TestingT(t)
}

// A wrapprt of pongo2.RenderTemplateString
func getResult(s string, ctx pongo2.Context) string {
	result, _ := pongo2.RenderTemplateString(s, ctx)
	return result
}

type TestSuite1 struct{}

var _ = Suite(&TestSuite1{})

func (s *TestSuite1) TestFilters(c *C) {
	c.Assert(getResult("{{ num|json }}", pongo2.Context{"num": 1}), Equals, "1")
	c.Assert(getResult("{{ str|json }}", pongo2.Context{"str": "1"}), Equals, "&quot;1&quot;")
	c.Assert(getResult("{{ map|json }}", pongo2.Context{
		"map": map[string]interface{}{
			"z": []string{"a", "b", "c"},
		},
	}), Equals, "{&quot;z&quot;:[&quot;a&quot;,&quot;b&quot;,&quot;c&quot;]}")
}
