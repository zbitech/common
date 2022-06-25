package object

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"
	//	"github.com/zbi/common/pkg/logger"
)

var (
	NO_FUNCS = template.FuncMap{}
)

type FileTemplate struct {
	Name    string
	Content string
	tmpl    *template.Template
}

func NewFileTemplate(path string, fmap template.FuncMap) (*FileTemplate, error) {

	name := strings.Split(filepath.Base(path), ".")[0]
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	content := string(data)
	return NewTextTemplate(name, content, fmap)
}

func NewTextTemplate(name, content string, fmap template.FuncMap) (*FileTemplate, error) {
	var tmpl *template.Template
	var err error

	if fmap != nil && len(fmap) > 0 {
		tmpl, err = template.New(name).Funcs(fmap).Parse(content)
	} else {
		tmpl, err = template.New(name).Parse(content)
	}
	if err != nil {
		return nil, err
	}

	return &FileTemplate{Name: name, Content: content, tmpl: tmpl}, nil
}

func (f *FileTemplate) ExecuteTemplate(name string, data interface{}) (string, error) {

	var buffer = new(bytes.Buffer)
	var err error

	err = f.tmpl.ExecuteTemplate(buffer, name, data)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func (f *FileTemplate) ExecuteTemplates(names []string, data interface{}) ([]string, error) {

	results := make([]string, len(names))

	for index, name := range names {
		var buffer = new(bytes.Buffer)
		var err error

		err = f.tmpl.ExecuteTemplate(buffer, name, data)
		if err != nil {
			return nil, fmt.Errorf("failed to generate template for %s - %s", name, err)
		}
		results[index] = buffer.String()
	}

	return results, nil
}

func (f *FileTemplate) Execute(data interface{}, fmap template.FuncMap) (string, error) {

	var buffer = new(bytes.Buffer)
	var err error

	if len(fmap) > 0 {
		err = template.Must(f.tmpl.Clone()).Funcs(fmap).Execute(buffer, data)
	} else {
		err = f.tmpl.Execute(buffer, data)
	}

	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
