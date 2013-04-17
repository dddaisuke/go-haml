package main

import (
  "github.com/realistschuckle/gohaml"
  "io/ioutil"
  "os"
  "text/template"
)

type Person struct {
    Name string
}

func main() {
  var scope = make(map[string]interface{})
  scope["lang"] = "HAML"
  content, _ := ioutil.ReadFile("sample.haml")
  engine, _ := gohaml.NewEngine(string(content))
  output := engine.Render(scope)

  t := template.New(output)
  t, _ = t.Parse(output)

  p := Person { Name: "@dddaisuke" }
  t.Execute(os.Stdout, p)
}
