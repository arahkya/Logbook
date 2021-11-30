package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {

	resourceFiles := []string{
		"resources/pages/IndexPage.tmpl",
		"resources/pages/AuthPartialPage.tmpl",
	}

	templateFiles := template.Must(template.ParseFiles(resourceFiles...))

	jsonText := `{
		"UserContext": {
			"FirstName": "Arahk",
			"LastName": "Yambupah"
		},
		"GreetingText": "Hi"
	}`

	var jsonContext map[string]interface{}
	json.Unmarshal([]byte(jsonText), &jsonContext)

	templateFiles.Execute(w, &jsonContext)
}
