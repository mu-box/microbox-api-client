## Microbox V2 API Client

This is the Microbox V2 API client written in Go (golang)

### Installation

Run `go get github.com/mu-box/microbox-api-client` to download the package, then add the client to your Go file's imports list:

##### standard import:

```go
import (
    "github.com/mu-box/microbox-api-client"
)
```

This will give you access to the API client in the form of `microbox-api-client`.

##### aliased import:

```go
import (
    microAPI "github.com/mu-box/microbox-api-client"
)
```

This will give you access to the API client in the nicer `microAPI` form.

### Getting Started

All of the following examples will assume an **aliased import** method:
```go
// create a new client
apiClient = microAPI.NewClient()
```

A user auth token, found on your [dashboard](https://dashboard.microbox.cloud/users/me/auth-token), is required to communicate with the Microbox CLI:

```go
// set api auth token
apiClient.AuthToken = "abc123"
```

Alternately, you can use the GetAuthToken() method to retrieve it using a username and password:
```go
// get a user auth token using a username and password
user, err := client.GetAuthToken(username, password)
// set api auth token
apiClient.AuthToken = user.AuthenticationToken
```

### Usage

The API client uses verbose method names in an attempt to make it clear and consistent to use.

To get a list of all your apps:
```go
// get an index of all apps
apps, err := client.GetApps()
if err != nil { //handle error }
```

To list all of an apps environment variables:

```go
// apps can be found by either name or id
appSlug := "app-name"
// get an index of all an app's evars
evars, err := client.GetAppEVars(appSlug)
if err != nil { //handle error }
```

The API will return both critical (internal) errors from Go, or special API errors. To handle these special errors:

```go
if apiError, ok := err.(microAPI.Error); ok {
    // handle error
}
```

These API errors have a variety of fields available to use when determining the type of error and how to handle/display it.

* `.Error()` - The entire error (ex. {"error":"Not Found"})
* `Code` - The HTTP response status code (ex. 404)
* `Status` - The HTTP response status (ex. "404 Not Found")
* `StatusCode` - A parsed code from the Status (ex. "404")
* `StatusText` - A parsed text from the Status (ex. "Not Found")
* `Body` - The 'body' of a custom API error (ex. "Upgrade Required")

### Optional Parameters

The Microbox API has many actions with optional parameters. For example, when creating an app, you can either specify a name, or allow the API to select one for you.

Options are provided as a pointer to a struct with all available options as fields. For example when creating an app the `AppCreateOptions` struct is passed into the
`CreateApp()` method. Any empty fields on an options struct will be disregarded.

##### options:

```go
// create options field values
name := "my-app"

// create an options struct
options := &microAPI.AppCreateOptions{}

// set option field as a pointer to name
options.Name := &name

// create a new app
app, err := client.CreateApp(options)
if err != nil {
    //handle error
}

fmt.Println("Created ", app.Name)
// Created my-app
```

If you do not wish to set any options, simply pass `nil` into the method.

##### no options:

```go
// pass 'nil' when no additional options are desired
app, err := client.CreateApp(nil)
if err != nil {
    //handle error
}

fmt.Println("Created ", app.Name)
// Created happy-hippo
```

### Documentation

Complete documentation is available on [godoc](http://godoc.org/github.com/mu-box/microbox-api-client).

### Contact

For help using the API client or if you have any questions or suggestions, please find us on [Discord](https://discord.gg/MCDdHfy)
