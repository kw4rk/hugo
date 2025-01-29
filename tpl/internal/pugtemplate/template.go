package pugtemplate

import (
	"bytes"
	"io"
	"github.com/Joker/jade"
	"github.com/gohugoio/hugo/tpl"
)

type PugTemplate struct {
	name string
	tpl  *jade.Jade
}

func (p *PugTemplate) Name() string {
	return p.name
}

func (p *PugTemplate) Prepare() (*tpl.Template, error) {
	return nil, nil
}

func (p *PugTemplate) Parse(name, tpl string) (tpl.Template, error) {
	jadeTpl, err := jade.Parse(name, tpl)
	if err != nil {
		return nil, err
	}
	return &PugTemplate{name: name, tpl: jadeTpl}, nil
}

func (p *PugTemplate) Execute(t tpl.Template, wr io.Writer, data any) error {
	var buf bytes.Buffer
	err := p.tpl.Execute(&buf, data)
	if err != nil {
		return err
	}
	_, err = wr.Write(buf.Bytes())
	return err
}
