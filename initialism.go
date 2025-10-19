package gocase

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	runeOffsetUpperToLower = 32
)

// DefaultInitialisms is a list of default initialisms.
// This list is based on [Staticcheck].
// For more details, please see [initialisms section].
//
// [Staticcheck]: https://staticcheck.io/
// [initialisms section]: https://staticcheck.io/docs/configuration/options/#initialisms
var DefaultInitialisms = []string{
	"ACL", "API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP",
	"HTTPS", "JSON", "QPS", "RAM", "RPC", "SLA", "SMTP", "SQL", "SSH", "TCP",
	"TLS", "TTL", "UDP", "GID", "UUID", "URI", "URL", "UTF8", "VM", "XML",
	"XMPP", "XSRF", "XSS", "SIP", "RTP", "AMQP", "DB", "TS",
	// Lower priority due to collision
	"UID", "ID", "IP", "UI",
}

// initialism is a type that describes initialization rule.
// It contains an all-uppercase string, a capitalized string (only the first letter uppercase),
// and pre-compiled regular expressions for efficient pattern matching.
type initialism struct {
	allUpperStr string
	capUpperStr string
	notEndRegex *regexp.Regexp
	endRegex    *regexp.Regexp
}

// newInitialism creates a new initialism with pre-compiled regular expressions.
// The allUpper parameter should be the all-uppercase form (e.g., "API").
// The capUpper parameter should be the capitalized form (e.g., "Api").
// Regular expressions are compiled once during initialization for better performance.
func newInitialism(allUpper, capUpper string) initialism {
	// Pre-compile regular expressions for better performance
	notEndRegex := regexp.MustCompile(fmt.Sprintf("%s([^a-z])", capUpper))
	endRegex := regexp.MustCompile(fmt.Sprintf("%s$", capUpper))

	return initialism{
		allUpperStr: allUpper,
		capUpperStr: capUpper,
		notEndRegex: notEndRegex,
		endRegex:    endRegex,
	}
}

func (i initialism) allUpper() string {
	return i.allUpperStr
}

func (i initialism) capUpper() string {
	return i.capUpperStr
}

func createInitialisms(initialisms ...string) ([]initialism, error) {
	results := make([]initialism, 0, len(initialisms))
	for _, i := range initialisms {

		s, err := convertToOnlyFirstLetterCapitalizedString(i)
		if err != nil {
			return nil, err
		}

		results = append(results, newInitialism(strings.ToUpper(i), s))
	}
	return results, nil
}

func convertToOnlyFirstLetterCapitalizedString(str string) (string, error) {
	if !utf8.ValidString(str) {
		return "", errors.New("input is not valid UTF-8")
	}

	result := []rune{}
	for i, r := range str {
		if 'A' <= r && r <= 'Z' {
			if i == 0 {
				result = append(result, r)
			} else {
				result = append(result, rune(int(r)+runeOffsetUpperToLower))
			}
		} else if 'a' <= r && r <= 'z' {
			if i == 0 {
				result = append(result, rune(int(r)-runeOffsetUpperToLower))
			} else {
				result = append(result, r)
			}
		} else if '0' <= r && r <= '9' {
			result = append(result, r)
		} else {
			return "", fmt.Errorf("input %q contains non-alphanumeric character %q at position %d", str, r, i)
		}
	}

	return string(result), nil
}
