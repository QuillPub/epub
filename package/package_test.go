package packagefile_test

import (
	"testing"
	"time"

	packagefile "github.com/quillpub/epub/package"
)

type attribute interface {
	String() string
	Name() string
}

func expectAttribute(t *testing.T, expected, got attribute) {
	if got != expected {
		t.Errorf("expect %s value %q, got %s value %q",
			expected.Name(),
			expected.String(),
			got.Name(),
			got.String(),
		)
	}
}

func expectString(t *testing.T, expected, got string) {
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func expectCount(t *testing.T, expected, got int, description string) {
	if got != expected {
		t.Errorf("expected number of %s to be %d, but was %d", description, expected, got)
	}
}

func expectDate(t *testing.T, expected time.Time, got packagefile.Date) {
	if got.T == nil {
		t.Error("expected a time but got none")
		return
	}
	if expected != *got.T {
		t.Errorf("expected %s, got %s", expected, got.T)
	}
}
