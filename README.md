# gocase

[![GoDoc](https://godoc.org/github.com/takuoki/gocase?status.svg)](https://godoc.org/github.com/takuoki/gocase)
[![CircleCI](https://circleci.com/gh/takuoki/gocase/tree/master.svg?style=shield&circle-token=06a2582cde9cc3c182873c3ec5dddb67e9388cf6)](https://circleci.com/gh/takuoki/gocase/tree/master)
[![codecov](https://codecov.io/gh/takuoki/gocase/branch/master/graph/badge.svg)](https://codecov.io/gh/takuoki/gocase)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

A golang package to convert normal CamelCase to Golang's CamelCase and vice versa.
Golang's CamelCase means a string that takes into account to Go's common initialisms.

## Install

```cmd
go get -u github.com/takuoki/gocase
```

## Usage

### Default converter

```go
str := gocase.To("IpAddress") // "IPAddress"
```

```go
str := gocase.To("defaultDNSServer") // "defaultDnsServer"
```

### Custom converter

```go
converter, _ := gocase.New(gocase.WithInitialisms("JSON", "CSV"))

str1 := converter.To("IpAddress") // "IpAddress" (no convert)
str2 := converter.To("JsonFile")  // "JSONFile"
str3 := converter.To("CsvFile")   // "CSVFile"
```

## Example

The default converter converts as follows

- `To`: converts a string from `From` to `To`.
- `Revert`: converts a string from `To` to `From`.

| From                 | To                   |
| :------------------- | :------------------- |
| ""                   | ""                   |
| "jsonFile"           | "jsonFile"           |
| "IpAddress"          | "IPAddress"          |
| "defaultDnsServer"   | "defaultDNSServer"   |
| "somethingHttpApiId" | "somethingHTTPAPIID" |
| "somethingUuid"      | "somethingUUID"      |
| "somethingSip"       | "somethingSIP"       |

## Reference

For more details, please see [initialisms section](https://staticcheck.io/docs/configuration/options/#initialisms) in [Staticcheck](https://staticcheck.io/).
