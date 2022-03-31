package object

import "testing"

var (
	content = `
{{define "section1"}}
apiVersion: v1
kind: ConfigMap
metadata:
  name: cm-{{.Name}}
{{end}}

{{define "section2"}}
apiVersion: v1
kind: Secret
metadata:
  name: secret-{{.Name}}
{{end}}

{{define "section3"}}
apiVersion: v1
kind: Service
metadata:
  name: svc-{{.Name}}
{{end}}
	`

	sections = []string{"section1", "section2", "section3"}
)

func Test_NewFileTemplate(t *testing.T) {

}

func Test_NewTextTemplate(t *testing.T) {
	ft, err := NewTextTemplate("main", content)
	if err != nil {
		t.Errorf("Failed to parse content - %s", err)
	}

	if ft == nil {
		t.Error("Failed to create template")
	}

	if ft.Content != content {
		t.Errorf("want %s but got %s", content, ft.Content)
	}
}

func Test_ExecuteTemplate(t *testing.T) {

	var input struct {
		Name string
	}

	input.Name = "main"

	ft, err := NewTextTemplate("main", content)
	if err != nil {
		t.Errorf("Failed to parse content - %s", err)
	}

	for _, section := range sections {
		t.Run("", func(t *testing.T) {
			content, err := ft.ExecuteTemplate(section, input)
			if err != nil {
				t.Errorf("Failed to generate content - %s", err)
			}

			t.Logf("Section %s = %s", section, content)
		})
	}

}

func Test_ExecuteTemplates(t *testing.T) {

	var input struct {
		Name string
	}

	input.Name = "main"

	ft, err := NewTextTemplate("main", content)
	if err != nil {
		t.Errorf("Failed to parse content - %s", err)
	}

	contents, err := ft.ExecuteTemplates(sections, input)
	if err != nil {
		t.Errorf("Failed to generate content - %s", err)
	}

	for index, content := range contents {
		t.Logf("%d. %s", index+1, content)
	}
}

func Test_Execute(t *testing.T) {

}
