# InnaticalID SDK

InnaticalID is an authentication platform for Innatical, designed to provide a secure and easy-to-use authentication system for your applications. This SDK enables developers to quickly integrate InnaticalID's authentication features into their Golang projects with minimal effort.

## Installation

To install InnaticalID in your Golang project, use the following command:

```
go get github.com/innatical/id-sdk
```

## Usage

To use InnaticalID SDK in your project, simply import the package and start using the provided functions and methods. Here's a basic example to get started:

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/innatical/id-sdk"
)

func main() {
	// Initialize InnaticalID
	idSDK.New("client_id", "client_secret", "redirect_uri")

	// Listen for authorize flow, return generate oauth2 url
    http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		url := idSDK.CreateURL("identity team", "state")

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})

	// callback, return token
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		token, err := idSDK.GetToken(code)

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(token.AccessToken))
		return
	})

	println("Visit http://localhost:8080/authorize to start the flow")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
```

[//]: # (For more advanced usage, refer to the [official documentation]&#40;https://github.com/innatical/id-sdk/wiki&#41;.)

[//]: # (## Documentation)

[//]: # ()
[//]: # (The official documentation is available on the [GitHub wiki]&#40;https://github.com/innatical/id-sdk/wiki&#41;. Here you'll find comprehensive guides, detailed explanations, and example code to help you integrate InnaticalID into your application.)

## Contributing

Contributions to InnaticalID are welcome! If you'd like to report a bug, request a feature, or contribute code, please follow the [contribution guidelines](https://github.com/innatical/id-sdk/blob/main/CONTRIBUTING.md).

## License

InnaticalID SDK is released under the [MIT License](https://github.com/innatical/id-sdk/blob/main/LICENSE).