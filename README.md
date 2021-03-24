# Hopex Golang SDK

This is Hopex Go SDK,  you can install to your Golang project and use this SDK to query all market data and trading.

The SDK supports RESTful API invoking.

## Quick Start

The SDK is compiled by Go 1.16, you can import this SDK in your Golang project:

* go get -u github.com/hopex-hk/go_sdk
* Create one of the clients (under package **client**) instance by **Init** or **InitByDefault** method
* Call the method provided by client.

```go
package main

import (
	"fmt"
	"github.com/hopex-hk/go_sdk/client"
	"github.com/hopex-hk/go_sdk/core"
)

func main() {
	cfg := core.NewConfig(
		"https://api1.hopex.com",
		"your app key",
		"your app secret",
		"{your app name}/{app version}",
	)

	// Get the user info from Hopex server and print on console
	client := new(client.AccountClient).InitByDefault(cfg)
	res, err := client.GetUserInfo()

	if err != nil {
		fmt.Printf("has error: %+v", err)
	} else {
		fmt.Printf("get userinfo: %+v", res)
	}
}

    
```

## Usage

After above section, this SDK should be already download to your local machine, this section introduce this SDK and how to use it correctly.

### Folder Structure

This is the folder and package structure of SDK source code and the description

- **client**: The client struct that are responsible to access data
- **model**: The request and response data model
- **core** The core package defines Logger, HttpRequester, Config, etc.
  - **logging**: It define **Loggger** interface and provide a default implementation
  - **model**: The common data model
  - **utils**: Provide common utils
- **example**: The examples how to use **core** and **client** to access API and read response.

### Run examples

This SDK provides examples that under **example** folder, if you want to run the examples to access private data, you need below additional steps:

1. Create an **API Key** first from [Hopex official website](https://web.hopex.com)
2. Create **config.go** into your **example/config** folder (package). The purpose of this file is to prevent submitting SecretKey into repository by accident, so this file is already added in the *.gitignore* file. 
3. **cd example; go run ./main.go** ([view source code](https://github.com/hopex-hk/go_sdk/tree/main/example))
```go
// config.go file
package config

import "github.com/hopex-hk/go_sdk/core"

var DemoConfig *core.Config

func init() {
	DemoConfig = core.NewConfig(
		"https://api1.hopex.com",
		"your app key",
		"your app secret",
		"{your app name}/{app version}",
	)
}

```

If you don't need to access private data, you can ignore the secret key.

Regarding the difference between public data and private data you can find details in [Client](#Client) section below.



### Client

In this SDK, the client is the object to access the Hopex API. In order to isolate the private data with public data, and isolated different kind of data, the client category is designated to match the API category.

All the client is listed in below table. Each client is very small and simple, it is only responsible to operate its related data, you can pick up multiple clients to create your own application based on your business.

| Data Category | Client        | Privacy | Access Type |
| ------------- | ------------- | ------- | ----------- |
| Account       | AccountClient | Private | Rest        |
| Home          | HomeClient    | Public  | Rest        |
| Market        | MarketClient  | Public  | Rest        |
| Trade         | TradeClient   | Private | Rest        |

### Logging

The SDK support custom log library (by implementing the **Logger** interface under the **core/logging** package), and it also provides a default log implementation.
The default log implementation uses the high performance logging library [zap](https://github.com/uber-go/zap), which provide different kind of loggers. To better support format message, this SDK uses the SugaredLogger. It has below features:

1. Logging target is console (In the future we will support output to file)
2. Support multiple levels (Fatal, Error, Panic, Warn, Info and Debug) and minimum log level
3. Support colorful text (by default)