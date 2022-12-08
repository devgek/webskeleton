package template

import (
	"embed"
	"io/fs"
	"strings"
	"sync"
	"text/template"

	"github.com/devgek/webskeleton/config"
)

//go:embed templates/*.html
var templateDir embed.FS
var baseDir = "templates/"

// TStore ...
type TStore interface {
	GetTemplate(name string) (*template.Template, error)
}

// FSBasedTemplateStore ...
type FSBasedTemplateStore struct {
	sync.Mutex
	templateDir embed.FS
	templates   map[string]*template.Template
}

// NewFSBasedTemplateStore ...
func NewFSBasedTemplateStore() TStore {
	return &FSBasedTemplateStore{templateDir: templateDir, templates: map[string]*template.Template{}}
}

// GetTemplate ...
func (ts *FSBasedTemplateStore) GetTemplate(fileName string) (*template.Template, error) {
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
		templ, err = parseEmbedded(ts.templateDir, templateName(fileName))
	case strings.Contains(fileName, "page"):
		templ, err = parseEmbedded(ts.templateDir, templateName("layout"), templateName(fileName))
	default:
		templ, err = parseEmbedded(ts.templateDir, templateName("layout"), templateName(fileName), templateName(fileName+"-edit"), templateName("confirm-delete"))
	}

	if err == nil {
		ts.templates[fileName] = templ
	}

	return templ, err
}

func templateName(template string) string {
	return baseDir + template + ".html"
}

// ParseEmbedded parses html templates from embedded Filesystem
// template is nil, it is created from the first file.
func parseEmbedded(templateDir embed.FS, filenames ...string) (*template.Template, error) {
	var t *template.Template

	for _, filename := range filenames {
		b, err := fs.ReadFile(templateDir, filename)
		if err != nil {
			return nil, err
		}
		s := string(b)
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
