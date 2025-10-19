# gocase

[![Go Reference](https://pkg.go.dev/badge/github.com/takuoki/gocase.svg)](https://pkg.go.dev/github.com/takuoki/gocase)
![CI](https://github.com/takuoki/gocase/actions/workflows/auto-test.yml/badge.svg)
[![codecov](https://codecov.io/gh/takuoki/gocase/branch/main/graph/badge.svg?token=s2jxPXhDjF)](https://codecov.io/gh/takuoki/gocase)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

A golang package to convert normal CamelCase to Golang's CamelCase and vice versa.

## Install

```bash
go get github.com/takuoki/gocase
```

## Usage

### Default converter

```go
str := gocase.To("IpAddress") // "IPAddress"
```

```go
str := gocase.Revert("defaultDNSServer") // "defaultDnsServer"
```

### Custom converter

```go
converter, _ := gocase.New(gocase.WithInitialisms("JSON", "CSV"))

str1 := converter.To("IpAddress") // "IpAddress" (no convert)
str2 := converter.To("JsonFile")  // "JSONFile"
str3 := converter.To("CsvFile")   // "CSVFile"
```

To add custom initialisms to the default list, use `gocase.DefaultInitialisms`:

```go
initialisms := append([]string{"JSON", "CSV"}, gocase.DefaultInitialisms...)
converter, _ := gocase.New(gocase.WithInitialisms(initialisms...))

str1 := converter.To("IpAddress") // "IPAddress"
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

## !!! WARNING !!!

Note that it is impossible to accurately determine the word break in a string of consecutive uppercase words,
so the string converted with `To` and `Revert` may not match the original string.

**example**

- `IdB` --To-> `IDB` --Revert-> `IDb`
- `UUid` --To-> `UUID` --Revert-> `Uuid`

## Reference

For more details, please see [initialisms section](https://staticcheck.io/docs/configuration/options/#initialisms) in [Staticcheck](https://staticcheck.io/).
