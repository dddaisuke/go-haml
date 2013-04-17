go-haml
=======

How to use haml with go lang?
---

You can find out more about HAML at [haml-lang.com](http://haml.info/)

What is it?
---

This is sample code for [gohaml](https://github.com/realistschuckle/gohaml)

```
git checkout v1
go run main.go
```

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

### Using template engine
```
git checkout v3
go run main.go
```

Replace {{.Name}} at sample.haml to @dddaisuke string.

```
#main_id
  welcome to {{.Name}}
  .foo
    bar
```

```
<div id="main_id">
        welcome to @dddaisuke
        <div class="foo">
                bar
        </div>
</div>
```

### Using web server
```
git checkout v4
go run main.go
```
and access to [http://localhost:8080/](http://localhost:8080/)
