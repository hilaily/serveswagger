## ServeSwagger

Render a swagger file in local with http server.

### Install

```bash
go install github.com/hilaily/serveswagger@latest
```

### Usage

```shell
serveswagger
```

It will find swagger.json or swagger.yaml in current directory and serve it 

```shell
serveswagger -p 8080 -f myswagger.json
```

It will render myswagger.json 

