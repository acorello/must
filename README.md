# Generic functions to reduce error handling boilerplate

```go
// must.Get example
var alwaysWorks = must.Get(canNotFail("with static input")) // panic if there is an error, or discard it

// must.Have example
var alwaysSet = must.Have(staticMap["key"]) // panics if (value, found) is not found 
```