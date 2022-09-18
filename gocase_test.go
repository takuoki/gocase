package gocase_test

import (
	"testing"

	"github.com/takuoki/gocase"
)

func TestConverter_To(t *testing.T) {
	dc, _ := gocase.New()
	cc, _ := gocase.New(gocase.WithInitialisms("JSON", "CSV"))

	cases := []struct {
		conv    *gocase.Converter
		s, want string
	}{
		{conv: dc, s: "", want: ""},
		{conv: dc, s: "jsonFile", want: "jsonFile"},
		{conv: dc, s: "IpAddress", want: "IPAddress"},
		{conv: dc, s: "defaultDnsServer", want: "defaultDNSServer"},
		{conv: dc, s: "somethingHttpApiId", want: "somethingHTTPAPIID"},
		{conv: dc, s: "somethingUuid", want: "somethingUUID"},
		{conv: dc, s: "somethingSip", want: "somethingSIP"},
		{conv: cc, s: "JsonFile", want: "JSONFile"},
		{conv: cc, s: "CsvFile", want: "CSVFile"},
		{conv: cc, s: "IpAddress", want: "IpAddress"},
	}

	for _, c := range cases {
		r := c.conv.To(c.s)
		if r != c.want {
			t.Errorf("value doesn't match: %s (want %s)", r, c.want)
		}
	}
}

func TestConverter_Revert(t *testing.T) {
	dc, _ := gocase.New()
	cc, _ := gocase.New(gocase.WithInitialisms("JSON", "CSV"))

	cases := []struct {
		conv    *gocase.Converter
		s, want string
	}{
		{conv: dc, s: "", want: ""},
		{conv: dc, s: "jsonFile", want: "jsonFile"},
		{conv: dc, s: "IPAddress", want: "IpAddress"},
		{conv: dc, s: "defaultDNSServer", want: "defaultDnsServer"},
		{conv: dc, s: "somethingHTTPAPIID", want: "somethingHttpApiId"},
		{conv: dc, s: "somethingUUID", want: "somethingUuid"},
		{conv: dc, s: "somethingSIP", want: "somethingSip"},
		{conv: cc, s: "JSONFile", want: "JsonFile"},
		{conv: cc, s: "CSVFile", want: "CsvFile"},
		{conv: cc, s: "somethingSIP", want: "somethingSIP"},
	}

	for _, c := range cases {
		r := c.conv.Revert(c.s)
		if r != c.want {
			t.Errorf("value doesn't match: %s (want %s)", r, c.want)
		}
	}
}
