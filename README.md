![go-graphql-api](https://socialify.git.ci/hephhay/go-graphql-api/image?description=1&language=1&name=1&pattern=Circuit%20Board&theme=Light)

## Features

- [x] DI - [Google Wire](https://github.com/google/wire)
- [x] Logging - [Logrus](https://github.com/Sirupsen/logrus)
- [x] Linter - [Go Linter](https://github.com/golangci/golangci-lint)
- [x] Debugger [Delve](https://github.com/go-delve/delve)
- [x] Docker Support
- [x] Pre commit
- [ ] Dataloader
- [ ] Caching
- [ ] Testing

## Pre-commit

To install pre-commit simply run inside the shell:

```bash
pre-commit install
```

pre-commit is very useful to check your code before publishing it. It's configured using .pre-commit-config.yaml file.

### Run

Run go lint

```bash
make go-lint
```

Start server

```bash
make sv-start
```

## Common golang errors:

- outside available modules: remove go.sum and rerun `go get -d -v ./...`

## Screenshots

![demo](demo.png)
