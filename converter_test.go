package gocase_test

import (
	"testing"

	"github.com/takuoki/gocase"
)

func TestNew(t *testing.T) {
	cases := []struct {
		opts    []gocase.Option
		wantErr string
	}{
		{opts: []gocase.Option{gocase.WithInitialisms("JSON", "CSV")}},
		{opts: []gocase.Option{gocase.WithInitialisms("UTF8", "UTF!")}, wantErr: "input \"UTF!\" contains non-alphanumeric character '!' at position 3"},
	}

	for _, c := range cases {
		r, err := gocase.New(c.opts...)
		if c.wantErr == "" {
			if err != nil {
				t.Errorf("error must not be occurred: %v", err)
			} else if r == nil {
				t.Error("value must not be nil")
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
