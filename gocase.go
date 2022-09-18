// Package gocase is a package to convert normal CamelCase to Golang's CamelCase and vice versa.
// Golang's CamelCase means a string that takes into account to Go's common initialisms.
// For more details, please see [initialisms section] in [Staticcheck].
//
// [Staticcheck]: https://staticcheck.io/
// [initialisms section]: https://staticcheck.io/docs/configuration/options/#initialisms
package gocase

import (
	"strings"
)

// To returns a string converted to Go case.
func To(s string) string {
	return defaultConverter.To(s)
}

// To returns a string converted to Go case with converter.
func (c *Converter) To(s string) string {
	for _, i := range c.initialisms {
		s = strings.Replace(s, i[1], i[0], -1)
	}
	return s
}

// Revert returns a string converted from Go case to normal case.
func Revert(s string) string {
	return defaultConverter.Revert(s)
}

// Revert returns a string converted from Go case to normal case with converter.
func (c *Converter) Revert(s string) string {
	for _, i := range c.initialisms {
		s = strings.Replace(s, i[0], i[1], -1)
	}
	return s
}
