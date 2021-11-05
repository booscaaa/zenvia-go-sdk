<p align="center">
  <h1 align="center">Zenvia API - Golang SDK</h1>
  <p align="center">Zenvia API (2.0) </p>
  <p align="center">
    <a href="https://pkg.go.dev/github.com/booscaaa/zenvia-go-sdk"><img alt="Reference" src="https://img.shields.io/badge/go-reference-purple?style=for-the-badge"></a>
    <a href="https://github.com/booscaaa/zenvia-go-sdk/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/booscaaa/zenvia-go-sdk.svg?style=for-the-badge"></a>
    <a href="/LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-red.svg?style=for-the-badge"></a>
    <a href="https://github.com/booscaaa/zenvia-go-sdk/actions/workflows/test.yaml"><img alt="Test status" src="https://img.shields.io/github/workflow/status/booscaaa/zenvia-go-sdk/Test?label=TESTS&style=for-the-badge"></a>
    <a href="https://codecov.io/gh/booscaaa/zenvia-go-sdk"><img alt="Coverage" src="https://img.shields.io/codecov/c/github/booscaaa/zenvia-go-sdk/master.svg?style=for-the-badge"></a>
  </p>
</p>

<br>

## Why?

This project is part of my personal portfolio, so, I'll be happy if you could provide me any feedback about the project, code, structure or anything that you can report that could make me a better developer!

Email-me: boscardinvinicius@gmail.com

Connect with me at [LinkedIn](https://www.linkedin.com/in/booscaaa/).

<br>

## Functionalities

- Send single SMS

<br>

## Getting Started

### Prerequisites

To run this project in the development mode, you'll need to have a basic environment to run:

- A Golang SDK, that can be found [here](https://golang.org/).

<br>

### Installing

```bash
$ go get github.com/booscaaa/zenvia-go-sdk
```

<br>

### Create a config.json file inside your project like this

**The access api_token can be found into Zenvia account**

```json
{
  "zenvia": {
    "api_token": "api token"
  }
}
```

<br>
<br>

### Create main.go file

```go
package main

import (
    "fmt"

    "github.com/booscaaa/zenvia-go-sdk/zenvia"
    "github.com/spf13/viper"
)

func init() {
    viper.SetConfigFile(`config.json`)
    err := viper.ReadInConfig()
    if err != nil {
        panic(err)
    }
}

func main() {
    // Get configuration variavables
    authToken := viper.GetString("zenvia.auth_token")

    // Configure Sandbox access into Zenvia api

    zenviaAccess := zenvia.ZenviaConfig().
      ApiToken(authToken).
      Sandbox()

    zenviaSdk := zenvia.Instance(zenviaAccess)

    // From and To with this format: PREFIX DDD NUMBER WITHOUT MASK
    zenviaSdk.SendSingleSMS("5554999999999", "5554999999999", "Message")

    //see more examples into examples folder
}
```

# Running local

```bash
go run main.go
```

<br>
<br>
<br>

## Api application built With

- [Viper](https://github.com/spf13/viper)

<br>
<br>
<br>

## Contributing

You can send how many PR's do you want, I'll be glad to analyze and accept them! And if you have any question about the project...

Email-me: boscardinvinicius@gmail.com

Connect with me at [LinkedIn](https://www.linkedin.com/in/booscaaa/)

Thank you!

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/booscaaa/zenvia-go-sdk/blob/master/LICENSE) file for details
