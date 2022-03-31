package object

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"
	//	"github.com/zbi/common/pkg/logger"
)

type FileTemplate struct {
	Name    string
	Content string
	tmpl    *template.Template
}

func NewFileTemplate(path string) (*FileTemplate, error) {

	name := strings.Split(filepath.Base(path), ".")[0]
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	content := string(data)
	return NewTextTemplate(name, content)
}

func NewTextTemplate(name, content string) (*FileTemplate, error) {
	tmpl, err := template.New(name).Parse(content)
	if err != nil {
		return nil, err
	}

	return &FileTemplate{Name: name, Content: content, tmpl: tmpl}, nil
}

func (f *FileTemplate) ExecuteTemplate(name string, data interface{}) (string, error) {

	buffer := new(bytes.Buffer)
	err := f.tmpl.ExecuteTemplate(buffer, name, data)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func (f *FileTemplate) ExecuteTemplates(names []string, data interface{}) ([]string, error) {

	results := make([]string, len(names))

	for index, name := range names {
		buffer := new(bytes.Buffer)
		err := f.tmpl.ExecuteTemplate(buffer, name, data)
		if err != nil {
			return nil, err
		}
		//		logger.Infof(context.Background(), "Adding %s to content", buffer.String())
		results[index] = buffer.String()
	}

	return results, nil
}

func (f *FileTemplate) Execute(data interface{}) (string, error) {

	buffer := new(bytes.Buffer)
	if err := f.tmpl.Execute(buffer, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
