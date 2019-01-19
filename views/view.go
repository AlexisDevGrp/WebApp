package views

import "html/template"

// NewView manages all used views
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

// View struc definition
type View struct {
	Template *template.Template
}
