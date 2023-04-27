# go-acrolinx

An Acrolinx API client library for writing Go programs to interact
with the Acrolinx Platform in an idiomatic way.

## Note

Please note that this library is currently in an experimental state
and not ready for production use.

## Coverage

So far, this client library can do nothing much besides authenticate a
user, get checking capabilities, submit a check request and get
results back.

## Usage

To get started you need to import this library

```go
import "github.com/acrolinx/go-acrolinx"
```

First, you need to create a client object, using an Acrolinx provided
signature and the URL where your Acrolinx Platform is deployed

```go
client, err := acrolinx.NewClient("some-signature", "https://acrolinx.example.com")
if err != nil {
    log.Fatalf("Error creating Acrolinx client: %v", err)
}
```

Next, you need to authenticate a user:

```go
err = client.SignIn("username", "password")
if err != nil {
    log.Fatalf("Error signing in: %v", err)
}
```

Authentication can also be done using an API token created through the
Acrolinx UI by passing an option function when creating the client:

```go
client, err := acrolinx.NewClient("some-signature",
    "https://acrolinx.example.com",
    acrolinx.WithAPIToken("some-api-token")) 
if err != nil {
    log.Fatalf("Error creating Acrolinx client: %v", err)
}
```

Now you're good to go! Get the checking capabilities of your Acrolinx
Platform and use it to check a text. Typically, API methods have
options parameters to further configure the behaviour of the API.

```go
caps, _, err := client.Checking.GetCapabilities(&acrolinx.GetCapabilitiesOptions{})
if err != nil {
    log.Fatalf("Error getting capabilities: %v", err)
}

check, _, err := client.Checking.SubmitCheck(&acrolinx.SubmitCheckOptions{
    Content: "This is a text",
    CheckOptions: &acrolinx.CheckOptions{
        GuidanceProfileID: caps.DefaultGuidanceProfileID,
        ContentFormat:     "text",
        CheckType:         "interactive",
    },
})
```

## Full Example

```go
package main

import (
    "log"
    "time"

    "github.com/acrolinx/go-acrolinx"
)

func main() {
    client, err := acrolinx.NewClient("some-signature", "https://acrolinx.example.com"))
    if err != nil {
        log.Fatalf("Error creating Acrolinx client: %v", err)
    }

    err = client.SignIn("username", "password")
    if err != nil {
        log.Fatalf("Error signing in: %v", err)
    }

    caps, _, err := client.Checking.GetCapabilities(&acrolinx.GetCapabilitiesOptions{})
    if err != nil {
        log.Fatalf("Error getting capabilities: %v", err)
    }

    check, _, err := client.Checking.SubmitCheck(&acrolinx.SubmitCheckOptions{
        Content: "This is a text",
        CheckOptions: &acrolinx.CheckOptions{
            GuidanceProfileID: caps.DefaultGuidanceProfileID,
            ContentFormat:     "text",
            CheckType:         "interactive",
        },
    })

    if err != nil {
        log.Fatalf("Error submitting check request: %v", err)
    }

    for {
        result, _, _ := client.Checking.GetCheckResult(check)
        if result.Progress != nil {
            log.Printf("Check %s is still in progress: %v", check.ID, result.Progress)
            time.Sleep(time.Second)
            continue
        }

        log.Printf("Received check result: %#v", result.Quality)
        break
    }
}
```
