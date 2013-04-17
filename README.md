go-haml
=======

How to use haml with go lang?
---

You can find out more about HAML at [haml-lang.com](http://haml.info/)

What is it?
---

This is sample code for [gohaml](https://github.com/realistschuckle/gohaml)

```go run main.go```

This haml file is changed into following html.

```
#main_id
  title
  .foo
    bar
```

```
<div id="main_id">
        title
        <div class="foo">
                bar
        </div>
</div>
```
