package model

type {{.Model}} struct {
    {{range .Fields}}{{.}}
    {{end}}
}

