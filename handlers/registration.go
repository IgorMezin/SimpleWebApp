package registration

import (
	template "html/template"
	http "net/http"
)

type ContactDetails struct {
	Login         string
	Password      string
	Success       bool
	StorageAccess string
}

var (
	tmpl = template.Must(template.ParseFiles("forms.html"))
)

func Handler(w http.ResponseWriter, r *http.Request) {
	data := ContactDetails{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}
	tmpl.Execute(w, data)
}
