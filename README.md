# ciinfo

ciinfo is a library for detecting Continuous Integration environments and Pull Request contexts.

It is inspired by [ci-info](https://github.com/watson/ci-info). And reuses the same vendor dataset and detection semantics, reimplement with Go.

```sh
go get github.com/startracex/ciinfo
```

```go
info := ciinfo.GetInfo()
if info.IsCI {
    println("CI ID:", info.ID)
    println("CI Provider:", info.Name)
    println("Is PR:", info.IsPR)
}
```

Vendors came from `ci-info`'s `vendors.json`, they are configurable.

Syntax can unmarshal from json.

```go
envs := os.Environ()
var myVendor vendors.Vendor
json.Unmarshal([]byte(`
{
    "name": "My Vendor",
    "constant": "MY_VENDOR",
    "env": "MY_VENDOR",
    "pr": {
        "any": ["MY_PULL_REQUEST_NUMBER", "MY_PULL_REQUEST_ID"]
    }
}
`), &myVendor)
myOtherVendor := vendors.Vendor{
    Name:     "My Other Vendor",
    Constant: "MY_OTHER_VENDOR",
    Env: syntax.EnvList{{
        StrictEqual: "MY_OTHER_VENDOR",
    }},
    PR: &syntax.PR{
        EqualsAnyOf: []string{"MY_OTHER_PULL_REQUEST_NUMBER", "MY_OTHER_PULL_REQUEST_ID"},
    },
}
ciinfo.GetInfoFrom(ciinfo.EnvironMap(os.Environ()), []vendors.Vendor{myVendor, myOtherVendor})
```
