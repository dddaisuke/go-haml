package main

import (
  "github.com/realistschuckle/gohaml"
  "fmt"
  "io/ioutil"
)

func main() {
  var scope = make(map[string]interface{})
  scope["lang"] = "HAML"
  content, _ := ioutil.ReadFile("sample.haml")
  engine, _ := gohaml.NewEngine(string(content))
  output := engine.Render(scope)
  fmt.Println(output)
}
