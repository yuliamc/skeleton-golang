# Golang Skeleton

Golang application skeleton by Modal Rakyat Indonesia.

## Table of Contents

- [Requirements](#requirements)
- [Structure](#structure)
- [Development Tools](#development-tools)
- [Known Issues](#known-issues)

## Requirements

- Golang: 1.19.2
- [Gin](https://github.com/gin-gonic/gin)
- [Godotenv](https://github.com/joho/godotenv)
- [Gorm](https://github.com/go-gorm/gorm)
- [Validator](https://github.com/go-playground/validator)

## Structure

#### app

Application main process layer, consists of

- Component initialization
- Configuration reader
- Locale
- Routes
- Handler
- Service locator

#### config

Configuration read from environment variable

#### internal

Internal application functionality layer

- Service
- Resources
- Internal utility
- Middleware

#### pkg

Independent utility function

- Error
- Message
- Custom utility function

## Development Tools

VSCode extension

[Rich Go language support for Visual Studio Code
](https://marketplace.visualstudio.com/items?itemName=golang.go)

### VSCode format/lint setting

```json
  "go.formatTool": "goimports",
  "go.lintOnSave": "file",
  "go.lintTool": "golangci-lint",
  "go.lintFlags": [
    "--fast",
    "--exclude-use-default=false",
    "--print-issued-lines=false",
    "-Egolint",
    "-Egoimports",
    "--exclude-use-default=false"
  ],
```

### Hot Deploy

[Golang Air live reload / hot deploy](https://github.com/cosmtrek/air)

### Thunder Client

VS Code Extension for local postman, import and export if needed.

## LICENSE

[MIT-licensed](https://github.com/yuliamc/skeleton-golang/blob/master/LICENSE)
