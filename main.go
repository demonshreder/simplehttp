package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":2000", nil)
}

var workDir, _ = os.Getwd()
var templatePath = filepath.Join(workDir, "templates/")

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	// t := template.Must(template.ParseFiles(templatePath+"/app/base.html", templatePath+"/app/home.html"))
	t := template.Must(template.ParseFiles(templatePath + "/home.html"))
	t.Execute(w, nil)
	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Category: %v\n", vars["id"])
}
