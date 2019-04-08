# gocase

[![CircleCI](https://circleci.com/gh/takuoki/gocase/tree/master.svg?style=shield&circle-token=06a2582cde9cc3c182873c3ec5dddb67e9388cf6)](https://circleci.com/gh/takuoki/gocase/tree/master)
[![codecov](https://codecov.io/gh/takuoki/gocase/branch/master/graph/badge.svg)](https://codecov.io/gh/takuoki/gocase)
[![GoDoc](https://godoc.org/github.com/takuoki/gocase?status.svg)](https://godoc.org/github.com/takuoki/gocase)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

A golang package to convert normal CamelCase to Golang's CamelCase and vice versa.
Golang's CamelCase means a string that takes into account to Go's common initialisms.

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
