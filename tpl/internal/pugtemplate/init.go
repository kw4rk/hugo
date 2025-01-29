package pugtemplate

import "github.com/gohugoio/hugo/tpl"

func init() {
	pugEngine := &PugTemplateEngine{}
	tpl.RegisterTemplateEngine("pug", pugEngine)
}
