package template

import (
	"strings"
	"sync"
	"text/template"

	"github.com/devgek/webskeleton/config"
	"github.com/gobuffalo/packr/v2"
)

//TemplateRoot rootdir for template files
// var TemplateRoot = "./web/templates/"

//TStore ...
type TStore interface {
	GetTemplate(name string) (*template.Template, error)
}

// BoxBasedTemplateStore ...
type BoxBasedTemplateStore struct {
	sync.Mutex
	Box       *packr.Box
	templates map[string]*template.Template
}

//NewBoxBasedTemplateStore ...
func NewBoxBasedTemplateStore(box *packr.Box) TStore {
	return &BoxBasedTemplateStore{Box: box, templates: map[string]*template.Template{}}
}

//GetTemplate ...
func (ts *BoxBasedTemplateStore) GetTemplate(fileName string) (*template.Template, error) {
	ts.Lock()
	defer ts.Unlock()

	//if dev mode, then parse the template on each request
	if val, ok := ts.templates[fileName]; ok && !config.IsDev() {
		return val, nil
	}

	var templ *template.Template
	var err error

	switch {
	case fileName == "login":
		templ, err = parsePacked(ts.Box, fileName+".html")
	case strings.Contains(fileName, "page"):
		templ, err = parsePacked(ts.Box, "layout.html", fileName+".html")
	default:
		templ, err = parsePacked(ts.Box, "layout.html", fileName+".html", fileName+"-edit.html", "confirm-delete.html")
	}

	if err == nil {
		ts.templates[fileName] = templ
	}

	return templ, err
}

// ParsePacked parses html templates from packr box
// template is nil, it is created from the first file.
func parsePacked(box *packr.Box, filenames ...string) (*template.Template, error) {
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
