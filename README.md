# poc Go template

go template implement with hexeganal architecture concept.

- [x] rest api example.
- [x] swagger documents.
- [x] hot reload.
- [x] rest-api validate input, and mapping error message.
- [ ] error catcher, error handler middleware.
- [ ] unit test example, using mockery to generate dependency.
- [ ] logger.
- [x] example connect mysql.
- [ ] database migrate.
- [ ] database seed.

## project structure

```
.
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   └── app
│       └── main.go
├── config.json
├── configs
│   ├── config.go
│   ├── db.go
│   ├── public.go
│   └── server.go
├── controllers
│   ├── example
│   │   ├── example.go
│   │   └── interface.go
│   └── system.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── pkg
│   ├── database
│   │   ├── database.go
│   │   ├── mongo
│   │   ├── mysql
│   │   └── postgres
│   ├── model
│   │   └── model.go
│   ├── redis
│   └── restapi
│       ├── context
│       ├── echo
│       ├── fasthttp
│       ├── gin
│       ├── response
│       └── restapi.go
├── repositories
│   └── example
│       ├── example.go
│       ├── interface.go
│       └── models
├── routes
│   └── routes.go
├── services
│   ├── entities
│   │   ├── request
│   │   └── response
│   └── example
│       ├── example.go
│       └── interface.go
```
