package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

const (
	LayoutDir   = "views/layout/"
	TemplateDir = "views/"
	TemplateExt = ".html"
)

type View struct {
	Template *template.Template
	Layout   string
}

// render is used to render the views with the predefined layout
func (v *View) Render(w http.ResponseWriter, data interface{}) {
	var buf bytes.Buffer
	if err := v.Template.ExecuteTemplate(&buf, v.Layout, data); err != nil {
		log.Println("Render() error:", err)
		http.Error(w, "Something went wrong!", http.StatusInternalServerError)
		return
	}
	if _, err := io.Copy(w, &buf); err != nil {
		log.Println("Copy error:", err)
		return
	}
}

func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	fmt.Println(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil)
}

// takes in a slice of strings representing file path for templates
// and it prepends the TemplateDir directory to each string in the slice
//
// Eg the input {"home"} would result in the output
// {"views/home"} if TemplateDir == "views/"
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

// takes in a slice of strings representing file path for templates and it appends
// the TemplateExt extension to each string in slice
//
// Eg the input {"home"} would result in the output {"home.html"} if TemplateExt == ".html"
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
