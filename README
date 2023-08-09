# Error Handling Example
Inspired by [this article from Dave Cheney](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully). This is an example of checking for error behavior rather than checking for specific values.

The example contains two example error behaviors we can check for:
- ***Not Found*** representing that something wasn't found where it was expected to be
- ***InvalidRequest*** representing that some area of the code received an invalid request

Common errors behaviors can be defined in the error package and re-used throughout the code while still attaching any domain specific context for logging. If a domain has a specific behavior it wants to check, those behaviors can be defined with the domain itself

Running main.go yields the following output
```
"basic" is not found error: true
not found: couldn't find foo in place

"thinlyWrapped" is not found error: true
not found: some database error saying not found

"deeplyWrapped" is not found error: true
not found: outermost context: middle context: inner context: some database error saying not found

"otherKindOfError" is not found error: false
some other error
```

