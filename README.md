### EBonding Library

This is a Go library would allow to call different Wireless Providers, this is for important information that we use on the CPD application.

### The Client Interface

In order to guarantee that all the clients provide the same sort of operations each of the carrier implementations should satisfy the ebonding.Client interface which is:

```go
package ebonding

type Client interface {
    // FindDevices return all the devices available for that particular client.
    func FindDevices() ([]Device, error)
    
    // FindAccessories returns accessories in that particular client catalog.
    func FindAccessories() ([]Accessory, error)
}
```

### Credentials

As each of these EBonding API calls will require authenticated access, the API provides a struct to pass those when creating the client.

```go
package ebonding

type Credentials struct {
    Username string
    Password string
    ExtraSettings map[string]interface{}
}
```

### Creating a Client

In order to create a client the library provides a `NewClient` function, which can be used like in the following example:

```go
client = ebonding.NewClient(ebonding.ClientTypeATT, ebonding.Credentials{...})
```

As you might have noticed there is a `ebonding.ClientTypeATT` constant being passed to the `NewClient` function. The library should have a number of ClientType constants defined like:

```go
package ebonding

//ClientType constants
const (
    _ = iota
    ClientTypeATT
    ClientTypeVerizon
    ClientTypeSprint
    ClientTypeTmobile
    ...
)
```

That will be used by the CPD app when calling the `NewClient` function.

### Devices

Ebonding library provides a Device struct that represents the devices that each carrier has in their catalog.

```go
type Device struct {
    SKU                 string
    ManufacturerPartID  string

    Color       string 
    OS          string
    Capacity    string
    
    URL string
    CompatibleAccessories []Accessory
}
```

### Device Catalog

Ebonding library allows the CPD platform to perform operations over carriers. This is achieved by creating an ebonding client and calling the appropriate method, in this case: `FindDevices(options FilterOptions)`.

```go
import (
    ...
    "motus/ebonding"
    ...
)

...

// AT&T client
client := ebonding.NewClient(ebonding.ClientATT, ebonding.Credentials{
    Username: "username", 
    Password: "password", 
    ExtraSettings: map[string]string {
        "GroupID": "groupID",
    }
})

devices, err := client.FindDevices()
if err != nil {
    // Do something with the error, this error could be things like:
    // - Network error
    // - Auth error
    // - Quota exceeded
}

for _, device := range devices {
    //Should print SKU and description in the response p.e. sku11920100001 : Apple iPhone XR (256GB Black)
    fmt.Sprintf("%v : %v\n", device.SKU, device.Description) 
}
...
```