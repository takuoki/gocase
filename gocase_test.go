package gocase_test

import (
	"testing"

	"github.com/takuoki/gocase"
)

func TestTo(t *testing.T) {
	cases := []struct {
		s, expected string
	}{
		{s: "", expected: ""},
		{s: "jsonFile", expected: "jsonFile"},
		{s: "IpAddress", expected: "IPAddress"},
		{s: "defaultDnsServer", expected: "defaultDNSServer"},
		{s: "somethingHttpApiId", expected: "somethingHTTPAPIID"},
	}

	for _, c := range cases {
		r := gocase.To(c.s)
		if r != c.expected {
			t.Errorf("value doesn't match: %s (expected %s)", r, c.expected)
		}
	}
}

func TestRevert(t *testing.T) {
	cases := []struct {
		s, expected string
	}{
		{s: "", expected: ""},
		{s: "jsonFile", expected: "jsonFile"},
		{s: "IPAddress", expected: "IpAddress"},
		{s: "defaultDNSServer", expected: "defaultDnsServer"},
		{s: "somethingHTTPAPIID", expected: "somethingHttpApiId"},
	}

	for _, c := range cases {
		r := gocase.Revert(c.s)
		if r != c.expected {
			t.Errorf("value doesn't match: %s (expected %s)", r, c.expected)
		}
	}
}
