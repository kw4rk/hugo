package template_test

import (
	"bytes"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/pugtemplate"
	qt "github.com/frankban/quicktest"
)

func TestPugTemplateEngine(t *testing.T) {
	c := qt.New(t)

	engine := &pugtemplate.PugTemplateEngine{}

	// Test parsing a Pug template
	tpl, err := engine.Parse("test", "p Hello, World!")
	c.Assert(err, qt.IsNil)
	c.Assert(tpl.Name(), qt.Equals, "test")

	// Test executing the Pug template
	var buf bytes.Buffer
	err = engine.Execute(tpl, &buf, nil)
	c.Assert(err, qt.IsNil)
	c.Assert(buf.String(), qt.Equals, "<p>Hello, World!</p>")
}
