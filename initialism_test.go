package gocase

import (
	"testing"
)

func TestCreateInitialisms(t *testing.T) {
	cases := []struct {
		in      []string
		want    []initialism
		wantErr string
	}{
		{in: []string{"ACL"}, want: []initialism{newInitialism("ACL", "Acl")}},
		{in: []string{"api", "aSCII"}, want: []initialism{newInitialism("API", "Api"), newInitialism("ASCII", "Ascii")}},
		{in: []string{"UTF!"}, wantErr: "input \"UTF!\" contains non-alphanumeric character '!' at position 3"},
	}

	for _, c := range cases {
		r, err := createInitialisms(c.in...)
		if c.wantErr == "" {
			if err != nil {
				t.Errorf("error must not be occurred: %v", err)
			} else if len(r) != len(c.want) {
				t.Errorf("value length doesn't match: %d (want %d)", len(r), len(c.want))
			} else {
				for i, w := range c.want {
					if !equalInitialism(r[i], w) {
						t.Errorf("value doesn't match at index %d: {allUpper: %s, capUpper: %s} (want {allUpper: %s, capUpper: %s})",
							i, r[i].allUpper(), r[i].capUpper(), w.allUpper(), w.capUpper())
					}
				}
			}
		} else {
			if err == nil {
				t.Error("error must be occurred")
			} else if err.Error() != c.wantErr {
				t.Errorf("error doesn't match: %v (want %s)", err, c.wantErr)
			}
		}
	}
}

func TestConvertToOnlyFirstLetterCapitalizedString(t *testing.T) {
	cases := []struct {
		in      string
		want    string
		wantErr string
	}{
		{in: "ACL", want: "Acl"},
		{in: "api", want: "Api"},
		{in: "aSCII", want: "Ascii"},
		{in: "cPu", want: "Cpu"},
		{in: "UTF8", want: "Utf8"},
		{in: "UTF!", wantErr: "input \"UTF!\" contains non-alphanumeric character '!' at position 3"},
		{in: "aa\xe2", wantErr: "input is not valid UTF-8"},
	}

	for _, c := range cases {
		r, err := convertToOnlyFirstLetterCapitalizedString(c.in)
		if c.wantErr == "" {
			if err != nil {
				t.Errorf("error must not be occurred: %v", err)
			} else if r != c.want {
				t.Errorf("value doesn't match: %s (want %s)", r, c.want)
			}
		} else {
			if err == nil {
				t.Error("error must be occurred")
			} else if err.Error() != c.wantErr {
				t.Errorf("error doesn't match: %v (want %s)", err, c.wantErr)
			}
		}
	}
}

// equalInitialism compares two initialism values for equality.
// Since initialism contains regexp pointers, we compare the string values
// and verify that regex fields are not nil.
func equalInitialism(a, b initialism) bool {
	if a.allUpper() != b.allUpper() || a.capUpper() != b.capUpper() {
		return false
	}
	// Verify that regex patterns are not nil (they should be initialized)
	if a.notEndRegex == nil || a.endRegex == nil {
		return false
	}
	if b.notEndRegex == nil || b.endRegex == nil {
		return false
	}
	return true
}
