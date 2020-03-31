package template

import (
	"github.com/devgek/webskeleton/global"
	"github.com/devgek/webskeleton/packrfix"
	"sync"
	"text/template"
)

//TemplateRoot rootdir for template files
var TemplateRoot = "./web/templates/"

//TStore ...
type TStore interface {
	GetTemplate(name string) *template.Template
}

// BoxBasedTemplateStore ...
type BoxBasedTemplateStore struct {
	sync.Mutex
	Box       *packrfix.BoxExtended
	templates map[string]*template.Template
}

//NewBoxBasedTemplateStore ...
func NewBoxBasedTemplateStore(box *packrfix.BoxExtended) TStore {
	return &BoxBasedTemplateStore{Box: box, templates: map[string]*template.Template{}}
}

//GetTemplate ...
func (ts *BoxBasedTemplateStore) GetTemplate(fileName string) *template.Template {
	ts.Lock()
	defer ts.Unlock()

	if val, ok := ts.templates[fileName]; ok && !global.Debug {
		return val
	}

	var templ *template.Template

	if fileName == "login.html" {
		templ = template.Must(parsePacked(ts.Box, fileName))
	} else {
		templ = template.Must(parsePacked(ts.Box, "layout.html", fileName, "user-edit.html", "confirm-delete.html"))
	}

	ts.templates[fileName] = templ

	return templ
}

// ParsePacked parses html templates from packr box
// template is nil, it is created from the first file.
func parsePacked(box *packrfix.BoxExtended, filenames ...string) (*template.Template, error) {
	var t *template.Template

	for _, filename := range filenames {
		s, err := box.FindString(filename)
		if err != nil {
			return nil, err
		}
		name := filename
		// First template becomes return value if not already defined,
		// and we use that one for subsequent New calls to associate
		// all the templates together. Also, if this file has the same name
		// as t, this file becomes the contents of t, so
		//  t, err := New(name).Funcs(xxx).ParseFiles(name)
		// works. Otherwise we create a new template associated with t.
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
