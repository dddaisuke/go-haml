package main

import (
  "github.com/realistschuckle/gohaml"
  "io/ioutil"
  "net/http"
  "text/template"
  "fmt"
)

type Person struct {
    Name string
}

var welcomeTemplate = makeTemplate()

func main() {
  fmt.Println("Please access to http://localhost:8080/ by browser.")
  http.HandleFunc("/", handleRoot)
  http.ListenAndServe("localhost:8080", nil)
}

func makeTemplate() *template.Template {
  var scope = make(map[string]interface{})
  scope["lang"] = "HAML"
  content, _ := ioutil.ReadFile("sample.haml")
  engine, _ := gohaml.NewEngine(string(content))
  output := engine.Render(scope)

  welcomeTemplate := template.Must(template.New("").Parse(output))
  return welcomeTemplate
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
  p := Person { Name: "@dddaisuke" }

  welcomeTemplate.Execute(w, p)
}
