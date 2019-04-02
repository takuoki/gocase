# gocase

A golang package to convert to normal CamelCase or Golang's CamelCase.

## Example

* `To`: converts a string from `From` to `To`.
* `Revert`: converts a string from `To` to `From`.

| From | To |
|:-|:-|
|`""`|`""`|
|`"jsonFile"`|`"jsonFile"`|
|`"IpAddress"`|`"IPAddress"`|
|`"defaultDnsServer"`|`"defaultDNSServer"`|
|`"somethingHttpApiId"`|`"somethingHTTPAPIID"`|

## Reference

For more details, please see [golint package](https://github.com/golang/lint/blob/d0100b6bd8b389f0385611eb39152c4d7c3a7905/lint.go#L768-L810).
