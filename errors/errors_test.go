package apierrors

import (
	"errors"
	"testing"
)

func TestIs(t *testing.T) {
	err := fmtWrap(New(CodeConflict, "revision mismatch"))
	if !Is(err, CodeConflict) {
		t.Fatal("expected conflict code")
	}
	if Is(err, CodeNotFound) {
		t.Fatal("unexpected not_found code")
	}
}

func fmtWrap(err error) error {
	return errors.Join(err)
}
